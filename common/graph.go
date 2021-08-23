package common

type Node struct {
	n string
}

type Graph struct {
	Edges map[string][]*Node
}

func NewGraph() *Graph {
	return &Graph{Edges: make(map[string][]*Node)}
}

// AddEdge 有向图添加边
func (g *Graph) AddEdge(from *Node, to *Node) {
	e, ok := g.Edges[from.n]
	if ok {
		g.Edges[from.n] = append(e, to)
	} else {
		g.Edges[from.n] = []*Node{to}
	}
}
