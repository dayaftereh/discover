package texture

import "math"

type MapData struct {
	MinValue float64
	MaxValue float64
	width    int
	height   int
	values   []float64
}

func NewMapData(width, height int) *MapData {
	return &MapData{
		MinValue: math.Inf(1),
		MaxValue: math.Inf(-1),
		width:    width,
		height:   height,
		values:   make([]float64, width*height),
	}
}

func (mapData *MapData) index(x, y int) int {
	index := y*mapData.width + x
	return index
}

func (mapData *MapData) Get(x, y int) float64 {
	index := mapData.index(x, y)
	return mapData.values[index]
}

func (mapData *MapData) GetNormalized(x, y int) float64 {
	index := mapData.index(x, y)
	value := mapData.values[index]

	return (value - mapData.MinValue) / (mapData.MaxValue - mapData.MinValue)
}

func (mapData *MapData) Set(x, y int, value float64) {
	mapData.MinValue = math.Min(value, mapData.MinValue)
	mapData.MaxValue = math.Max(value, mapData.MaxValue)

	index := mapData.index(x, y)
	mapData.values[index] = value
}
