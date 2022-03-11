package tree

import (
  "fmt"
)

type Record struct {
  ID     int
  Parent int
}

type Node struct {
  ID       int
  Children []*Node
}

// Build receive slice of Record with process information and add respective Node
func Build(r []Record) (*Node, error) {

  // encerra em casos de records vazios
  if len(r) == 0 {
    return nil, nil
  }

  // se houve somente o processo root
  if len(r) == 1 {
    return &Node{}, nil
  }

  // tratamento outros processo
  // n recebe o inicio do processo
  n := Node{0, nil}

  // aplica os processo no root process
  for i, v := range r {
    // Sabemos que temos 2 processos para registrar
    fmt.Println("Trabalhar com:", len(r) - 1)
    fmt.Println("root", n.Children)
    if v.ID != 0 {
      n.Children = append(n.Children, &Node{v.ID, nil})
    }
    fmt.Println(i, v)
  }
  return &n, nil
}

const rootID = 0
func Solved(records []Record) (*Node, error) {
  if len(records) == 0 {
    return nil, nil
  }
  positions := make([]int, len(records))
  for i, r := range records {
    if r.ID < rootID || r.ID >= len(records) {
      return nil, fmt.Errorf("out of bounds record id %d", r.ID)
    }
    positions[r.ID] = i
  }
  nodes := make([]Node, len(records))
  for i := range positions {
    r := records[positions[i]]
    if r.ID != i {
      return nil, fmt.Errorf("non-contiguous node %d (want %d)", r.ID, i)
    }
    validParentForChild := (r.ID > r.Parent) || (r.ID == rootID && r.Parent == rootID)
    if !validParentForChild {
      return nil, fmt.Errorf("node %d has impossible parent %d", r.ID, r.Parent)
    }
    nodes[i].ID = i
    if i != rootID {
      p := &nodes[r.Parent]
      p.Children = append(p.Children, &nodes[i])
    }
  }
  return &nodes[0], nil
}
