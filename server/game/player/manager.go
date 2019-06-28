package player

import "sync"

type Manager struct {
	lock    sync.RWMutex
	players map[string]*Player
}

func NewPlayerManager() *Manager {
	return &Manager{
		players: make(map[string]*Player),
	}
}

// Get the player for the given id, otherwise null
func (manager *Manager) Get(id string) *Player {
	// lock for read
	manager.lock.RLock()
	defer manager.lock.RUnlock()

	player, ok := manager.players[id]
	if !ok {
		return nil
	}
	return player
}

func (manager *Manager) Remove(id string) *Player {
	// lock for read
	manager.lock.Lock()
	defer manager.lock.Unlock()

	// get the player
	player, ok := manager.players[id]
	if !ok {
		return nil
	}
	// remove the player
	delete(manager.players, id)
	return player
}

func (manager *Manager) GetOrCreate(id string) *Player {
	// lock for read
	manager.lock.RLock()
	// get the player
	player, ok := manager.players[id]
	// unlock from read
	manager.lock.RUnlock()

	if !ok {
		// create a new player
		player = NewPlayer(id)
		// lock for write
		manager.lock.Lock()
		defer manager.lock.Unlock()
		// store the player
		manager.players[id] = player
	}

	return player
}
