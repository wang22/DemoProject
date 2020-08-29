package core

type Obstacle struct {
	*Object
	Type string `json:"type"`
}
