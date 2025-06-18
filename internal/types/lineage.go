package types

type Node struct {
    ID       string `json:"id"`
    Type     string `json:"type"` // e.g., "script", "dataset"
    Name     string `json:"name"`
    Location string `json:"location"`
}

type Link struct {
    Source string `json:"source"` // Node ID
    Target string `json:"target"` // Node ID
    Type   string `json:"type"`   // e.g., "reads", "writes"
}