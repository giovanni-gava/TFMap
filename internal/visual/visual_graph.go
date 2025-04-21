package visual

import (
	"github.com/giovanni-gava/tfmap/internal/graph"
)

type VisualNode struct {
	ID         string
	Label      string
	Type       string
	Group      string
	PositionX  float64
	PositionY  float64
	Color      string
	Shape      string
	Icon       string
	Highlight  bool
	SourceFile string
	LineNumber int
}

type VisualEdge struct {
	From   string
	To     string
	Label  string
	Dashed bool
	Color  string
}

// VisualGraph é o grafo visual completo da infraestrutura.
type VisualGraph struct {
	Nodes []VisualNode
	Edges []VisualEdge
}

// NewFromInfraGraph converte um InfraGraph semântico em um VisualGraph com estilo default.
func NewFromInfraGraph(infra *graph.InfraGraph) *VisualGraph {
	vg := &VisualGraph{}

	for _, res := range infra.Resources {
		vg.Nodes = append(vg.Nodes, VisualNode{
			ID:         res.ID,
			Label:      res.Name,
			Type:       res.Type,
			Group:      res.Type, // agrupar por tipo inicialmente
			Color:      defaultColorForType(res.Type),
			Shape:      "rect",
			Icon:       iconForType(res.Type),
			SourceFile: res.SourceFile,
			LineNumber: res.LineNumber,
		})
	}

	for _, rel := range infra.Relationships {
		vg.Edges = append(vg.Edges, VisualEdge{
			From:   rel.From,
			To:     rel.To,
			Label:  rel.Type,
			Dashed: rel.Type != "depends_on",
			Color:  edgeColor(rel.Type),
		})
	}

	return vg
}

// Helpers (mock simples — podem evoluir para dicionários visuais reais)

func defaultColorForType(t string) string {
	if t == "aws_s3_bucket" {
		return "#fcd34d" // amarelo bucket
	} else if t == "aws_vpc" {
		return "#60a5fa" // azul VPC
	} else if t == "aws_iam_policy" {
		return "#f87171" // vermelho policy
	}
	return "#d1d5db" // cinza neutro
}

func iconForType(t string) string {
	if t == "aws_s3_bucket" {
		return "s3"
	} else if t == "aws_vpc" {
		return "vpc"
	} else if t == "aws_iam_policy" {
		return "shield"
	}
	return "box"
}

func edgeColor(t string) string {
	if t == "depends_on" {
		return "#6b7280" // cinza
	}
	return "#9ca3af" // mais claro
}
