package tests

import (
	"fmt"
	"im/model"
	"strconv"
	"sync"
	"testing"
)

// TestBatchUser 并发场景扛不住
func TestBatchUser(t *testing.T) {
	group := sync.WaitGroup{}
	lock := sync.Mutex{}
	repeats := make(map[int]int)
	users := make([]*model.User, 0)
	for i := 0; i < 1000000; i++ {
		group.Add(1)
		go func(x int) {
			u := model.NewUser(strconv.Itoa(x))
			users = append(users, u)
			if lock.TryLock() {
				repeats[u.Id]++
				lock.Unlock()
			}
			group.Add(-1)
		}(i)
	}
	group.Wait()
	for val, count := range repeats {
		if count > 2 {
			fmt.Println("res>2", count, val)
		}
	}
}
