package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/giovanni-gava/tfmap/internal/exporter/dot"
	"github.com/giovanni-gava/tfmap/internal/parser/terraform"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "tfmap",
		Usage: "Parse and visualize Terraform/Terragrunt infrastructure as code",
		Commands: []*cli.Command{
			{
				Name:  "parse",
				Usage: "Parse IaC and export the infrastructure graph",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "input",
						Usage:    "Path to IaC files or directory",
						Required: true,
					},
					&cli.StringFlag{
						Name:     "format",
						Usage:    "Export format (json, dot)",
						Required: true,
					},
					&cli.StringFlag{
						Name:     "output",
						Usage:    "Path to output file",
						Required: true,
					},
				},
				Action: func(c *cli.Context) error {
					inputPath := c.String("input")
					format := c.String("format")
					outputPath := c.String("output")

					absPath, err := filepath.Abs(inputPath)
					if err != nil {
						return fmt.Errorf("failed to resolve input path: %w", err)
					}

					fmt.Printf("🔍 Parsing infrastructure from %s...\n", absPath)

					parser := terraform.NewParser()
					infraGraph, err := parser.Parse(absPath)
					if err != nil {
						return fmt.Errorf("failed to parse infrastructure: %w", err)
					}

					switch format {
					case "json":
						file, err := os.Create(outputPath)
						if err != nil {
							return fmt.Errorf("failed to create output file: %w", err)
						}
						defer file.Close()

						encoder := json.NewEncoder(file)
						encoder.SetIndent("", "  ")
						if err := encoder.Encode(infraGraph); err != nil {
							return fmt.Errorf("failed to write graph: %w", err)
						}
						fmt.Printf("✅ InfraGraph exported to %s (JSON format)\n", outputPath)

					case "dot":
						exporter := dot.NewExporter()
						if err := exporter.Export(infraGraph, outputPath); err != nil {
							return fmt.Errorf("failed to export DOT: %w", err)
						}
						fmt.Printf("✅ InfraGraph exported to %s (DOT format)\n", outputPath)

					default:
						return fmt.Errorf("format %s not supported yet", format)
					}

					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatalf("❌ %v\n", err)
	}
}
