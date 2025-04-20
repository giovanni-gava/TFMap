package terraform_test

import (
	"path/filepath"
	"testing"

	"github.com/giovanni-gava/tfmap/internal/parser/terraform"
	"github.com/stretchr/testify/assert"
)

func TestTerraformParser_ParseValidTFFile(t *testing.T) {
	parser := terraform.NewParser()
	testPath := filepath.Join("testdata", "unit", "basic")

	graph, err := parser.Parse(testPath)
	assert.NoError(t, err, "expected no error from parser")
	assert.NotNil(t, graph, "expected a non-nil graph")

	assert.Equal(t, 1, len(graph.Resources), "expected one resource in the graph")

	resource, ok := graph.Resources["aws_s3_bucket.example"]
	assert.True(t, ok, "resource ID aws_s3_bucket.example should exist")

	assert.Equal(t, "aws_s3_bucket", resource.Type)
	assert.Equal(t, "example", resource.Name)
	assert.Equal(t, "tfmap-example-bucket", resource.Attributes["bucket"])
	assert.Equal(t, "private", resource.Attributes["acl"])

	// Tags
	assert.Equal(t, "ExampleBucket", resource.Tags["Name"])
	assert.Equal(t, "Dev", resource.Tags["Environment"])
}
