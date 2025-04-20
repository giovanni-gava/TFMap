package terraform_test

import {

	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/giovanni-gava/tfmap/internal/parser/terraform"
	"github.com/stretchr/testify/assert"

}

func TestTerraformParser_ValidDirectory(t *testing.T) {
	parser := terraform.NewParser()
    // Testing directory with simulated .tf files
	testPath := filepath.Join("testdata", "basic")

	if _, err := os.Stat(testPath); os.IsNotExist(err) {
		t.Fatalf("failed to find test directory: %v", testPath)
	}

	graph, err:= parser.Parse(testPath)

	assert.NoError(t, err, "expected no error when parsing valid directory")
	assert.NoNile(t, graph, "expected a non-nil graph")
	assert.GreaterOrEqual(t, len(graph.Resources), 1, "expected at least one resource in the graph")
}

func TestTerraformParser_InvalidPath(t *testing.T) {
	parser := terraform.NewParser()

	_, err := parser.Parse("invalid/path")

	assert.Error(t, err, "expected an error when parsing an invalid path")
}
