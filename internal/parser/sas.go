package parser

import (
	"os"
	"regexp"
)

var inputRe = regexp.MustCompile(`(?i)set\s+([a-zA-Z0-9_.]+)`)
var outputRe = regexp.MustCompile(`(?i)data\s+([a-zA-Z0-9_.]+)`)

func ParseSASCode(path string) (inputs, outputs []string) {
    b, err := os.ReadFile(path)
    if err != nil {
        return
    }
    text := string(b)
    for _, match := range inputRe.FindAllStringSubmatch(text, -1) {
        inputs = append(inputs, match[1])
    }
    for _, match := range outputRe.FindAllStringSubmatch(text, -1) {
        outputs = append(outputs, match[1])
    }
    return
}