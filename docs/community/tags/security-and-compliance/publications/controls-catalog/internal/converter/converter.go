package converter

import (
	_ "embed"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"text/template"

	"github.com/gemaraproj/go-gemara"
	"gopkg.in/yaml.v3"
)

//go:embed markdown.tmpl
var markdownTemplate string

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

// familyWithGuidelines represents a family with its associated guidelines for template rendering.
type familyWithGuidelines struct {
	gemara.Family
	Guidelines []gemara.Guideline
}

// templateData holds the data structure for the markdown template.
type templateData struct {
	FamiliesWithGuidelines []familyWithGuidelines
}

// ToMarkdown converts a GuidanceDocument to Markdown format for website rendering.
func ToMarkdown(doc *gemara.GuidanceDocument) (string, error) {
	familyGuidelines := make(map[string][]gemara.Guideline)
	for _, guideline := range doc.Guidelines {
		if guideline.Family != "" {
			familyGuidelines[guideline.Family] = append(familyGuidelines[guideline.Family], guideline)
		}
	}

	for familyID := range familyGuidelines {
		sort.Slice(familyGuidelines[familyID], func(i, j int) bool {
			return familyGuidelines[familyID][i].Id < familyGuidelines[familyID][j].Id
		})
	}

	var familiesWithGuidelines []familyWithGuidelines
	for _, family := range doc.Families {
		guidelines, hasGuidelines := familyGuidelines[family.Id]
		if hasGuidelines && len(guidelines) > 0 {
			familiesWithGuidelines = append(familiesWithGuidelines, familyWithGuidelines{
				Family:     family,
				Guidelines: guidelines,
			})
		}
	}

	data := templateData{
		FamiliesWithGuidelines: familiesWithGuidelines,
	}

	tmpl, err := template.New("markdown").Funcs(template.FuncMap{
		"anchor": func(s string) string {
			return strings.ToLower(strings.ReplaceAll(s, " ", "-"))
		},
		"lower": func(s string) string {
			return strings.ToLower(s)
		},
	}).Parse(markdownTemplate)
	if err != nil {
		return "", fmt.Errorf("failed to parse template: %w", err)
	}

	var b strings.Builder
	if err := tmpl.Execute(&b, data); err != nil {
		return "", fmt.Errorf("failed to execute template: %w", err)
	}

	return b.String(), nil
}
