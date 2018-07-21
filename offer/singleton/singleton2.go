package offer_singleton

import "sync"

func GetSingleton() *Singleton {
	if singleton2 == nil {
		lock.Lock()
		defer lock.Unlock()
		if singleton2 == nil {
			singleton2 = &Singleton{"test"}
		}
	}
	return singleton2
}

var singleton2 *Singleton
// sync lock
var lock sync.Mutex
