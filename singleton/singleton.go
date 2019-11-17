package main

import "sync"

type singleton struct{}

var (
	instance *singleton
	once     sync.Once
	mutex    sync.Mutex
)

func getInstanceIf() *singleton {
	if instance == nil {
		instance = &singleton{}
	}
	return instance
}

func getInstanceMutex() *singleton {
	if instance == nil {
		mutex.Lock()
		defer mutex.Unlock()
		instance = &singleton{}
	}
	return instance
}

func getInstanceMutexOptimised() *singleton {
	if instance == nil {
		mutex.Lock()
		defer mutex.Unlock()
		if instance == nil {
			instance = &singleton{}
		}
	}
	return instance
}

func getInstanceDo() *singleton {
	once.Do(func() {
		instance = &singleton{}
	})
	return instance
}

func main() {}
