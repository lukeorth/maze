package main

import "errors"

type Stack struct {
    cell []Cell
}

func NewStack() *Stack {
    return &Stack{make([]Cell, 0)}
}

func (s *Stack) Push(v *Cell) {
    s.cell = append(s.cell, *v)
}

func (s *Stack) Pop() (*Cell, error) {
    l := len(s.cell)
    if l == 0 {
        return nil, errors.New("Empty Stack")
    }

    res := s.cell[l-1]
    s.cell = s.cell[:l-1]
    return &res, nil
}
