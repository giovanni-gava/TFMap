package lint

type LintLevel string

const (
	LevelError   LintLevel = "error"
	LevelWarning LintLevel = "warning"
	LevelInfo    LintLevel = "info"
	LevelHint    LintLevel = "hint"
)

type LintResult struct {
	ResourceID string     // ID completo do recurso (ex: aws_s3_bucket.example)
	Level      LintLevel  // error, warning, info, hint
	Rule       string     // Código ou nome da regra (ex: no_tags, wildcard_iam)
	Message    string     // Mensagem explicativa
	Suggestion string     // Dica prática ou documentação recomendada
	File       string     // Caminho do arquivo onde o recurso foi definido (se disponível)
	Line       int        // Linha de origem (opcional)
}

