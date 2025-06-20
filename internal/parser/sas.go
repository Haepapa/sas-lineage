package parser

import (
	"os"
	"regexp"
	"strings"

	"github.com/Haepapa/sas-lineage/internal/types"
	"github.com/Haepapa/sas-lineage/internal/utils"
	"github.com/google/uuid"
)

var inputRe = regexp.MustCompile(`(?i)set\s+([a-zA-Z0-9_.]+)`)
var outputRe = regexp.MustCompile(`(?i)data\s+([a-zA-Z0-9_.]+)`)

func ParseSASCode(path string, nodes *[]types.Node, links *[]types.Link, sasEGName string) error {
    b, err := os.ReadFile(path)
    if err != nil {
        return err
    }
    text := string(b)

    ignoreOutputNames := map[string]bool{
        "_null_": true,
    }

    blockCommentRe := regexp.MustCompile(`(?s)/\*.*?\*/`)
    text = blockCommentRe.ReplaceAllString(text, "")

    lineCommentRe := regexp.MustCompile(`(?m)^\s*\*.*?;`)
    text = lineCommentRe.ReplaceAllString(text, "")

    var inputs []string
    var outputs []string
    for _, match := range inputRe.FindAllStringSubmatch(text, -1) {
        inputs = append(inputs, match[1])
    }
    for _, match := range outputRe.FindAllStringSubmatch(text, -1) {
        name := strings.ToLower(match[1])
        if !ignoreOutputNames[name] {
            outputs = append(outputs, match[1])
        }
    }
    scriptID := utils.GetOrCreateNodeID(nodes, types.Node{
        Label: "label",
        Name:     func() string {
            if sasEGName != "" {
                return sasEGName
            }
            return path
        }(),
        Fill:     "#e66557",
        Size:     10,
        Type: func() string {
            if sasEGName != "" {
                return "SAS Enterprise Guide"
            }
            return "SAS Program"
        }(),
    })
    for _, in := range inputs {
        dataID := utils.GetOrCreateNodeID(nodes, types.Node{
            Label:     "sas-dataset",
            Name:     in,
            Fill:     "#5782e6",
            Size:     10,
            Type:      "SAS Dataset",
        })
        utils.AppendUniqueLink(links, types.Link{
            ID:          uuid.New().String(),
            Label:       "reads",
            Source:     dataID,
            Target:     scriptID,
        })
    }
    for _, out := range outputs {
        dataID := utils.GetOrCreateNodeID(nodes, types.Node{
            Label:     "sas-dataset",
            Name:     out,
            Fill:     "#5782e6",
            Size:     10,
            Type:      "SAS Dataset",
        })
        utils.AppendUniqueLink(links, types.Link{
            ID:          uuid.New().String(),
            Label:       "writes",
            Source:     dataID,
            Target:     scriptID,
        })
    }
    return nil
}