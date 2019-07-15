package starsystem

import "log"

type StarSystemFunction func(starSystem *StarSystem) error

func (starSystem *StarSystem) execute(function StarSystemFunction) {
	err := function(starSystem)
	if err != nil {
		log.Printf("fail to execute function in star-system [ %d ], because %v", starSystem.ID, err)
	}
}
