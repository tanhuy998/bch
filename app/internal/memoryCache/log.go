package memoryCache

import (
	"errors"
	"fmt"
	"io"
	"os"
	"sync"
)

var (
	log_mgr = &log_manager{
		writers: make(map[string]io.Writer),
	}
)

type (
	log_manager struct {
		sync.Mutex
		writers map[string]io.Writer
	}
)

func init() {

	log_mgr.AddListener("stdout", os.Stdout)
}

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

func (this *log_manager) Write(b []byte) (int, error) {

	for key, w := range this.writers {

		if _, ok := this.writers[key]; !ok {

			continue
		}

		if w == nil {

			continue
		}

		_, err := w.Write(b)

		switch {
		case errors.Is(err, io.EOF):
			this.RemoveListener(key)
			continue
		}
	}

	return len(b), nil
}

func AddLogListener(key string, writter io.Writer) {

	log_mgr.AddListener(key, writter)
}

func RemoveListener(key string) {

	log_mgr.RemoveListener(key)
}

func formatLogContent[Key_T any](key Key_T, op string, msg string) string {

	return fmt.Sprintf(`-key[%v]|operation(%s) %s`, key, op, msg)
}
