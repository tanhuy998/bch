package atomic

import (
	"sync"
)

type (
	/*
		SynchronizerEmitter struct can distributes SynchronizeListeners.
		when Emitter is syncing, it's distributed listeners's locking accquistions
		are delayed until the synchronization done
	*/
	SynchronizeEmitter struct {
		sync.RWMutex
		sync_ch        chan struct{}
		currentSyncNum int
		lastSyncNum    int
	}
)

func (this *SynchronizeEmitter) Sync(num int) {

	this.Lock()
	defer this.Unlock()

	// this.lastSyncNum = this.currentSyncNum
	// this.currentSyncNum = num

	this.lastSyncNum, this.currentSyncNum = this.currentSyncNum, num

	this.sync_ch = make(chan struct{}, num)
	defer close(this.sync_ch)

	if num <= 0 {
		return
	}

	for range num {
		this.sync_ch <- struct{}{}
	}
}

func (this *SynchronizeEmitter) LockAndSyncLater() func(num int) {

	this.Lock()

	return func(num int) {

		defer this.Unlock()

		this.lastSyncNum = this.currentSyncNum
		this.currentSyncNum = num

		this.sync_ch = make(chan struct{}, num)
		defer close(this.sync_ch)

		if num <= 0 {
			return
		}

		for range num {
			this.sync_ch <- struct{}{}
		}
	}
}

func (this *SynchronizeEmitter) Distribute() *SynchronizeListener {

	this.Sync(this.lastSyncNum + 1)

	return (*SynchronizeListener)(this)
}

type (
	SynchronizeListener SynchronizeEmitter
)

func (this *SynchronizeListener) IsInSyncCycle() bool {
	return this.sync_ch != nil && this.lastSyncNum != cap(this.sync_ch) && this.lastSyncNum != this.currentSyncNum
}

func (this *SynchronizeListener) Sync(notifyAfterSync func()) {

	this.RLock()
	defer this.RUnlock()

	if !this.IsInSyncCycle() {
		return
	}

	<-this.sync_ch
	notifyAfterSync()
}
