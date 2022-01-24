package structures

type Stack []string

func NewStack() *Stack {
	return &Stack{}
}

func (s *Stack) Push(data string) {
	*s = append(*s, data)
}

func (s *Stack) Pop() *string {
	if s.Count() == 0 {
		return nil
	}
	i := s.Count() - 1
	data := (*s)[i]
	*s = (*s)[:i]
	return &data
}

func (s *Stack) Count() int {
	return len(*s)
}
