package types

type Node struct {
    ID       string `json:"id"`
    Label     string `json:"label"` // e.g., "script", "dataset"
    ClassName     string `json:"className"`
    Shape string `json:"shape"`
    SizeX string `json:"sizeX"`
    SizeY string `json:"sizeY"`
    X string `json:"x"`
    Y string `json:"y"`
}

type Link struct {
    ID       string `json:"id"`
    Label     string `json:"label"` // e.g., "script", "dataset"
    ClassName     string `json:"className"`
    Direction     string `json:"direction"`
    LeftNodeID     string `json:"leftNodeID"`
    RightNodeID    string `json:"rightNodeID"`
}