package utils

import (
	"github.com/Haepapa/sas-lineage/internal/types"
	"github.com/google/uuid"
)

func GetOrCreateNodeID(nodes *[]types.Node, node types.Node) string {
	for _, n := range *nodes {
		if n.Label == node.Label && n.Name == node.Name {
			return n.ID
		}
	}
	node.ID = uuid.New().String()
	*nodes = append(*nodes, node)
	return node.ID
}

func AppendUniqueLink(links *[]types.Link, link types.Link) {
	for _, l := range *links {
		if l.Source == link.Source && l.Target == link.Target && l.Label == link.Label {
			return
		}
	}
	*links = append(*links, link)
}