package utils

import (
	"encoding/json"
	"os"
	"path/filepath"
)

func ExportLineage(nodes []model.Node, links []model.Link, outDir string) error {
    nData, _ := json.MarshalIndent(nodes, "", "  ")
    lData, _ := json.MarshalIndent(links, "", "  ")
    os.WriteFile(filepath.Join(outDir, "nodes.json"), nData, 0644)
    os.WriteFile(filepath.Join(outDir, "links.json"), lData, 0644)
    return nil
}