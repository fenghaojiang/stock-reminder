package utils

import "sync"

var isSendMap map[string]bool
var mu sync.RWMutex

func init() {
	isSendMap = make(map[string]bool)
}

func ResetIsSendMap() {
	mu.Lock()
	isSendMap = make(map[string]bool)
	mu.Unlock()
}

func IsSendToday(mailID string) bool {
	mu.RLock()
	defer mu.RUnlock()
	return isSendMap[mailID]
}

func SendToday(mailID string) {
	mu.Lock()
	isSendMap[mailID] = true
	mu.Unlock()
}
