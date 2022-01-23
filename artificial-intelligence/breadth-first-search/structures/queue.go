package structures

func NewQueue() *Queue {
	return &Queue{}
}

type Queue []string

func (q *Queue) Push(data string) {
	*q = append(*q, data)
}

func (q *Queue) Pop() *string {
	data := (*q)[0]
	*q = (*q)[1:]
	return &data
}

func (q *Queue) Count() int {
	return len(*q)
}
