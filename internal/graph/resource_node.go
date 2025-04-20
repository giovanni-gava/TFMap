// File: internal/graph/resource_node.go

package graph

type ResourceNode struct {
	ID         string
	Type       string
	Name       string
	Provider   string
	Attributes map[string]interface{}
	DependsOn  []string
	ModulePath string
	SourceFile string
	LineNumber int
	Tags       map[string]string
	Metadata   ResourceMetadata
}

type Edge struct {
	From string
	To   string
	Type string
}

type InfraMetadata struct {
	Project     string
	Workspace   string
	Environment string
	SourcePaths []string
	CreatedBy   string
	Timestamp   string
	Version     string
}

type ResourceMetadata struct {
	ManagedBy     string
	SourceVersion string
	Warnings      []string
	Validated     bool
	Drifted       bool
}
