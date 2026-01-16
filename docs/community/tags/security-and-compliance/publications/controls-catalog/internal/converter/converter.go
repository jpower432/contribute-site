package converter

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/gemaraproj/go-gemara"
	"gopkg.in/yaml.v3"
)

// ToGemara loads all family YAML files and concatenates them into a single GuidanceDocument.
func ToGemara(catalogDir string, familyOrder []string) (*gemara.GuidanceDocument, error) {
	metadataFile := filepath.Join(catalogDir, "metadata.yaml")
	var metadataDoc gemara.GuidanceDocument
	if err := metadataDoc.LoadFile("file://" + metadataFile); err != nil {
		return nil, fmt.Errorf("failed to load metadata.yaml: %w", err)
	}

	familiesFile := filepath.Join(catalogDir, "families.yaml")
	familiesData, err := os.ReadFile(familiesFile)
	if err != nil {
		return nil, fmt.Errorf("failed to read families.yaml: %w", err)
	}

	var familiesDoc struct {
		Families []gemara.Family `yaml:"families"`
	}
	if err := yaml.Unmarshal(familiesData, &familiesDoc); err != nil {
		return nil, fmt.Errorf("failed to parse families.yaml: %w", err)
	}

	var allGuidelines []gemara.Guideline

	for _, familyID := range familyOrder {
		familyFilePath := filepath.Join(catalogDir, fmt.Sprintf("%s.yaml", familyID))

		if _, err := os.Stat(familyFilePath); os.IsNotExist(err) {
			continue
		}

		// Read YAML file
		data, err := os.ReadFile(familyFilePath)
		if err != nil {
			return nil, fmt.Errorf("failed to read %s: %w", familyFilePath, err)
		}

		var familyData struct {
			Guidelines []gemara.Guideline `yaml:"guidelines"`
		}
		if err := yaml.Unmarshal(data, &familyData); err != nil {
			return nil, fmt.Errorf("failed to parse %s: %w", familyFilePath, err)
		}

		allGuidelines = append(allGuidelines, familyData.Guidelines...)
	}

	sort.Slice(allGuidelines, func(i, j int) bool {
		return allGuidelines[i].Id < allGuidelines[j].Id
	})

	doc := &gemara.GuidanceDocument{
		Title:        metadataDoc.Title,
		Metadata:     metadataDoc.Metadata,
		DocumentType: metadataDoc.DocumentType,
		Families:     familiesDoc.Families,
		Guidelines:   allGuidelines,
	}

	return doc, nil
}

// ToMarkdown converts a GuidanceDocument to Markdown format for website rendering.
func ToMarkdown(doc *gemara.GuidanceDocument) string {
	var b strings.Builder

	b.WriteString("---\n")
	b.WriteString("title: Cloud Native Security Controls Catalog\n")
	b.WriteString("sidebar_position: 1\n")
	b.WriteString("---\n\n")

	b.WriteString("# Cloud Native Security Controls Catalog\n\n")
	b.WriteString("The Cloud Native Security Controls Catalog provides comprehensive guidance for securing cloud-native applications and workloads.\n\n")

	familyMap := make(map[string]gemara.Family)
	for _, family := range doc.Families {
		familyMap[family.Id] = family
	}

	familyGuidelines := make(map[string][]gemara.Guideline)
	for _, guideline := range doc.Guidelines {
		if guideline.Family != "" {
			familyGuidelines[guideline.Family] = append(familyGuidelines[guideline.Family], guideline)
		}
	}

	// Sort guidelines within each family by ID
	for familyID := range familyGuidelines {
		sort.Slice(familyGuidelines[familyID], func(i, j int) bool {
			return familyGuidelines[familyID][i].Id < familyGuidelines[familyID][j].Id
		})
	}

	b.WriteString("## Table of Contents\n\n")
	for _, family := range doc.Families {
		if _, hasGuidelines := familyGuidelines[family.Id]; hasGuidelines {
			anchor := strings.ToLower(strings.ReplaceAll(family.Id, " ", "-"))
			b.WriteString(fmt.Sprintf("- [%s](#%s)\n", family.Title, anchor))
		}
	}
	b.WriteString("\n---\n\n")

	for _, family := range doc.Families {
		guidelines, hasGuidelines := familyGuidelines[family.Id]
		if !hasGuidelines || len(guidelines) == 0 {
			continue
		}

		anchor := strings.ToLower(strings.ReplaceAll(family.Id, " ", "-"))
		b.WriteString(fmt.Sprintf("## %s {#%s}\n\n", family.Title, anchor))
		if family.Description != "" {
			b.WriteString(fmt.Sprintf("%s\n\n", family.Description))
		}

		for _, guideline := range guidelines {
			b.WriteString(fmt.Sprintf("### %s {#%s}\n\n", guideline.Title, strings.ToLower(guideline.Id)))
			b.WriteString(fmt.Sprintf("**Guideline ID**: `%s`\n\n", guideline.Id))

			if guideline.Objective != "" {
				b.WriteString("#### Objective\n\n")
				b.WriteString(fmt.Sprintf("%s\n\n", guideline.Objective))
			}

			// Guideline Mappings
			if len(guideline.GuidelineMappings) > 0 {
				b.WriteString("#### Guideline Mappings\n\n")
				for _, mapping := range guideline.GuidelineMappings {
					b.WriteString(fmt.Sprintf("**%s**\n\n", mapping.ReferenceId))
					if len(mapping.Entries) > 0 {
						b.WriteString("| Reference ID | Strength | Remarks |\n")
						b.WriteString("|--------------|----------|----------|\n")
						for _, entry := range mapping.Entries {
							strength := ""
							if entry.Strength > 0 {
								strength = fmt.Sprintf("%d", entry.Strength)
							}
							b.WriteString(fmt.Sprintf("| %s | %s | %s |\n",
								entry.ReferenceId, strength, entry.Remarks))
						}
						b.WriteString("\n")
					}
				}
			}

			if len(guideline.Statements) > 0 {
				b.WriteString("#### Statements\n\n")
				for _, stmt := range guideline.Statements {
					if stmt.Title != "" {
						b.WriteString(fmt.Sprintf("**%s**\n\n", stmt.Title))
					}
					b.WriteString(fmt.Sprintf("%s\n\n", stmt.Text))
				}
			}

			if len(guideline.Recommendations) > 0 {
				b.WriteString("#### Recommendations\n\n")
				for _, rec := range guideline.Recommendations {
					b.WriteString(fmt.Sprintf("- %s\n", rec))
				}
				b.WriteString("\n")
			}

			b.WriteString("---\n\n")
		}
	}

	return b.String()
}
