// SPDX-FileCopyrightText: Copyright 2025 CNCF Contributors
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/ossf/gemara"
	"gopkg.in/yaml.v3"
)

// FamilyYAML represents a family YAML file in security-baseline format
type FamilyYAML struct {
	Title       string            `yaml:"title"`
	Description string            `yaml:"description"`
	Controls    []gemara.Guideline `yaml:"controls"`
}

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <families-directory> [output-file]\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "  families-directory: Directory containing family YAML files (e.g., access.yaml, compute.yaml)\n")
		fmt.Fprintf(os.Stderr, "  output-file: Output markdown file (default: index.md)\n")
		os.Exit(1)
	}

	familiesDir := os.Args[1]
	outputFile := "index.md"
	if len(os.Args) > 2 {
		outputFile = os.Args[2]
	}

	// Define family order
	familyOrder := []string{
		"access",
		"compute",
		"deploy",
		"develop",
		"distribute",
		"securing-artefacts",
		"securing-build-pipelines",
		"securing-deployments",
		"securing-materials",
		"securing-the-source-code",
		"security-assurance",
		"storage",
	}

	// Collect all families
	families := make(map[string]*FamilyYAML)

	// Read all family YAML files
	for _, familyID := range familyOrder {
		familyFile := filepath.Join(familiesDir, fmt.Sprintf("%s.yaml", familyID))
		
		// Check if file exists
		if _, err := os.Stat(familyFile); os.IsNotExist(err) {
			continue
		}

		// Read YAML file
		data, err := os.ReadFile(familyFile)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Warning: failed to read %s: %v\n", familyFile, err)
			continue
		}

		var family FamilyYAML
		if err := yaml.Unmarshal(data, &family); err != nil {
			fmt.Fprintf(os.Stderr, "Warning: failed to parse %s: %v\n", familyFile, err)
			continue
		}

		families[familyID] = &family
	}

	// Sort controls within each family by ID
	for _, family := range families {
		sort.Slice(family.Controls, func(i, j int) bool {
			return family.Controls[i].Id < family.Controls[j].Id
		})
	}

	// Generate markdown
	markdown := generateMarkdown(families, familyOrder)

	// Write output
	if err := os.WriteFile(outputFile, []byte(markdown), 0644); err != nil {
		fmt.Fprintf(os.Stderr, "Error writing output: %v\n", err)
		os.Exit(1)
	}

	totalControls := 0
	for _, family := range families {
		totalControls += len(family.Controls)
	}

	fmt.Printf("✅ Generated %s with %d families and %d total controls\n", outputFile, len(families), totalControls)
}

func generateMarkdown(families map[string]*FamilyYAML, familyOrder []string) string {
	var b strings.Builder

	// Header
	b.WriteString("---\n")
	b.WriteString("title: Cloud Native Security Controls Catalog\n")
	b.WriteString("sidebar_position: 1\n")
	b.WriteString("---\n\n")

	b.WriteString("# Cloud Native Security Controls Catalog\n\n")
	b.WriteString("The Cloud Native Security Controls Catalog provides comprehensive guidance for securing cloud-native applications and workloads.\n\n")

	// Table of Contents
	b.WriteString("## Table of Contents\n\n")
	for _, familyID := range familyOrder {
		if family, exists := families[familyID]; exists {
			b.WriteString(fmt.Sprintf("- [%s](#%s)\n", family.Title, familyID))
		}
	}
	b.WriteString("\n---\n\n")

	// Generate content for each family
	for _, familyID := range familyOrder {
		family, exists := families[familyID]
		if !exists {
			continue
		}

		b.WriteString(fmt.Sprintf("## %s {#%s}\n\n", family.Title, familyID))
		if family.Description != "" {
			b.WriteString(fmt.Sprintf("%s\n\n", family.Description))
		}

		// List all controls in this family
		for _, control := range family.Controls {
			b.WriteString(fmt.Sprintf("### %s {#%s}\n\n", control.Title, strings.ToLower(control.Id)))
			b.WriteString(fmt.Sprintf("**Control ID**: `%s`\n\n", control.Id))

			if control.Objective != "" {
				b.WriteString("#### Objective\n\n")
				b.WriteString(fmt.Sprintf("%s\n\n", control.Objective))
			}

			// Guideline Mappings
			if len(control.GuidelineMappings) > 0 {
				b.WriteString("#### Guideline Mappings\n\n")
				for _, mapping := range control.GuidelineMappings {
					b.WriteString(fmt.Sprintf("**%s**\n\n", mapping.ReferenceId))
					if len(mapping.Entries) > 0 {
						b.WriteString("| Reference ID | Strength | Remarks |\n")
						b.WriteString("|--------------|----------|----------|\n")
						for _, entry := range mapping.Entries {
							b.WriteString(fmt.Sprintf("| %s | %d | %s |\n",
								entry.ReferenceId, entry.Strength, entry.Remarks))
						}
						b.WriteString("\n")
					}
				}
			}

			// Recommendations
			if len(control.Recommendations) > 0 {
				b.WriteString("#### Recommendations\n\n")
				for _, rec := range control.Recommendations {
					b.WriteString(fmt.Sprintf("- %s\n", rec))
				}
				b.WriteString("\n")
			}

			b.WriteString("---\n\n")
		}
	}

	return b.String()
}
