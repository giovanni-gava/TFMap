package dot_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/giovanni-gava/tfmap/internal/exporter/dot"
	"github.com/giovanni-gava/tfmap/internal/graph"
	"github.com/stretchr/testify/assert"
)

func TestDotExporter_BasicGraph(t *testing.T) {
	exporter := dot.NewExporter()
	tmpDir := t.TempDir()
	outPath := filepath.Join(tmpDir, "test.dot")

	g := graph.NewInfraGraph()
	g.AddResource(&graph.ResourceNode{
		ID:   "aws_s3_bucket.example",
		Type: "aws_s3_bucket",
		Name: "example",
	})
	g.AddResource(&graph.ResourceNode{
		ID:   "aws_iam_role.myrole",
		Type: "aws_iam_role",
		Name: "myrole",
	})
	g.AddEdge("aws_iam_role.myrole", "aws_s3_bucket.example", "explicit")

	err := exporter.Export(g, outPath)
	assert.NoError(t, err)
	assert.FileExists(t, outPath)

	content, _ := os.ReadFile(outPath)
	assert.Contains(t, string(content), "aws_s3_bucket\\nexample")
	assert.Contains(t, string(content), "aws_iam_role\\nmyrole")
	assert.Contains(t, string(content), "->")
}
