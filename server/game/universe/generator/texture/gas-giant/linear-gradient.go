package gasgiant

import (
	"container/list"

	"github.com/dayaftereh/discover/server/mathf"

	"github.com/dayaftereh/discover/server/utils/container"

	textureset "github.com/dayaftereh/discover/server/game/universe/generator/texture/texture-set"
)

type linearGradientElement struct {
	value float64
	color *textureset.Color
}

type LinearGradient1D struct {
	colorRange *list.List
}

func NewLinearGradient1D() *LinearGradient1D {
	return &LinearGradient1D{
		colorRange: list.New(),
	}
}

func (linearGradient *LinearGradient1D) AddColor(x float64, color *textureset.Color) {
	x = mathf.Clamp(x, 0.0, 1.0)
	colorElement := &linearGradientElement{
		value: x,
		color: color,
	}

	// find the element for the insert
	element := container.Find(linearGradient.colorRange, func(value interface{}, index int64) bool {
		colorRangeElement := value.(*linearGradientElement)
		return x < colorRangeElement.value
	})

	if element == nil {
		linearGradient.colorRange.PushBack(colorElement)
	} else {
		linearGradient.colorRange.InsertBefore(colorElement, element)
	}

}

func (linearGradient *LinearGradient1D) Color(x float64) *textureset.Color {
	x = mathf.Clamp(x, 0.0, 1.0)
	// no color added
	if linearGradient.colorRange.Len() < 2 {
		return nil
	}

	element := container.Find(linearGradient.colorRange, func(value interface{}, index int64) bool {
		colorRangeElement := value.(*linearGradientElement)
		return x < colorRangeElement.value
	})

	if element == nil {
		element = linearGradient.colorRange.Back()
	}

	prevElement := element.Prev()
	endColorRange := element.Value.(*linearGradientElement)
	if prevElement == nil {
		return endColorRange.color
	}
	startColorRange := prevElement.Value.(*linearGradientElement)

	// norm the range
	value := (x - startColorRange.value) / (endColorRange.value - startColorRange.value)

	color := startColorRange.color.Lerp(endColorRange.color, value)
	return color
}
