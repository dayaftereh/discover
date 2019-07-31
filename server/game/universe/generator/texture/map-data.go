package texture

import "math"

type MapData struct {
	width    int
	height   int
	data     []float64
	MinValue float64
	MaxValue float64
}

func NewMapData(width int, height int) *MapData {
	return &MapData{
		width:    width,
		height:   height,
		MinValue: math.Inf(1),
		MaxValue: math.Inf(-1),
		data:     make([]float64, width*height),
	}
}

func (mapData *MapData) index(x, y int) int {
	// y = i / W
	// x = i % W
	return y*mapData.width + x
}

func (mapData *MapData) updateMinMax(value float64) {
	if value > mapData.MaxValue {
		mapData.MaxValue = value
	}

	if value < mapData.MinValue {
		mapData.MinValue = value
	}
}

func (mapData *MapData) Get(x, y int) float64 {
	index := mapData.index(x, y)
	return mapData.data[index]
}

func (mapData *MapData) Set(x, y int, value float64) {
	mapData.updateMinMax(value)
	index := mapData.index(x, y)
	mapData.data[index] = value
}
