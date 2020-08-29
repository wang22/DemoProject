package core

type MapConfig struct {
	Name    string         `json:"name"`
	Objects []interface{} `json:"objects"`
}

func NewMapConfig() *MapConfig {
	return &MapConfig{}
}
