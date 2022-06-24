package memory

import (
	"sync"

	"github.com/goodjobtech/candado/internal/errors"
)

type locks = map[string]bool

type Memory struct {
	locks locks
	mu    sync.Mutex
}

func New() *Memory {
	return &Memory{
		locks: make(locks),
	}
}

func (m *Memory) Lock(id string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	_, ok := m.locks[id]
	if ok {
		return errors.ErrAlreadyLocked
	} else {
		m.locks[id] = true
		return nil
	}
}

func (m *Memory) Unlock(id string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	_, ok := m.locks[id]
	if ok {
		delete(m.locks, id)
		return nil
	} else {
		return errors.ErrAlreadyUnlocked
	}
}

func (m *Memory) Heartbeat(id string) (uint16, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	_, ok := m.locks[id]
	if ok {
		return 1, nil
	} else {
		return 0, nil
	}
}
