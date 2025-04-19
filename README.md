# TFMap – IaC Intelligence Engine

TFMap is a GitOps-ready, open-source engine for intelligent analysis and visualization of infrastructure as code.

Built for DevOps, SRE, SecOps, and Platform Engineers who need clarity, security, and architecture-level insights from Terraform and Terragrunt configurations.

## 🌍 What it does

- Parses `.tf`, `.tf.json`, `terragrunt.hcl`, `tfstate`, and `tfplan`
- Builds a semantic graph of resources, modules, outputs, and dependencies
- Detects risks (IAM wildcards, missing tags, implicit dependencies)
- Exports to `.json`, `.dot`, `.md`, `.svg`
- GitHub/GitLab CI-ready
- CLI-first, API-ready, UI-compatible

## 📦 Getting Started

```bash
go install github.com/giovanni-gava/tfmap/cmd/tfmap@latest

tfmap parse --input ./infra --format json --output ./tfmap.json


💡 Why TFMap?
Terraform shows what you wrote.
TFMap shows what you built — and where it could fail.

📈 Roadmap
✅ CLI-based parser and graph exporter

🛠️ Terraform + Terragrunt cross-parsing

🔒 Security linting (IAM, exposure, drift)

🌐 REST API (optional, local or SaaS-ready)

🖥️ Web UI (zoomable, filterable graph explorer)

🤖 GitHub Bot for IaC Pull Requests

📊 Audit reports and compliance artifacts

🧠 License
Apache 2.0

🗣️ Contribute
Open an issue, submit a PR or start a discussion — all feedback welcome.
Let’s build the future of infrastructure understanding.

Built with purpose. Built for teams. Built for insight.
— The TFMap Project.