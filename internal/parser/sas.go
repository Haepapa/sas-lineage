package parser

import (
	"os"
	"path/filepath"
	"regexp"

	"github.com/Haepapa/sas-lineage/internal/types"
	"github.com/Haepapa/sas-lineage/internal/utils"
	"github.com/google/uuid"
)

var inputRe = regexp.MustCompile(`(?i)set\s+([a-zA-Z0-9_.]+)`)
var outputRe = regexp.MustCompile(`(?i)data\s+([a-zA-Z0-9_.]+)`)

func ParseSASCode(path string, nodes *[]types.Node, links *[]types.Link) error {
    b, err := os.ReadFile(path)
    if err != nil {
        return err
    }
    text := string(b)
    var inputs []string
    var outputs []string
    for _, match := range inputRe.FindAllStringSubmatch(text, -1) {
        inputs = append(inputs, match[1])
    }
    for _, match := range outputRe.FindAllStringSubmatch(text, -1) {
        outputs = append(outputs, match[1])
    }
    scriptID := utils.GetOrCreateNodeID(nodes, types.Node{
        Label:     "script",
        ClassName: filepath.Base(path),
        Shape:     "rectangle",
        SizeX:     "100",
        SizeY:     "60",
        X:         "150",
        Y:         "150",
    })
    for _, in := range inputs {
        dataID := utils.GetOrCreateNodeID(nodes, types.Node{
            Label:     "dataset",
            ClassName: in,
            Shape:     "circle",
            SizeX:     "80",
            SizeY:     "80",
            X:         "150",
            Y:         "150",
        })
        utils.AppendUniqueLink(links, types.Link{
            ID:          uuid.New().String(),
            Label:       "reads",
            ClassName:   "",
            Direction:   "left-to-right",
            LeftNodeID:  dataID,
            RightNodeID: scriptID,
        })
    }
    for _, out := range outputs {
        dataID := utils.GetOrCreateNodeID(nodes, types.Node{
            Label:     "dataset",
            ClassName: out,
            Shape:     "circle",
            SizeX:     "80",
            SizeY:     "80",
            X:         "150",
            Y:         "150",
        })
        utils.AppendUniqueLink(links, types.Link{
            ID:          uuid.New().String(),
            Label:       "writes",
            ClassName:   "",
            Direction:   "left-to-right",
            LeftNodeID:  scriptID,
            RightNodeID: dataID,
        })
    }
    return nil
}