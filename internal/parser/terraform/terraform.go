package terraform

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/giovanni-gava/tfmap/internal/graph"
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclparse"
)

type TerraformParser struct{}

func NewParser() *TerraformParser {
	return &TerraformParser{}
}

func (p *TerraformParser) Name() string {
	return "terraform"
}

func (p *TerraformParser) Parse(path string) (*graph.InfraGraph, error) {
	info, err := os.Stat(path)
	if err != nil {
		return nil, fmt.Errorf("invalid path: %w", err)
	}
	if !info.IsDir() {
		return nil, fmt.Errorf("only directories are supported")
	}

	parser := hclparse.NewParser()
	infra := graph.NewInfraGraph()

	err = filepath.WalkDir(path, func(fpath string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() || !strings.HasSuffix(fpath, ".tf") {
			return nil
		}

		file, diag := parser.ParseHCLFile(fpath)
		if diag.HasErrors() {
			return fmt.Errorf("failed to parse %s: %s", fpath, diag.Error())
		}

		body := file.Body
		content, _, diag := body.PartialContent(&hcl.BodySchema{
			Blocks: []hcl.BlockHeaderSchema{
				{
					Type:       "resource",
					LabelNames: []string{"type", "name"}},
			},
		})
		if diag.HasErrors() {
			return fmt.Errorf("failed to read blocks in %s: %s", fpath, diag.Error())
		}

		for _, block := range content.Blocks {
			if block.Type != "resource" || len(block.Labels) != 2 {
				continue
			}

			resType := block.Labels[0]
			resName := block.Labels[1]
			id := fmt.Sprintf("%s.%s", resType, resName)

			attrs, diag := block.Body.JustAttributes()
			if diag.HasErrors() {
				return fmt.Errorf("failed to extract attributes: %s", diag.Error())
			}

			attrMap := make(map[string]interface{})
			tagMap := make(map[string]string)

			for name, attr := range attrs {
				val, diag := attr.Expr.Value(nil)
				if diag.HasErrors() {
					continue // silenciosamente ignora erros de eval por enquanto
				}

				switch val.Type().FriendlyName() {
				case "string":
					attrMap[name] = val.AsString()
				case "object":
					if name == "tags" {
						for tagKey, tagVal := range val.AsValueMap() {
							if tagVal.Type().FriendlyName() == "string" {
								tagMap[tagKey] = tagVal.AsString()
							}
						}
					}
				default:
					attrMap[name] = val.GoString()
				}
			}

			infra.AddResource(&graph.ResourceNode{
				ID:         id,
				Type:       resType,
				Name:       resName,
				Attributes: attrMap,
				Tags:       tagMap,
				SourceFile: fpath,
				ModulePath: "root", // placeholder por enquanto
			})
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return infra, nil
}
