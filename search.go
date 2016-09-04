package astar

import "container/heap"

type Graph interface {
	Edges(node int) []Edge
	Estimate(src, dst int) float64
}

type Edge struct {
	Dst  int
	Cost float64
}

func Search(graph Graph, src, dst int) Result {
	scores := make(map[int]float64)
	var queue PriorityQueue
	heap.Push(&queue, NewItem(src))
	for len(queue) > 0 {
		item := heap.Pop(&queue).(*Item)
		if item.ID == dst {
			return CreateResult(item)
		}
		if score, ok := scores[item.ID]; ok && score <= item.Score {
			continue
		}
		scores[item.ID] = item.Score
		for _, edge := range graph.Edges(item.ID) {
			estimate := graph.Estimate(edge.Dst, dst)
			newItem := item.Follow(edge, estimate)
			if score, ok := scores[newItem.ID]; ok && score <= newItem.Score {
				continue
			}
			heap.Push(&queue, newItem)
		}
	}
	return Result{}
}

type Result struct {
	Nodes []int
	Cost  float64
}

func CreateResult(item *Item) Result {
	cost := item.Cost
	path := make([]int, item.Depth+1)
	for item != nil {
		path[item.Depth] = item.ID
		item = item.Next
	}
	return Result{path, cost}
}
