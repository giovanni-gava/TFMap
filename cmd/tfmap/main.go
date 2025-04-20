package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/giovanni-gava/tfmap/internal/exporter/dot"
	"github.com/giovanni-gava/tfmap/internal/lint"
	"github.com/giovanni-gava/tfmap/internal/parser/terraform"
	"github.com/urfave/cli/v2"
)

var Version = "dev"

func main() {
	app := &cli.App{
		Name:    "tfmap",
		Usage:   "Parse and visualize Terraform/Terragrunt infrastructure as code",
		Version: Version,
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
			{
				Name:  "lint",
				Usage: "Analyze IaC and detect missing tags or bad practices",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "input",
						Usage:    "Path to IaC files or directory",
						Required: true,
					},
					&cli.StringFlag{
						Name:  "format",
						Usage: "Output format (default: human, options: json)",
						Value: "human",
					},
					&cli.StringFlag{
						Name:  "output",
						Usage: "Path to output file (optional)",
					},
					&cli.BoolFlag{
						Name:  "strict",
						Usage: "Exit with code 1 if warnings or errors are found",
					},
				},
				Action: func(c *cli.Context) error {
					inputPath := c.String("input")
					format := c.String("format")
					outputPath := c.String("output")
					strict := c.Bool("strict")

					absPath, err := filepath.Abs(inputPath)
					if err != nil {
						return fmt.Errorf("failed to resolve input path: %w", err)
					}

					fmt.Printf("🔍 Linting infrastructure from %s...\n", absPath)

					parser := terraform.NewParser()
					graph, err := parser.Parse(absPath)
					if err != nil {
						return fmt.Errorf("failed to parse IaC: %w", err)
					}

					results := lint.RunAll(graph)

					// FORMAT: JSON
					if format == "json" {
						jsonData, err := json.MarshalIndent(results, "", "  ")
						if err != nil {
							return fmt.Errorf("failed to marshal JSON: %w", err)
						}

						if outputPath != "" {
							if err := os.WriteFile(outputPath, jsonData, 0644); err != nil {
								return fmt.Errorf("failed to write JSON to file: %w", err)
							}
							fmt.Printf("✅ JSON lint report written to %s\n", outputPath)
						} else {
							fmt.Println(string(jsonData))
						}

						if strict {
							for _, r := range results {
								if r.Level == lint.LevelWarning || r.Level == lint.LevelError {
									fmt.Println("❌ Issues detected in strict mode.")
									os.Exit(1)
								}
							}
						}

						return nil
					}

					// FORMAT: HUMAN-READABLE
					if len(results) == 0 {
						fmt.Println("✅ No lint issues found.")
						return nil
					}

					for _, r := range results {
						prefix := "[INFO]"
						if r.Level == lint.LevelWarning {
							prefix = "[WARNING]"
						} else if r.Level == lint.LevelError {
							prefix = "[ERROR]"
						}

						fmt.Printf("%s %s (%s)\n", prefix, r.ResourceID, r.Rule)
						fmt.Printf("    › %s\n", r.Message)
						if r.Suggestion != "" {
							fmt.Printf("    ↳ Suggestion: %s\n", r.Suggestion)
						}
						if r.File != "" {
							fmt.Printf("    ↳ Location: %s:%d\n", r.File, r.Line)
						}
					}

					if strict {
						for _, r := range results {
							if r.Level == lint.LevelWarning || r.Level == lint.LevelError {
								fmt.Println("❌ Issues detected in strict mode.")
								os.Exit(1)
							}
						}
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
