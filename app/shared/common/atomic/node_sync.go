package atomic

import (
	"context"
	"sync"
)

type (
	node_mutex_sync_ch_t chan struct{}
)

type (
	INodeSynchronization interface {
		sync(ch node_mutex_sync_ch_t)
	}
)

type (
	node_mutex_t struct {
		synchronizer *SynchronizeListener
		mutex        sync.RWMutex
		sync_context context.Context
		linked       bool
	}
)

func (this *node_mutex_t) isSynchronizing() bool {

	return this.sync_context == nil || this.sync_context.Err() != nil
}

func (this *node_mutex_t) Sync() {

	if this.isSynchronizing() {
		return
	}

	ctx, cancel := context.WithCancel(context.TODO())

	this.sync_context = ctx

	this.synchronizer.Sync(cancel)
}

func (this *node_mutex_t) Lock() {

	this.Sync()

	if !this.linked {
		return
	}

	this.mutex.Lock()
}

func (this *node_mutex_t) Unlock() {

	if !this.linked {
		return
	}

	this.mutex.Unlock()
}

func (this *node_mutex_t) RLock() {

	this.Sync()

	if !this.linked {
		return
	}

	this.mutex.RLock()
}

func (this *node_mutex_t) RUnlock() {

	this.Sync()

	if !this.linked {
		return
	}

	this.mutex.RUnlock()
}

func (this *node_mutex_t) TryLock() bool {

	return this.linked && !this.isSynchronizing() && this.mutex.TryLock()
}

func (this *node_mutex_t) TryRLock() bool {

	return this.linked && !this.isSynchronizing() && this.mutex.TryRLock()
}
