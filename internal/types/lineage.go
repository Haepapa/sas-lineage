package types

type Node struct {
    ID       string `json:"id"`
    Label     string `json:"label"`
    Name     string `json:"name"`
    Fill     string `json:"fill"`
    Size int `json:"size"`
    Type     string `json:"type"`
}

type Link struct {
    ID       string `json:"id"`
    Label     string `json:"label"`
    Source     string `json:"source"`
    Target     string `json:"target"`
}