package lint

import (
	"fmt"
	"strings"

	"github.com/giovanni-gava/tfmap/internal/graph"
)

func checkMissingTags(g *graph.InfraGraph) []LintResult {
	var results []LintResult

	for _, res := range g.Resources {
		// Ignora tipos que geralmente não precisam de tags
		skipTypes := map[string]bool{
			"aws_iam_policy": true,
			"aws_iam_role":   true,
			"aws_iam_user":   true,
		}
		if skipTypes[res.Type] {
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

func checkWildcardIAM(g *graph.InfraGraph) []LintResult {
	var results []LintResult

	for _, res := range g.Resources {
		if res.Type != "aws_iam_policy" && res.Type != "aws_iam_role" {
			continue
		}

		policyRaw, ok := res.Attributes["policy"]
		if !ok {
			continue
		}

		policyStr, ok := policyRaw.(string)
		if !ok {
			continue
		}

		// Normalização simples
		policyStr = strings.ToLower(strings.ReplaceAll(policyStr, " ", ""))

		if strings.Contains(policyStr, `"action":"*"`) ||
			strings.Contains(policyStr, `"resource":"*"`) {

			results = append(results, LintResult{
				ResourceID: res.ID,
				Level:      LevelError,
				Rule:       "wildcard_iam",
				Message:    "IAM policy contains wildcard action or resource",
				Suggestion: "Restrict IAM policies to least privilege. Avoid use of \"*\" in Action or Resource.",
				File:       res.SourceFile,
				Line:       res.LineNumber,
			})
		}
	}

	return results
}
