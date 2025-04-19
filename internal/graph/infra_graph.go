package graph

// InfraGraph is the central structure holding the entire infrastructure graph
type InfraGraph struct {
	Resources map[string]*ResourceNode // keyed by resource ID
	Edges     []*Edge
	Metadata  *InfraMetadata
}

func NewInfraGraph() *InfraGraph {
	return &InfraGraph{
		Resources: make(map[string]*ResourceNode),
		Edges:     []*Edge{},
		Metadata:  &InfraMetadata{},
	}
}

func (g *InfraGraph) AddResource(node *ResourceNode) {
	g.Resources[node.ID] = node
}

func (g *InfraGraph) AddEdge(from, to, edgeType string) {
	edge := &Edge{
		From: from,
		To:   to,
		Type: edgeType,
	}
	g.Edges = append(g.Edges, edge)
}

func (g *InfraGraph) GetResourceByID(id string) (*ResourceNode, bool) {
	n, ok := g.Resources[id]
	return n, ok
}
