package lint

import (
	"fmt"

	"github.com/giovanni-gava/tfmap/internal/graph"
)

func checkMissingTags(g *graph.InfraGraph) []LintResult {
	var results []LintResult

	for _, res := range g.Resources {
		// Ignore recursos internos ou de providers sem suporte
		if res.Type == "" || res.Name == "" {
			continue
		}

		if len(res.Tags) == 0 {
			results = append(results, LintResult{
				ResourceID: res.ID,
				Level:      LevelWarning,
				Rule:       "missing_tags",
				Message:    "Resource has no tags assigned",
				Suggestion: "Add at least Name and Environment tags for traceability and cost management",
				File:       res.SourceFile,
				Line:       res.LineNumber,
			})
			continue
		}

		// Verifica se tags importantes estão faltando
		required := []string{"Name", "Environment"}
		missing := []string{}

		for _, key := range required {
			if _, ok := res.Tags[key]; !ok {
				missing = append(missing, key)
			}
		}

		if len(missing) > 0 {
			results = append(results, LintResult{
				ResourceID: res.ID,
				Level:      LevelWarning,
				Rule:       "missing_tags",
				Message:    fmt.Sprintf("Missing required tags: %v", missing),
				Suggestion: "Include standard tags like Name and Environment to comply with best practices",
				File:       res.SourceFile,
				Line:       res.LineNumber,
			})
		}
	}

	return results
}

