package queue

import "sync"

type ConcurrentQueue struct {
	// Mutex lock
	lock *sync.Mutex

	// empty and full locks
	notEmpty *sync.Cond
	notFull  *sync.Cond

	// Queue storage
	backend *QueueBackend
}
