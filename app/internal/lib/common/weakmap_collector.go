package libCommon

import (
	"app/internal/internal/interupt"
	"app/shared/common/weak"
	"time"
)

var (
	weak_map_watch_list map[weak.Pointer[IWeakMapCleaner]]struct{} = make(map[weak.Pointer[IWeakMapCleaner]]struct{})
)

type (
	IWeakMapCleaner interface {
		clean()
	}
)

func init() {

	go __collect_deallocated_weak_map_values__()
}

func __collect_deallocated_weak_map_values__() {

	ticker := time.NewTicker(5 * time.Minute)

	for {
		select {
		case <-interupt.C():
			ticker.Stop()
			return
		case <-ticker.C:

			for weakPointer := range weak_map_watch_list {

				weakMapCleaner := weakPointer.Value()

				if weakMapCleaner == nil {
					delete(weak_map_watch_list, weakPointer)
					continue
				}

				(*weakMapCleaner).clean()
			}
		}
	}
}

func __watch_weak_map__(weakMap IWeakMapCleaner) {

	weak_map_watch_list[weak.Make(&weakMap)] = struct{}{}
}
