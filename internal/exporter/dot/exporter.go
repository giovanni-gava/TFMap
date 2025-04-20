package dot

import (
	"fmt"
	"os"
	"strings"

	"github.com/giovanni-gava/tfmap/internal/graph"
)

type Exporter struct{}

func NewExporter() *Exporter {
	return &Exporter{}
}

func (e *Exporter) Export(g *graph.InfraGraph, outputPath string) error {
	var b strings.Builder

	b.WriteString("digraph tfmap {\n")
	b.WriteString("  rankdir=LR;\n")
	b.WriteString("  node [shape=box style=filled fillcolor=lightgrey fontname=Helvetica];\n")

	// Nós (resources)
	for _, res := range g.Resources {
		label := fmt.Sprintf("%s\\n%s", res.Type, res.Name)
		color := inferColor(res.Type)

		b.WriteString(fmt.Sprintf("  \"%s\" [label=\"%s\" fillcolor=\"%s\"];\n", res.ID, label, color))
	}

	// Arestas (edges)
	for _, e := range g.Edges {
		style := inferEdgeStyle(e.Type)
		b.WriteString(fmt.Sprintf("  \"%s\" -> \"%s\" [style=%s];\n", e.From, e.To, style))
	}

	b.WriteString("}\n")

	return os.WriteFile(outputPath, []byte(b.String()), 0644)
}

func inferColor(resourceType string) string {
	switch {
	case strings.HasPrefix(resourceType, "aws_iam"):
		return "lightcoral"
	case strings.HasPrefix(resourceType, "aws_s3"):
		return "lightblue"
	case strings.HasPrefix(resourceType, "aws_lambda"):
		return "lightgoldenrod"
	case strings.HasPrefix(resourceType, "aws_"):
		return "lightgreen"
	default:
		return "white"
	}
}

func inferEdgeStyle(edgeType string) string {
	switch edgeType {
	case "explicit":
		return "bold"
	case "implicit":
		return "dashed"
	default:
		return "solid"
	}
}
