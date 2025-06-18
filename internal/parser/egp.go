package parser

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"
	"strings"
)
func ExtractEGP(path string) (string, error) {
    tempDir := strings.TrimSuffix(path, ".egp")
    r, err := zip.OpenReader(path)
    if err != nil {
        return "", err
    }
    defer r.Close()
    for _, f := range r.File {
        outPath := filepath.Join(tempDir, f.Name)
        if err := os.MkdirAll(filepath.Dir(outPath), 0755); err != nil {
            return "", err
        }
        rc, _ := f.Open()
        outFile, _ := os.Create(outPath)
        io.Copy(outFile, rc)
        rc.Close()
        outFile.Close()
    }
    return tempDir, nil
}