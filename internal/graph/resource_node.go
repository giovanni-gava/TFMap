package graph

// ResourceNode representa um recurso da infraestrutura como código.
type ResourceNode struct {
	ID         string                 // ex: aws_instance.web
	Type       string                 // ex: aws_instance
	Name       string                 // ex: web
	Provider   string                 // ex: aws
	Attributes map[string]interface{} // atributos brutos do recurso (ex: ami, instance_type, tags...)
	DependsOn  []string               // IDs de recursos dos quais este depende
	ModulePath string                 // caminho lógico do módulo (ex: root.module.vpc.module.subnet)
	SourceFile string                 // caminho do arquivo onde o recurso foi declarado
	LineNumber int                    // linha onde o recurso começa (opcional, para UI/debug)
	Tags       map[string]string      // tags detectadas (separado de Attributes para facilitar linting)
	Metadata   ResourceMetadata       // informações adicionais enriquecidas
}

// Edge representa uma ligação entre dois nós no grafo.
type Edge struct {
	From string // ID do recurso origem
	To   string // ID do recurso destino
	Type string // "explicit", "implicit", "output", "module", "computed", etc
}

// InfraMetadata representa informações de alto nível sobre a infra analisada.
type InfraMetadata struct {
	Project     string   // ex: my-fintech-app
	Workspace   string   // ex: dev, prod, staging
	Environment string   // ex: dev, qa, prod
	SourcePaths []string // arquivos analisados
	CreatedBy   string   // username ou processo
	Timestamp   string   // ISO8601 timestamp da análise
	Version     string   // versão do TFMap usada na análise
}

// ResourceMetadata representa metadados adicionais para rastreabilidade, debug ou exportação.
type ResourceMetadata struct {
	ManagedBy     string   // terraform, terragrunt, manual, etc
	SourceVersion string   // hash ou versão do commit
	Warnings      []string // mensagens detectadas no parsing
	Validated     bool     // se passou por análise semântica
	Drifted       bool     // se está diferente do estado real
}
