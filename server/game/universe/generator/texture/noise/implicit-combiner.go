package noise

type ImplicitCombinerType string

const (
	CombinerTypeAdd      ImplicitCombinerType = "add"
	CombinerTypeMultiply ImplicitCombinerType = "multiply"
	CombinerTypeMin      ImplicitCombinerType = "min"
	CombinerTypeMax      ImplicitCombinerType = "max"
	CombinerTypeAverage  ImplicitCombinerType = "average"
)

type ImplicitCombiner struct {
	combinerType ImplicitCombinerType
	sources      []ImplicitBase
}

func NewImplicitCombiner(combinerType ImplicitCombinerType) *ImplicitCombiner {
	return &ImplicitCombiner{
		combinerType: combinerType,
		sources:      make([]ImplicitBase, 0),
	}
}

func (implicitCombiner *ImplicitCombiner) AddSource(source ImplicitBase) {
	implicitCombiner.sources = append(implicitCombiner.sources, source)
}

func (implicitCombiner *ImplicitCombiner) Get2D(x, y float64) float64 {
	switch implicitCombiner.combinerType {
	case CombinerTypeAdd:
		return implicitCombiner.AddGet2D(x, y)
	case CombinerTypeMultiply:
		return implicitCombiner.MultiplyGet2D(x, y)
	case CombinerTypeMin:
		return implicitCombiner.MinGet2D(x, y)
	case CombinerTypeMax:
		return implicitCombiner.MaxGet2D(x, y)
	case CombinerTypeAverage:
		return implicitCombiner.AverageGet2D(x, y)
	}
	return 0.0
}

func (implicitCombiner *ImplicitCombiner) Get3D(x, y, z float64) float64 {
	switch implicitCombiner.combinerType {
	case CombinerTypeAdd:
		return implicitCombiner.AddGet3D(x, y, z)
	case CombinerTypeMultiply:
		return implicitCombiner.MultiplyGet3D(x, y, z)
	case CombinerTypeMin:
		return implicitCombiner.MinGet3D(x, y, z)
	case CombinerTypeMax:
		return implicitCombiner.MaxGet3D(x, y, z)
	case CombinerTypeAverage:
		return implicitCombiner.AverageGet3D(x, y, z)
	}
	return 0.0
}

func (implicitCombiner *ImplicitCombiner) Get4D(x, y, z, w float64) float64 {
	switch implicitCombiner.combinerType {
	case CombinerTypeAdd:
		return implicitCombiner.AddGet4D(x, y, z, w)
	case CombinerTypeMultiply:
		return implicitCombiner.MultiplyGet4D(x, y, z, w)
	case CombinerTypeMin:
		return implicitCombiner.MinGet4D(x, y, z, w)
	case CombinerTypeMax:
		return implicitCombiner.MaxGet4D(x, y, z, w)
	case CombinerTypeAverage:
		return implicitCombiner.AverageGet4D(x, y, z, w)
	}
	return 0.0
}

func (implicitCombiner *ImplicitCombiner) AddGet2D(x, y float64) float64 {
	sum := 0.0
	for _, source := range implicitCombiner.sources {
		sum += source.Get2D(x, y)
	}
	return sum
}

func (implicitCombiner *ImplicitCombiner) AddGet3D(x, y, z float64) float64 {
	sum := 0.0
	for _, source := range implicitCombiner.sources {
		sum += source.Get3D(x, y, z)
	}
	return sum
}

func (implicitCombiner *ImplicitCombiner) AddGet4D(x, y, z, w float64) float64 {
	sum := 0.0
	for _, source := range implicitCombiner.sources {
		sum += source.Get4D(x, y, z, w)
	}
	return sum
}

func (implicitCombiner *ImplicitCombiner) MultiplyGet2D(x, y float64) float64 {
	aggregate := 1.0
	for _, source := range implicitCombiner.sources {
		aggregate *= source.Get2D(x, y)
	}
	return aggregate
}

func (implicitCombiner *ImplicitCombiner) MultiplyGet3D(x, y, z float64) float64 {
	aggregate := 1.0
	for _, source := range implicitCombiner.sources {
		aggregate *= source.Get3D(x, y, z)
	}
	return aggregate
}

func (implicitCombiner *ImplicitCombiner) MultiplyGet4D(x, y, z, w float64) float64 {
	aggregate := 1.0
	for _, source := range implicitCombiner.sources {
		aggregate *= source.Get4D(x, y, z, w)
	}
	return aggregate
}

func (implicitCombiner *ImplicitCombiner) MinGet2D(x, y float64) float64 {
	min := 0.0
	for index, source := range implicitCombiner.sources {
		value := source.Get2D(x, y)
		if index == 0 || min > value {
			min = value
		}
	}
	return min
}

func (implicitCombiner *ImplicitCombiner) MinGet3D(x, y, z float64) float64 {
	min := 0.0
	for index, source := range implicitCombiner.sources {
		value := source.Get3D(x, y, z)
		if index == 0 || min > value {
			min = value
		}
	}
	return min
}

func (implicitCombiner *ImplicitCombiner) MinGet4D(x, y, z, w float64) float64 {
	min := 0.0
	for index, source := range implicitCombiner.sources {
		value := source.Get4D(x, y, z, w)
		if index == 0 || min > value {
			min = value
		}
	}
	return min
}

func (implicitCombiner *ImplicitCombiner) MaxGet2D(x, y float64) float64 {
	max := 0.0
	for index, source := range implicitCombiner.sources {
		value := source.Get2D(x, y)
		if index == 0 || max < value {
			max = value
		}
	}
	return max
}

func (implicitCombiner *ImplicitCombiner) MaxGet3D(x, y, z float64) float64 {
	max := 0.0
	for index, source := range implicitCombiner.sources {
		value := source.Get3D(x, y, z)
		if index == 0 || max < value {
			max = value
		}
	}
	return max
}

func (implicitCombiner *ImplicitCombiner) MaxGet4D(x, y, z, w float64) float64 {
	max := 0.0
	for index, source := range implicitCombiner.sources {
		value := source.Get4D(x, y, z, w)
		if index == 0 || max < value {
			max = value
		}
	}
	return max
}

func (implicitCombiner *ImplicitCombiner) AverageGet2D(x, y float64) float64 {
	sum := implicitCombiner.AddGet2D(x, y)
	return sum / float64(len(implicitCombiner.sources))
}

func (implicitCombiner *ImplicitCombiner) AverageGet3D(x, y, z float64) float64 {
	sum := implicitCombiner.AddGet3D(x, y, z)
	return sum / float64(len(implicitCombiner.sources))
}

func (implicitCombiner *ImplicitCombiner) AverageGet4D(x, y, z, w float64) float64 {
	sum := implicitCombiner.AddGet4D(x, y, z, w)
	return sum / float64(len(implicitCombiner.sources))
}
