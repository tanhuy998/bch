package memoryCache

import (
	"fmt"
	"io"
	"sync"
)

var (
	log_mgr = new(log_manager)
)

type (
	log_manager struct {
		sync.Mutex
		writers map[string]io.Writer
	}
)

func (this *log_manager) AddListener(key string, writter io.Writer) {

	this.Lock()
	defer this.Unlock()

	this.writers[key] = writter
}

func (this *log_manager) RemoveListener(key string) {

	this.Lock()
	defer this.Unlock()

	delete(this.writers, key)
}

func (this *log_manager) Write(msg string) {

	// log_mgr.RLock()
	// defer log_mgr.RUnlock()

	for key, w := range this.writers {

		if _, ok := this.writers[key]; !ok {

			continue
		}

		if w == nil {

			continue
		}

		w.Write([]byte(fmt.Sprintf("%s\n", msg)))
	}
}

func AddLogListener(key string, writter io.Writer) {

	log_mgr.AddListener(key, writter)
}

func RemoveListener(key string) {

	log_mgr.RemoveListener(key)
}

func writeLog(msg string) {

	log_mgr.Write(msg)
}
