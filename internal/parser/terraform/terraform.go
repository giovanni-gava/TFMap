package terraform

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/giovanni-gava/tfmap/internal/graph"
)

// Parser define o contrato que todos os parsers devem seguir.
type Parser interface {
	Name() string
	Parse(path string) (*graph.InfraGraph, error)
}

// TerraformParser implementa a interface Parser para arquivos .tf
type TerraformParser struct{}

func NewParser() *TerraformParser {
	return &TerraformParser{}
}

func (p *TerraformParser) Name() string {
	return "terraform"
}

// Parse simula a análise de arquivos .tf e monta um InfraGraph básico.
func (p *TerraformParser) Parse(path string) (*graph.InfraGraph, error) {
	stat, err := os.Stat(path)
	if err != nil {
		return nil, fmt.Errorf("invalid path: %w", err)
	}
	if !stat.IsDir() {
		return nil, fmt.Errorf("only directories supported for now")
	}

	g := graph.NewInfraGraph()

	// 🚧 STUB: aqui ainda não fazemos parsing real dos arquivos .tf
	// Apenas adicionamos recursos simulados para testes
	g.AddResource(&graph.ResourceNode{
		ID:         "aws_s3_bucket.logs",
		Type:       "aws_s3_bucket",
		Name:       "logs",
		Attributes: map[string]interface{}{"acl": "private"},
		Tags:       map[string]string{"project": "tfmap"},
		SourceFile: filepath.Join(path, "main.tf"),
		LineNumber: 1,
		ModulePath: "root",
	})

	return g, nil
}
