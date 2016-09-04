package astar

type Item struct {
	ID    int
	Cost  float64
	Score float64
	Depth int
	Next  *Item
}

func NewItem(id int) *Item {
	return &Item{ID: id}
}

func (item *Item) Follow(edge Edge, estimate float64) *Item {
	result := Item{}
	result.ID = edge.Dst
	result.Cost = item.Cost + edge.Cost
	result.Score = result.Cost + estimate
	result.Depth = item.Depth + 1
	result.Next = item
	return &result
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].Score < pq[j].Score
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(*Item))
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[:n-1]
	return item
}
