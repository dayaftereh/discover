package persistence

import (
	"fmt"
	"image"
	"path"
	"sync"

	"github.com/dayaftereh/discover/server/game/persistence/types"
)

const (
	playersFileName      = "players.json"
	universeFileName     = "universe.json"
	starSystemsDirectory = "star-systems"
	starSystemFileName   = "star-system.json"
)

type PersistenceManager struct {
	lock      sync.Mutex
	fileLocks map[string]*sync.Mutex
	directory string
}

func NewPersistenceManager(directory string) *PersistenceManager {
	return &PersistenceManager{
		directory: directory,
		fileLocks: make(map[string]*sync.Mutex),
	}
}

func (persistenceManager *PersistenceManager) playersFile() string {
	return path.Join(persistenceManager.directory, playersFileName)
}

func (persistenceManager *PersistenceManager) universeFile() string {
	return path.Join(persistenceManager.directory, universeFileName)
}

func (persistenceManager *PersistenceManager) starSystemDirectory(name string) string {
	return path.Join(persistenceManager.directory, starSystemsDirectory, name)
}

func (persistenceManager *PersistenceManager) starSystemFile(name string) string {
	directory := persistenceManager.starSystemDirectory(name)
	return path.Join(directory, starSystemFileName)
}

func (persistenceManager *PersistenceManager) requestFileLock(filename string) *sync.Mutex {
	// lock the persistenceManager for the file locks
	persistenceManager.lock.Lock()
	defer persistenceManager.lock.Unlock()

	// get the file lock
	lock, ok := persistenceManager.fileLocks[filename]
	// create a new file lock if needed
	if !ok {
		lock = &sync.Mutex{}
		persistenceManager.fileLocks[filename] = lock
	}
	return lock
}

func (persistenceManager *PersistenceManager) LoadPlayers() (map[string]*types.Player, error) {
	// get the path to the player file
	playersFile := persistenceManager.playersFile()

	// request the lock
	lock := persistenceManager.requestFileLock(playersFile)
	// lock the players file
	lock.Lock()
	defer lock.Unlock()

	var players map[string]*types.Player
	// read the players
	err := readJSON(playersFile, &players)
	if err != nil {
		return nil, err
	}
	return players, nil
}

func (persistenceManager *PersistenceManager) WritePlayers(players map[string]*types.Player) error {
	// get the path to the player file
	playersFile := persistenceManager.playersFile()

	// request the lock
	lock := persistenceManager.requestFileLock(playersFile)
	// lock the players file
	lock.Lock()
	defer lock.Unlock()
	// write the players
	err := writeJSON(playersFile, players)
	return err
}

func (persistenceManager *PersistenceManager) LoadUniverse() (*types.Universe, error) {
	// get the path to the universe file
	universeFile := persistenceManager.universeFile()

	// request the lock
	lock := persistenceManager.requestFileLock(universeFile)
	// lock the universe file
	lock.Lock()
	defer lock.Unlock()

	universe := &types.Universe{
		InitialStarSystem: nil,
		StarSystems:       make(map[string][]string),
	}
	// read the universe
	err := readJSON(universeFile, universe)
	if err != nil {
		return nil, err
	}
	return universe, nil
}

func (persistenceManager *PersistenceManager) WriteUniverse(universe *types.Universe) error {
	// get the path to the universe file
	universeFile := persistenceManager.universeFile()

	// request the lock
	lock := persistenceManager.requestFileLock(universeFile)
	// lock the universe file
	lock.Lock()
	defer lock.Unlock()

	// write the universe
	err := writeJSON(universeFile, universe)

	return err
}

func (persistenceManager *PersistenceManager) LoadStarSystem(name string) (*types.StarSystem, error) {
	// get the path to the star-system file
	starSystemFile := persistenceManager.starSystemFile(name)

	// request the lock
	lock := persistenceManager.requestFileLock(starSystemFile)
	// lock the star-system file
	lock.Lock()
	defer lock.Unlock()

	starSystem := &types.StarSystem{}
	// read the star-system
	err := readJSON(starSystemFile, starSystem)
	if err != nil {
		return nil, err
	}
	return starSystem, nil
}

func (persistenceManager *PersistenceManager) WriteStarSystem(starSystem *types.StarSystem) error {
	// get the path to the star-system file
	starSystemFile := persistenceManager.starSystemFile(starSystem.Name)

	// request the lock
	lock := persistenceManager.requestFileLock(starSystemFile)
	// lock the star-system file
	lock.Lock()
	defer lock.Unlock()

	// write the universe
	err := writeJSON(starSystemFile, starSystem)

	return err
}

func (persistenceManager *PersistenceManager) PlanetTexturesFile(starSystemName string, planetName string) string {
	starSystemDirectory := persistenceManager.starSystemDirectory(starSystemName)
	filename := fmt.Sprintf("planet_%s.png", planetName)
	return path.Join(starSystemDirectory, filename)
}

func (persistenceManager *PersistenceManager) WritePlanetTextures(starSystemName string, planetName string, img *image.RGBA) error {
	planetTexturesFile := persistenceManager.PlanetTexturesFile(starSystemName, planetName)

	err := writeImage(planetTexturesFile, img)
	return err
}
