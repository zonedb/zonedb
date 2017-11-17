package build

import "sync"

type semaphore struct {
	c  chan struct{}
	wg sync.WaitGroup
}

func newSemaphore(concurrency int) *semaphore {
	return &semaphore{
		c: make(chan struct{}, concurrency),
	}
}

func (s *semaphore) do(f func()) {
	s.c <- struct{}{}

	go func() {
		s.wg.Add(1)
		defer func() {
			<-s.c
			s.wg.Done()
		}()
		f()
	}()
}

func (s *semaphore) wait() {
	s.wg.Wait()
}
