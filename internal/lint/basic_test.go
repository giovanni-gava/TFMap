package lint_test

import (
	"testing"

	"github.com/giovanni-gava/tfmap/internal/graph"
	"github.com/giovanni-gava/tfmap/internal/lint"
	"github.com/stretchr/testify/assert"
)

func TestCheckMissingTags(t *testing.T) {
	g := graph.NewInfraGraph()

	// Caso 1: recurso sem nenhuma tag
	g.AddResource(&graph.ResourceNode{
		ID:         "aws_s3_bucket.no_tags",
		Type:       "aws_s3_bucket",
		Name:       "no_tags",
		SourceFile: "infra/no_tags.tf",
		LineNumber: 3,
	})

	// Caso 2: recurso com apenas uma tag
	g.AddResource(&graph.ResourceNode{
		ID:         "aws_s3_bucket.partial_tags",
		Type:       "aws_s3_bucket",
		Name:       "partial_tags",
		Tags:       map[string]string{"Name": "only-name"},
		SourceFile: "infra/partial.tf",
		LineNumber: 7,
	})

	// Caso 3: recurso válido com todas as tags
	g.AddResource(&graph.ResourceNode{
		ID:   "aws_s3_bucket.valid",
		Type: "aws_s3_bucket",
		Name: "valid",
		Tags: map[string]string{
			"Name":        "ok",
			"Environment": "dev",
		},
	})

	results := lint.RunAll(g)

	// Deve haver exatamente 2 problemas: no_tags e partial_tags
	assert.Len(t, results, 2)

	// Validação por conteúdo e não por ordem
	resultMap := map[string]lint.LintResult{}
	for _, r := range results {
		resultMap[r.ResourceID] = r
	}

	r1, ok := resultMap["aws_s3_bucket.no_tags"]
	assert.True(t, ok)
	assert.Equal(t, lint.LevelWarning, r1.Level)
	assert.Equal(t, "missing_tags", r1.Rule)
	assert.Contains(t, r1.Message, "no tags assigned")

	r2, ok := resultMap["aws_s3_bucket.partial_tags"]
	assert.True(t, ok)
	assert.Equal(t, lint.LevelWarning, r2.Level)
	assert.Equal(t, "missing_tags", r2.Rule)
	assert.Contains(t, r2.Message, "Missing required tags")
}

func TestCheckWildcardIAM(t *testing.T) {
	g := graph.NewInfraGraph()

	// Caso 1: Política com wildcard na Action
	g.AddResource(&graph.ResourceNode{
		ID:   "aws_iam_policy.bad_policy",
		Type: "aws_iam_policy",
		Name: "bad_policy",
		Attributes: map[string]interface{}{
			"policy": `{
				"Version": "2012-10-17",
				"Statement": [{
					"Effect": "Allow",
					"Action": "*",
					"Resource": "*"
				}]
			}`,
		},
		SourceFile: "iam/bad.tf",
		LineNumber: 10,
	})

	// Caso 2: Política válida e restrita
	g.AddResource(&graph.ResourceNode{
		ID:   "aws_iam_policy.good_policy",
		Type: "aws_iam_policy",
		Name: "good_policy",
		Attributes: map[string]interface{}{
			"policy": `{
				"Version": "2012-10-17",
				"Statement": [{
					"Effect": "Allow",
					"Action": "s3:ListBucket",
					"Resource": "arn:aws:s3:::example-bucket"
				}]
			}`,
		},
	})

	results := lint.RunAll(g)

	// Deve haver exatamente 1 erro
	assert.Len(t, results, 1)

	r := results[0]
	assert.Equal(t, "aws_iam_policy.bad_policy", r.ResourceID)
	assert.Equal(t, lint.LevelError, r.Level)
	assert.Equal(t, "wildcard_iam", r.Rule)
	assert.Contains(t, r.Message, "wildcard")
}
