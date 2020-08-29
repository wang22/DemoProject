package main

import (
	"bytes"
	"encoding/json"
	"github.com/beevik/etree"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"io/ioutil"
	"strings"
	"tilemap/core"
)

func ReadLevelFile(filePath string) *core.MapConfig {
	doc := etree.NewDocument()

	if err := doc.ReadFromFile(filePath); err != nil {
		panic(err)
	}
	tMap := doc.SelectElement("map")

	propMap := tMap.SelectElement("properties")
	properties := propMap.SelectElements("property")

	mapCfg := &core.MapConfig{}
	for _, prop := range properties {
		name := prop.SelectAttr("name").Value
		value := prop.SelectAttr("value").Value
		if name == "name" {
			mapCfg.Name = value
		}
	}

	layers := tMap.SelectElements("layer")

	for _, layer := range layers {
		name := layer.SelectAttr("name").Value
		obj := core.NewObject()
		obj.Name = name
		obj.ReadData(layer.SelectElement("data").Text())
		if name == "obstacle" {
			obstacle := core.Obstacle{
				Object: obj,
				Type:   "",
			}
			if hasProps := layer.SelectElement("properties"); hasProps != nil {
				props:=hasProps.SelectElements("property")
				for _, prop:=range props {
					name  := prop.SelectAttr("name").Value
					if name == "type" {
						value := prop.SelectAttr("value").Value
						obstacle.Type = value
					}
				}
			}
			mapCfg.Objects = append(mapCfg.Objects, obstacle)
		} else {
			mapCfg.Objects = append(mapCfg.Objects, obj)
		}
	}

	return mapCfg
}

func main() {
	dirPath := "./tiledmap"
	fileInfos, _ := ioutil.ReadDir(dirPath)
	var buf bytes.Buffer
	for _, file := range fileInfos {
		if file.IsDir() || !strings.HasPrefix(file.Name(), "level") {
			continue
		}
		filePath := dirPath + "/" + file.Name()
		mapCfg := ReadLevelFile(filePath)
		jsonStr, _ := json.Marshal(mapCfg)
		buf.WriteString(string(jsonStr))
		buf.WriteString(",")
	}
	jsonArr := "[" + buf.String()[0:len(buf.String()) - 1] + "]"
	ioutil.WriteFile("./static/level.json", []byte(jsonArr), 0677)
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"*"},
	}))
	e.Static("/static", "static")
	e.Logger.Fatal(e.Start(":1323"))
}
