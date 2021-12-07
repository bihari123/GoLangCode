package tree
import "fmt"
// Define the Record type

const rootID = 0

type Record struct {
    ID, Parent int
}

type Node struct {
    ID       int
    Children []*Node
}
// Build converts an unordered slice of Records into a hierarchical tree of Nodes,

// after validating that the tree is not a graph and that the records have

// a contiguous range of IDs.

func Build(records []Record) (*Node, error) {

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
