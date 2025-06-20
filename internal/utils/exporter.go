package utils

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/Haepapa/sas-lineage/internal/types"
)

func ExportLineage(nodes []types.Node, links []types.Link, outDir string) error {
	if err := os.MkdirAll(outDir, 0755); err != nil {
		return err
	}
    nData, _ := json.MarshalIndent(nodes, "", "  ")
    lData, _ := json.MarshalIndent(links, "", "  ")
    os.WriteFile(filepath.Join(outDir, "nodes.json"), nData, 0644)
    os.WriteFile(filepath.Join(outDir, "edges.json"), lData, 0644)
    return nil
}