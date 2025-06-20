package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/Haepapa/sas-lineage/internal/parser"
	"github.com/Haepapa/sas-lineage/internal/types"
	"github.com/Haepapa/sas-lineage/internal/utils"
)

func main() {
    var nodes []types.Node
    var links []types.Link
	var tempDir string
	flag.StringVar(&tempDir, "temp-dir", "", "Base directory to extract temporary EGP contents")
	flag.StringVar(&tempDir, "t", "", "Base directory to extract temporary EGP contents (shorthand)")
	flag.Parse()

	if flag.NArg() < 1 {
		fmt.Println("Usage: sas-lineage [--temp-dir path] root_path")
		os.Exit(1)
	}

	root := flag.Arg(0)
	if tempDir == "" {
		tempDir = root
	}
	paths, _ := utils.FindSASFiles(root)

	for _, p := range paths {
		if strings.HasSuffix(p, ".egp") {
			log.Printf("Processing EGP file: %s\n", p)
			if err := parser.ExtractEGP(p, tempDir, &nodes, &links); err != nil {
				log.Printf("Error extracting EGP file: %s\n", err)
				continue
			}
		} else if strings.HasSuffix(p, ".sas") {
			log.Printf("Processing SAS file: %s\n\n", p)
			if err := parser.ParseSASCode(p, &nodes, &links, ""); err != nil {
				log.Printf("Error parsing SAS file: %s\n", err)
				continue
			}
		}
	}
    utils.ExportLineage(nodes, links, "./output")
}