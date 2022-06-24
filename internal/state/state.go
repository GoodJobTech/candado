package state

type Locker interface {
	Lock(id string) error
	Unlock(id string) error
	Heartbeat(id string) (uint16, error)
}

type State struct {
	locker Locker
}

func New(locker Locker) *State {
	return &State{
		locker: locker,
	}
}

func (s *State) Lock(id string) error {
	return s.locker.Lock(id)
}

func (s *State) Unlock(id string) error {
	return s.locker.Unlock(id)
}
func (s *State) Heartbeat(id string) (uint16, error) {
	return s.locker.Heartbeat(id)
}
