# TFMap – IaC Intelligence Engine

**Created by [Giovanni Colognesi](https://github.com/giovanni-gava)**  
Infrastructure Auditor • DevOps Strategist • Open Source Builder

---

TFMap is an open-source, GitOps-ready platform for intelligent analysis and visual representation of infrastructure as code (IaC). Built to empower DevOps, SRE, Platform Engineering, and Security teams to deeply understand and audit their Terraform and Terragrunt-based infrastructure.

> "Infrastructure isn't just code — it's a system of truth. TFMap is the engine that reveals it."

---

## ✅ Core Features (Delivered)

- 🔍 **Deep parser** for `.tf` files using `hcl/v2`
- 🧠 Internal `InfraGraph` model with nodes, edges, and metadata
- 📦 Modular CLI with subcommands and flags
- 🎨 Exporter `.dot` for Graphviz visualizations
- 🧪 Unit tests for parser and exporter

---

## 🚧 Roadmap (Fase 1 – Core Engine & CLI)

| Fase 1 Microetapas                   | Status     |
|--------------------------------------|------------|
| Parser real de `.tf` com HCL         | ✅ Pronto  |
| CLI `parse` funcional                | ✅ Pronto  |
| Exportador `.json` e `.dot`          | ✅ Pronto  |
| Linter de boas práticas (`tags`)     | 🚧 Em breve |
| Exportador `.md`, `.svg`, `.yaml`    | 🔜          |
| Documentação de arquitetura          | 🔜          |

---

## 📦 Installation

```bash
go install github.com/giovanni-gava/tfmap/cmd/tfmap@latest
```

---

## 🚀 Usage

```bash
tfmap parse --input ./infra --format dot --output graph.dot
dot -Tsvg graph.dot -o graph.svg
```

---

## 📊 Example Output

![graph example](./graph.svg)

---

## 🔧 Architecture Overview

```
tfmap/
├── cmd/              # CLI entrypoint
├── internal/
│   ├── parser/       # Terraform, Terragrunt, tfplan, tfstate
│   ├── graph/        # InfraGraph model
│   ├── exporter/     # dot, json, yaml
│   ├── lint/         # rules, engine, result types
├── testdata/         # test infrastructure
```

---

## 🧠 Why TFMap?

Terraform tells you what you wrote.
TFMap tells you what you built — and where you're exposed.

---

## 💬 Contribute

TFMap is in active development.
PRs, issues, feedback, and ideas are welcome — join the mission.

---

## 📜 License

Apache 2.0 License — use freely and responsibly.

---

**Made with purpose by [Giovanni Colognesi](https://linkedin.com/in/giovanni-gava-21338115a)**  
Let’s turn infrastructure into knowledge.
