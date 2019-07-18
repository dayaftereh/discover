package player

import (
	"sync"

	"github.com/dayaftereh/discover/server/game/data"
)

type Manager struct {
	lock     sync.RWMutex
	sessions map[string]*Player
	storage  map[string]*data.Player
}

func NewPlayerManager() *Manager {
	return &Manager{
		sessions: make(map[string]*Player),
		storage:  make(map[string]*data.Player),
	}
}

func (manager *Manager) LoadPlayersFromData(gameData *data.Game) {
	// lock for write
	manager.lock.Lock()
	defer manager.lock.Unlock()

	// store all players
	for name, player := range gameData.Players {
		manager.storage[name] = player
	}
}

func (manager *Manager) WritePlayersToData(gameData *data.Game) {
	// lock for read
	manager.lock.RLock()
	defer manager.lock.RUnlock()

	//copy players to game data
	gameData.Players = make(map[string]*data.Player)
	for name, player := range manager.storage {
		gameData.Players[name] = player
	}
}

func (manager *Manager) SessionByName(id string, name string) *Player {
	// lock for read
	manager.lock.RLock()

	// check if player has already a session
	player, ok := manager.sessions[id]

	// unlock read
	manager.lock.RUnlock()

	// if session exists
	if ok {
		return player
	}

	// if not lock for write
	manager.lock.Lock()
	defer manager.lock.Unlock()

	// get the player data
	playerData, ok := manager.storage[name]

	// create a new player in storage
	if !ok {
		// create a new player
		playerData = &data.Player{}
		playerData.Name = name
		// storage the player
		manager.storage[name] = playerData
	}

	// create a new player
	player = NewPlayer(id, playerData)
	// store a new session for the player
	manager.sessions[id] = player

	return player

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
	// lock for read
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
