package generator

import (
	"fmt"

	"github.com/dayaftereh/discover/server/game/data"
	"github.com/dayaftereh/discover/server/utils"
)

// https://eldacur.com/~brons/NerdCorner/StarGen/StarGen.html

const (
	// AU is a factor to convert Astronomical unit (AU) to meter
	AU float64 = 149597870700
)

func GenerateStarSystem(id int64) *data.StarSystem {
	//name := randStarSystemName()
	return nil
}

func randStarSystemName() string {
	prefix, _ := utils.RandString(2)
	counter := utils.RandInt64(1, 999)
	return fmt.Sprintf("%s-%3.d", prefix, counter)
}

/**
Stellar Fusion Requirements: https://sites.uni.edu/morgans/astro/course/Notes/section2/fusion.html
*/
