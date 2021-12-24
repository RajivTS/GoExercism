package tree

import "errors"

// Define the Record type
type Record struct {
	ID     int
	Parent int
}

// Define the Node type
type Node struct {
	ID       int
	Children []*Node
}

var errOutOfRangeNode error = errors.New("node ID out of range")
var errOutOfOrderOrCyclicNode error = errors.New("nodes out of order or cyclic")
var errDisconnectedTree error = errors.New("the directory tree is disconnected")
var errDuplicateNodes error = errors.New("duplicate nodes in the directory tree")

func Build(records []Record) (*Node, error) {
	nodeCount := 0
	nodeLst := make([]*Node, len(records))
	parentMap := make(map[int]int)
	for _, record := range records {
		if record.ID < 0 || record.ID >= len(records) || record.Parent < 0 || record.Parent >= len(records) {
			return nil, errOutOfRangeNode
		}
		if record.ID < record.Parent || record.ID == record.Parent && record.ID != 0 {
			return nil, errOutOfOrderOrCyclicNode
		}
		if nodeLst[record.ID] == nil {
			nodeLst[record.ID] = &Node{ID: record.ID}
			nodeCount++
		} else {
			return nil, errDuplicateNodes
		}
		parentMap[record.ID] = record.Parent
	}
	if nodeCount != len(records) {
		return nil, errDisconnectedTree
	}
	for idx := 0; idx < len(records); idx++ {
		childNode, parentNode := nodeLst[idx], nodeLst[parentMap[idx]]
		if parentNode != childNode {
			parentNode.Children = append(parentNode.Children, childNode)
		}
	}
	var rootNode *Node = nil
	if nodeCount > 0 {
		rootNode = nodeLst[0]
	}
	return rootNode, nil
}
