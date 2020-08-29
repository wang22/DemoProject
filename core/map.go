package core

type MapConfig struct {
	Name    string    `json:"name"`
	Objects []*Object `json:"objects"`
}

func NewMapConfig() *MapConfig {
	return &MapConfig{}
}
