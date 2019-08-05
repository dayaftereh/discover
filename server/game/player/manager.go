package player

import (
	"sync"

	"github.com/dayaftereh/discover/server/game/persistence"
	"github.com/dayaftereh/discover/server/game/persistence/types"
	"github.com/pkg/errors"
)

type Manager struct {
	lock        sync.RWMutex
	sessions    map[string]*Player
	storage     map[string]*types.Player
	persistence *persistence.PersistenceManager
}

func NewPlayerManager(persistenceManager *persistence.PersistenceManager) *Manager {
	return &Manager{
		persistence: persistenceManager,
		sessions:    make(map[string]*Player),
		storage:     make(map[string]*types.Player),
	}
}

func (manager *Manager) Init() error {
	// lock for write
	manager.lock.Lock()
	defer manager.lock.Unlock()

	// load the player from the persistance manager
	players, err := manager.persistence.LoadPlayers()
	if err != nil {
		return err
	}

	// store the players
	manager.storage = players

	return nil
}

func (manager *Manager) SessionByName(id string, name string) (*Player, error) {
	// lock for read
	manager.lock.RLock()

	// check if player has already a session
	player, ok := manager.sessions[id]

	// unlock read
	manager.lock.RUnlock()

	// if session exists
	if ok {
		return player, nil
	}

	// if not lock for write
	manager.lock.Lock()
	defer manager.lock.Unlock()

	// get the player data
	playerData, ok := manager.storage[name]

	// create a new player in storage
	if !ok {
		// create a new player
		playerData = &types.Player{
			Name:  name,
			Admin: false,
		}
		// storage the player
		manager.storage[name] = playerData
		// storage the new created player
		err := manager.persistence.WritePlayers(manager.storage)
		if err != nil {
			return nil, err
		}
	}

	// create a new player
	player = NewPlayer(id, playerData)
	// store a new session for the player
	manager.sessions[id] = player

	return player, nil

}

// GetSession the player for the given id, otherwise null
func (manager *Manager) GetSession(id string) *Player {
	// lock for read
	manager.lock.RLock()
	defer manager.lock.RUnlock()

	player, ok := manager.sessions[id]
	if !ok {
		return nil
	}
	return player
}

func (manager *Manager) DropSession(id string) *Player {
	// lock for write
	manager.lock.Lock()
	defer manager.lock.Unlock()

	// get the player
	player, ok := manager.sessions[id]
	if !ok {
		return nil
	}
	// remove the player
	delete(manager.sessions, id)
	return player
}

func (manager *Manager) UpdatePlayerStarSystem(id string, starSystem string) error {
	// lock for write
	manager.lock.Lock()
	defer manager.lock.Unlock()

	// get the player
	player, ok := manager.sessions[id]
	if !ok {
		return errors.Errorf("fail to update star-system for player session [ %s ], because session not found", id)
	}

	// set the star system for the player
	player.StarSystem = &starSystem

	// get the player data
	playerData, ok := manager.storage[player.Name]

	// if data not exists
	if !ok {
		return errors.Errorf("fail to update star-system for player [ %s ], because not player found in storage", id)
	}

	// set the star system for the player
	playerData.StarSystem = &starSystem

	// write the changed player storage
	err := manager.persistence.WritePlayers(manager.storage)
	return err
}
