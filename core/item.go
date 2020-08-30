package core

type Item struct {
	*Object
	Type string `json:"type"`
}
