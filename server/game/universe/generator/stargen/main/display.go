package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"math"
	"os"
	"path"

	"github.com/dayaftereh/discover/server/mathf"

	"github.com/dayaftereh/discover/server/utils"

	"github.com/dayaftereh/discover/server/game/persistence/types"
)

type outputContext struct {
	Name    string
	Sun     *types.Sun
	Planets []*types.Planet
}

type Display struct {
	outputDirectory    string
	starSystemTemplate *template.Template
}

func defaultFuncMap() template.FuncMap {
	funcMap := template.FuncMap{
		"solarMass2KG": func(mass float64) float64 {
			return mass * 1.98847e30
		},
		"au2KM": func(au float64) float64 {
			return au * 149597870.700
		},
		"kelvin2Degrees": func(kelvin float64) float64 {
			return kelvin - 272.15
		},
		"decimal2hex": func(value int64) string {
			return fmt.Sprintf("#%X", value)
		},
		"cmPerS2mPerS": func(x float64) float64 {
			return x * 0.01
		},
		"cmPerSSqr2mPerSSqr": func(x float64) float64 {
			return x * 0.01
		},
		"earthGravities2mPerSSqr": func(x float64) float64 {
			return x * 9.80665
		},
		"percent": func(x float64) float64 {
			return x * 100.0
		},
		"infinity": func(x float64) float64 {
			if mathf.CloseEquals(x, math.MaxFloat64) {
				return math.Inf(1)
			}
			return x
		},
		"silicatesMassFraction": func(planet *types.Planet) float64 {
			return 1.0 - planet.CarbonMassFraction
		},
		"ironMassFraction": func(planet *types.Planet) float64 {
			return 1.0 - (planet.RockMassFraction + planet.IceMassFraction)
		},
		"json": func(planet *types.Planet) string {
			bytes, err := json.Marshal(planet)
			if err != nil {
				return err.Error()
			}

			return string(bytes)
		},
	}

	return funcMap
}

func NewDisplay(outputDirectory string, templateFile string) (*Display, error) {
	starSystemTemplate, err := template.New("template.html").Funcs(defaultFuncMap()).ParseGlob(templateFile)
	if err != nil {
		return nil, err
	}

	return &Display{
		outputDirectory:    outputDirectory,
		starSystemTemplate: starSystemTemplate,
	}, nil
}

func (display *Display) export(sun *types.Sun, planets []*types.Planet, name string) error {

	// check if the output directory exists
	exists, err := utils.Exists(display.outputDirectory)
	if err != nil {
		return err
	}

	if !exists {
		// create all directories
		err = os.MkdirAll(display.outputDirectory, os.ModePerm)
		if err != nil {
			return err
		}
	}

	outputFilename := fmt.Sprintf("%s.html", name)
	outputFile := path.Join(display.outputDirectory, outputFilename)

	// create the output file
	f, err := os.Create(outputFile)
	if err != nil {
		return err
	}
	// close the file at the end
	defer f.Close()

	context := &outputContext{
		Name:    name,
		Sun:     sun,
		Planets: planets,
	}

	// execute and output the template
	err = display.starSystemTemplate.Execute(f, context)

	return err

}
