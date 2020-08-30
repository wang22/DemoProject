package core

import (
	"strings"
)

type Object struct {
	Name       string `json:"name"`
	minX       int
	minY       int
	maxX       int
	maxY       int
	Width      int               `json:"width"`
	Height     int               `json:"height"`
	X          int               `json:"x"`
	Y          int               `json:"y"`
	Properties map[string]string `json:"properties"`
}

func (pos *Object) Info() (minx, minY, maxX, maxY int) {
	return pos.minX, pos.minY, pos.maxX, pos.maxY
}

func (pos *Object) ReadData(layerData string) {
	lines := strings.Split(layerData, "\n")
	y := 0
	for i := len(lines) - 1; i >= 0; i-- {
		line := lines[i]
		if strings.TrimSpace(line) == "" {
			continue
		}
		cols := strings.Split(line, ",")
		//var colArr []bool
		x := -1
		for _, col := range cols {
			x++
			if strings.TrimSpace(col) == "" || col == "0" {
				continue
			}
			if x > pos.maxX {
				pos.maxX = x
			}
			if x < pos.minX {
				pos.minX = x
			}
			if y > pos.maxY {
				pos.maxY = y
			}
			if y < pos.minY {
				pos.minY = y
			}
		}
		y++
	}
	pos.Width, pos.Height = pos.GetSize()
	pos.X, pos.Y = pos.GetPosition()
}

func (pos *Object) GetSize() (width, height int) {
	return pos.maxX - pos.minX + 1, pos.maxY - pos.minY + 1
}

func (pos *Object) GetPosition() (x, y int) {
	return pos.minX, pos.minY
}

func (pos *Object) PutProp(key, val string) {
	if pos.Properties == nil {
		pos.Properties = make(map[string]string)
	}
	pos.Properties[key] = val
}

func NewObject() *Object {
	return &Object{
		minX: 99999999999,
		minY: 99999999999,
		maxX: -1,
		maxY: -1,
	}
}
