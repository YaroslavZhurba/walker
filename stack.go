package main

type StackDir struct {
	slice []string
	len   int
}

func NewStackDir() *StackDir {
	return &StackDir{
		slice: make([]string, 0),
		len:   0,
	}
}

func (s *StackDir) Push(f string) {
	s.len++
	if len(s.slice) < s.len {
		s.slice = append(s.slice, f)
	}
	s.slice[s.len-1] = f
}

func (s *StackDir) Pop() string {
	if s.len == 0 {
		panic("Pop from empty string")
	}
	s.len--
	return s.slice[s.len]
}

func (s *StackDir) IsEmpty() bool {
	return s.len == 0
}
