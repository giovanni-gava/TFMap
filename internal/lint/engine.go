package lint

import (
	"github.com/giovanni-gava/tfmap/internal/graph"
)

// RuleFunc define a assinatura de uma regra de lint.
// Cada regra recebe um grafo e retorna uma lista de resultados.
type RuleFunc func(*graph.InfraGraph) []LintResult

// RunAll executa todas as regras de lint e retorna os resultados combinados.
func RunAll(g *graph.InfraGraph) []LintResult {
	var results []LintResult

	for _, rule := range rules() {
		results = append(results, rule(g)...)
	}

	return results
}

// rules retorna a lista de funções de regras que devem ser aplicadas.
// Cada função deve ser modular, testável e independente.
func rules() []RuleFunc {
	return []RuleFunc{
		checkMissingTags,
		// futuras regras:
		// checkWildcardIAM,
		// checkImplicitDependencies,
	}
}

