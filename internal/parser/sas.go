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
        Label:     func() string {
            if sasEGName != "" {
                return sasEGName
            }
            return path
        }(),
        ClassName: func() string {
            if sasEGName != "" {
                return "sas-egp"
            }
            return "sas-script"
        }(),
        Shape:     "rectangle",
        SizeX:     100,
        SizeY:     60,
        X:         150,
        Y:         150,
        Type: func() string {
            if sasEGName != "" {
                return "SAS Enterprise Guide"
            }
            return "SAS Program"
        }(),
    })
    for _, in := range inputs {
        dataID := utils.GetOrCreateNodeID(nodes, types.Node{
            Label:     in,
            ClassName: "sas-dataset",
            Shape:     "circle",
            SizeX:     80,
            SizeY:     80,
            X:         150,
            Y:         150,
            Type:      "SAS Dataset",
        })
        utils.AppendUniqueLink(links, types.Link{
            ID:          uuid.New().String(),
            Label:       "reads",
            ClassName:   "reads",
            Direction:   "left-to-right",
            LeftNodeId:  dataID,
            RightNodeId: scriptID,
        })
    }
    for _, out := range outputs {
        dataID := utils.GetOrCreateNodeID(nodes, types.Node{
            Label:     out,
            ClassName: "sas-dataset",
            Shape:     "circle",
            SizeX:     80,
            SizeY:     80,
            X:         150,
            Y:         150,
            Type:      "SAS Dataset",
        })
        utils.AppendUniqueLink(links, types.Link{
            ID:          uuid.New().String(),
            Label:       "writes",
            ClassName:   "writes",
            Direction:   "left-to-right",
            LeftNodeId:  scriptID,
            RightNodeId: dataID,
        })
    }
    return nil
}