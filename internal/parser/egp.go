package parser

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/Haepapa/sas-lineage/internal/types"
	"github.com/Haepapa/sas-lineage/internal/utils"
)
func ExtractEGP(path string, baseTempDir string, nodes *[]types.Node, links *[]types.Link) error {
	debug := false
	tempDir, err := os.MkdirTemp(baseTempDir, "egp-*")
	if err != nil {
		return err
	}
	r, err := zip.OpenReader(path)
	if err != nil {
		return err
	}
	defer r.Close()
	for _, f := range r.File {
		outPath := filepath.Join(tempDir, f.Name)
		if err := func() error {
			if err := os.MkdirAll(filepath.Dir(outPath), 0755); err != nil {
				if !strings.Contains(err.Error(), "not a directory") {
					return err
				}
				if debug {
					fmt.Printf("Warning: Skipping file due to directory error: %s\n", err)
				}
				return nil
			}
			rc, err := f.Open()
			if err != nil {
				if !strings.Contains(err.Error(), "not a directory") {
					return err
				}
				if debug {
					fmt.Printf("Warning: Skipping file due to directory error: %s\n", err)
				}
				return nil
			}
			defer rc.Close()
			outFile, err := os.Create(outPath)
			if err != nil {
				if !strings.Contains(err.Error(), "not a directory") {
					return err
				}
				if debug {
					fmt.Printf("Warning: Skipping file due to directory error: %s\n", err)
				}
				return nil
			}
			defer outFile.Close()
			_, err = io.Copy(outFile, rc)
			return err
		}(); err != nil {
			return err
		}
	}
	subPaths, _ := utils.FindSASFiles(tempDir)
	for _, sp := range subPaths {
		if strings.HasSuffix(sp, ".sas") {
			err := ParseSASCode(sp, nodes, links)
			if err != nil {
				return err
			}
		}
	}
	if !debug {
		os.RemoveAll(tempDir)
	}
	return nil
}