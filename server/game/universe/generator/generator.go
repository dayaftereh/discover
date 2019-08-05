package generator

import (
	"fmt"

	"github.com/dayaftereh/discover/server/game/persistence"
	"github.com/dayaftereh/discover/server/game/persistence/types"
	"github.com/dayaftereh/discover/server/game/universe/generator/stargen"
	"github.com/dayaftereh/discover/server/mathf"
	"github.com/dayaftereh/discover/server/utils"
)

type Generator struct {
	run         string
	persistence *persistence.PersistenceManager
}

func NewGenerator(persistenceManager *persistence.PersistenceManager) *Generator {
	run := utils.RandStringWithCharset(2, utils.UpperCaseCharset)
	return &Generator{
		run:         run,
		persistence: persistenceManager,
	}
}

func (generator *Generator) Generate(name string) (*types.StarSystem, error) {
	sun, planets := stargen.GenerateStellarSystem(true, true, true)

	// load all planets
	for index, planet := range planets {
		// find the planet name
		planet.Name = generator.planetName(name, index)
		// find the moon name
		for moonIndex, moon := range planet.Moons {
			moon.Name = generator.moonName(name, index, moonIndex)
		}
	}

	// default spawn location
	spawnLocation := mathf.NewVec3(0, 0, 0)

	// create the star system
	starSystem := &types.StarSystem{
		Name:          name,
		Planets:       planets,
		Sun:           sun,
		SpawnLocation: spawnLocation,
	}
	// write the star system
	err := generator.persistence.WriteStarSystem(starSystem)
	return starSystem, err
}

func (generator *Generator) RandomStarSystemName() string {
	n := utils.RandInt64(0, 999)
	letters := utils.RandStringWithCharset(4, utils.UpperCaseCharset)
	return fmt.Sprintf("%s-%03d-%s", letters, n, generator.run)
}

func (generator *Generator) planetName(starSystemName string, index int) string {
	return fmt.Sprintf("%s-%02d", starSystemName, index)
}

func (generator *Generator) moonName(starSystemName string, planetIndex int, index int) string {
	return fmt.Sprintf("%s-%02d.%d", starSystemName, planetIndex, index)
}
