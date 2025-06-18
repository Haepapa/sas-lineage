package utils

import (
	"os"
	"path/filepath"
	"strings"
)

func FindSASFiles(root string) ([]string, error) {
    var files []string
    err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
        if err != nil {
            return err
        }
        if !info.IsDir() && (strings.HasSuffix(path, ".sas") || strings.HasSuffix(path, ".egp")) {
            files = append(files, path)
        }
        return nil
    })
    return files, err
}