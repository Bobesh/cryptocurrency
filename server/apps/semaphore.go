package apps

const MaxProc = 4

type semaphore struct {
	sema chan struct{}
}

func newSemaphore() *semaphore {
	return &semaphore{
		sema: make(chan struct{}, MaxProc),
	}
}

func (s *semaphore) Acquire() {
	s.sema <- struct{}{}
}

func (s *semaphore) Release() {
	<-s.sema
}
