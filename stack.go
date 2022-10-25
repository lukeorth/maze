package main

import "errors"

type Stack struct {
    cells []Cell
}

func NewStack() *Stack {
    return &Stack{make([]Cell, 0)}
}

func (s *Stack) Push(v *Cell) {
    s.cells = append(s.cells, *v)
}

func (s *Stack) Pop() (*Cell, error) {
    l := len(s.cells)
    if l == 0 {
        return nil, errors.New("Empty Stack")
    }

    res := s.cells[l-1]
    s.cells = s.cells[:l-1]
    return &res, nil
}
