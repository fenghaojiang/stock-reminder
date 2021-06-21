package action

import "sync"

var isSendMap map[string]bool
var mu sync.Mutex

func init() {
	isSendMap = make(map[string]bool)
}

func ResetIsSendMap() {
	mu.Lock()
	isSendMap = make(map[string]bool)
	mu.Unlock()
}

func IsSendToday(mailID string) bool {
	mu.Lock()
	defer mu.Unlock()
	return isSendMap[mailID]
}
