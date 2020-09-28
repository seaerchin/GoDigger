package GoDigger

import gn "gonum.org/v1/gonum/graph"

// Modularity calculates the modularity for the goodness of a clustering
// refer to https://github.com/gonum/gonum/blob/master/graph/community/louvain_undirected.go
func Modularity(g gn.UndirectWeighted, communities []gn.Node) int {
	gEdges := NumEdges(g)
}


// NumEdges does a dfs traversal through g and returns the number of eges
func NumEdges(graph gn.Graph) int { 
	return -1 
}