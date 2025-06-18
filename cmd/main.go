package main

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/Haepapa/sas-lineage/internal/model"
	"github.com/Haepapa/sas-lineage/internal/parser/egp"
	"github.com/Haepapa/sas-lineage/internal/parser/sas"
	"github.com/Haepapa/sas-lineage/internal/utils/exporter"
	"github.com/Haepapa/sas-lineage/internal/utils/filewalker"
	"github.com/google/uuid"
)

func main() {
    root := os.Args[1]
    paths, _ := filewalker.FindSASFiles(root)
    var nodes []model.Node
    var links []model.Link

    for _, p := range paths {
        if strings.HasSuffix(p, ".egp") {
            temp, _ := egp.ExtractEGP(p)
            // Process extracted files recursively
        } else if strings.HasSuffix(p, ".sas") {
            inputs, outputs := sas.ParseSASCode(p)
            // Add node for the script
            scriptNode := model.Node{ID: uuid.New().String(), Type: "script", Name: filepath.Base(p), Location: p}
            nodes = append(nodes, scriptNode)

            for _, in := range inputs {
                inNode := model.Node{ID: in, Type: "dataset", Name: in, Location: ""}
                nodes = append(nodes, inNode)
                links = append(links, model.Link{Source: inNode.ID, Target: scriptNode.ID, Type: "reads"})
            }
            for _, out := range outputs {
                outNode := model.Node{ID: out, Type: "dataset", Name: out, Location: ""}
                nodes = append(nodes, outNode)
                links = append(links, model.Link{Source: scriptNode.ID, Target: outNode.ID, Type: "writes"})
            }
        }
    }

    exporter.ExportLineage(nodes, links, "./output")
}