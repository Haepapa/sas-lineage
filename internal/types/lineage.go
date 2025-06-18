package types

type Node struct {
    ID       string `json:"id"`
    Label     string `json:"label"` // e.g., "script", "dataset"
    ClassName     string `json:"className"`
    Shape string `json:"shape"`
    SizeX int `json:"sizeX"`
    SizeY int `json:"sizeY"`
    X int `json:"x"`
    Y int `json:"y"`
}

type Link struct {
    ID       string `json:"id"`
    Label     string `json:"label"` // e.g., "script", "dataset"
    ClassName     string `json:"className"`
    Direction     string `json:"direction"`
    LeftNodeId     string `json:"leftNodeId"`
    RightNodeId    string `json:"rightNodeId"`
}