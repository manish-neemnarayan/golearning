package main

import "sync/atomic"

type State struct {
	// mu    sync.Mutex
	count int32
}

func (s *State) setState(i int) {
	// s.mu.Lock()
	// defer s.mu.Unlock()

	// s.count = i

	atomic.AddInt32(&s.count, int32(i))
}

func main() {
	state := State{}

	for i := 0; i < 10; i++ {
		state.setState(i + 1)
	}

}
