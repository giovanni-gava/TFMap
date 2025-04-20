package lint_test

import (
	"path/filepath"
	"testing"

	"github.com/giovanni-gava/tfmap/internal/parser/terraform"
	"github.com/giovanni-gava/tfmap/internal/lint"
	"github.com/stretchr/testify/assert"
)

func TestLintDetectsMissingTagsInRealTF(t *testing.T) {
	parser := terraform.NewParser()
	path := filepath.Join("testdata", "no-tags")

	graph, err := parser.Parse(path)
	assert.NoError(t, err)
	assert.NotNil(t, graph)

	results := lint.RunAll(graph)
	assert.GreaterOrEqual(t, len(results), 1, "esperava pelo menos 1 erro de lint")

	assert.Equal(t, "missing_tags", results[0].Rule)
	assert.Equal(t, "aws_s3_bucket.broken", results[0].ResourceID)
}

func TestLintPassesValidTaggedResource(t *testing.T) {
	parser := terraform.NewParser()
	path := filepath.Join("testdata", "valid-tags")

	graph, err := parser.Parse(path)
	assert.NoError(t, err)
	assert.NotNil(t, graph)

	results := lint.RunAll(graph)
	assert.Len(t, results, 0, "esperava zero problemas com recurso corretamente tagueado")
}

