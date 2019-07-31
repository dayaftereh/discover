package noise_test

import (
	"testing"

	"github.com/dayaftereh/discover/server/game/universe/generator/texture/noise"
	"github.com/dayaftereh/discover/server/mathf"
	"github.com/stretchr/testify/assert"
)

type dummyImplicitBase struct{}

func (dummy *dummyImplicitBase) Get2D(x, y float64) float64 {
	return x * y
}

func (dummy *dummyImplicitBase) Get3D(x, y, z float64) float64 {
	return x * y * z
}

func (dummy *dummyImplicitBase) Get4D(x, y, z, w float64) float64 {
	return x * y * z * w
}

func CreateCombiner(combinerType noise.ImplicitCombinerType) *noise.ImplicitCombiner {
	i := noise.NewImplicitCombiner(combinerType)

	dummy := &dummyImplicitBase{}
	i.AddSource(dummy)
	i.AddSource(dummy)

	return i
}

func TestImplicitCombinerAdd2D(t *testing.T) {
	i := CreateCombiner(noise.CombinerTypeAdd)

	v := i.Get2D(1.0, 2.0)

	assert.InDelta(t, v, (1.0*2.0)+(1.0*2.0), mathf.Epsilon)
}

func TestImplicitCombinerAdd3D(t *testing.T) {
	i := CreateCombiner(noise.CombinerTypeAdd)

	v := i.Get3D(1.0, 2.0, 3.0)

	assert.InDelta(t, v, (1.0*2.0*3.0)+(1.0*2.0*3.0), mathf.Epsilon)
}

func TestImplicitCombinerAdd4D(t *testing.T) {
	i := CreateCombiner(noise.CombinerTypeAdd)

	v := i.Get4D(1.0, 2.0, 3.0, 4.0)

	assert.InDelta(t, v, (1.0*2.0*3.0*4.0)+(1.0*2.0*3.0*4.0), mathf.Epsilon)
}

func TestImplicitCombinerMultiply2D(t *testing.T) {
	i := CreateCombiner(noise.CombinerTypeMultiply)

	v := i.Get2D(1.0, 2.0)

	assert.InDelta(t, v, (1.0*2.0)*(1.0*2.0), mathf.Epsilon)
}

func TestImplicitCombinerMultiply3D(t *testing.T) {
	i := CreateCombiner(noise.CombinerTypeMultiply)

	v := i.Get3D(1.0, 2.0, 3.0)

	assert.InDelta(t, v, (1.0*2.0*3.0)*(1.0*2.0*3.0), mathf.Epsilon)
}

func TestImplicitCombinerMultiply4D(t *testing.T) {
	i := CreateCombiner(noise.CombinerTypeMultiply)

	v := i.Get4D(1.0, 2.0, 3.0, 4.0)

	assert.InDelta(t, v, (1.0*2.0*3.0*4.0)*(1.0*2.0*3.0*4.0), mathf.Epsilon)
}

func TestImplicitCombinerMin2D(t *testing.T) {
	i := CreateCombiner(noise.CombinerTypeMin)

	v := i.Get2D(1.0, 2.0)

	assert.InDelta(t, v, (1.0 * 2.0), mathf.Epsilon)
}

func TestImplicitCombinerMin3D(t *testing.T) {
	i := CreateCombiner(noise.CombinerTypeMin)

	v := i.Get3D(1.0, 2.0, 3.0)

	assert.InDelta(t, v, (1.0 * 2.0 * 3.0), mathf.Epsilon)
}

func TestImplicitCombinerMin4D(t *testing.T) {
	i := CreateCombiner(noise.CombinerTypeMin)

	v := i.Get4D(1.0, 2.0, 3.0, 4.0)

	assert.InDelta(t, v, (1.0 * 2.0 * 3.0 * 4.0), mathf.Epsilon)
}

func TestImplicitCombinerMax2D(t *testing.T) {
	i := CreateCombiner(noise.CombinerTypeMax)

	v := i.Get2D(1.0, 2.0)

	assert.InDelta(t, v, (1.0 * 2.0), mathf.Epsilon)
}

func TestImplicitCombinerMax3D(t *testing.T) {
	i := CreateCombiner(noise.CombinerTypeMax)

	v := i.Get3D(1.0, 2.0, 3.0)

	assert.InDelta(t, v, (1.0 * 2.0 * 3.0), mathf.Epsilon)
}

func TestImplicitCombinerMax4D(t *testing.T) {
	i := CreateCombiner(noise.CombinerTypeMax)

	v := i.Get4D(1.0, 2.0, 3.0, 4.0)

	assert.InDelta(t, v, (1.0 * 2.0 * 3.0 * 4.0), mathf.Epsilon)
}

func TestImplicitCombinerAverage2D(t *testing.T) {
	i := CreateCombiner(noise.CombinerTypeAverage)

	v := i.Get2D(1.0, 2.0)

	assert.InDelta(t, v, ((1.0*2.0)+(1.0*2.0))/2.0, mathf.Epsilon)
}

func TestImplicitCombinerAverage3D(t *testing.T) {
	i := CreateCombiner(noise.CombinerTypeAverage)

	v := i.Get3D(1.0, 2.0, 3.0)

	assert.InDelta(t, v, ((1.0*2.0*3.0)+(1.0*2.0*3.0))/2.0, mathf.Epsilon)
}

func TestImplicitCombinerAverage4D(t *testing.T) {
	i := CreateCombiner(noise.CombinerTypeAverage)

	v := i.Get4D(1.0, 2.0, 3.0, 4.0)

	assert.InDelta(t, v, ((1.0*2.0*3.0*4.0)+(1.0*2.0*3.0*4.0))/2.0, mathf.Epsilon)
}
