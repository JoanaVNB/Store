package elasticSearch

import (
	"BE-JoanaVidon/user-api/domain"
	"BE-JoanaVidon/user-api/presenter"
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"os"
	"sync"
)

func LoadUsersFromFile(ctx context.Context) context.Context {
const (
	concurrency = 5
	usersFile  = "user.json"
)

var (
	users    []domain.User
	waitGroup = new(sync.WaitGroup)
	workQueue = make(chan string)
	mutex     = &sync.Mutex{}
)

go func() {
	usersFile, err := os.Open(usersFile)
	if err != nil {
		panic(err)
	}
	defer usersFile.Close()
	scanner := bufio.NewScanner(usersFile)
	for scanner.Scan() {
		workQueue <- scanner.Text()
	}
	close(workQueue)
}()

for i := 0; i < concurrency; i++ {
	waitGroup.Add(1)
	go func(workQueue chan string, waitGroup *sync.WaitGroup) {
		for entry := range workQueue {
			userRaw := presenter.PresentUser{}
			err := json.Unmarshal([]byte(entry), &userRaw)
			if err == nil {
				user := func(userRaw presenter.PresentUser) domain.User {
					return domain.User{
						ID:          userRaw.ID,
						Name:        userRaw.Name,
						CPF:         userRaw.CPF,
						Email:       userRaw.Email,
						PhoneNumber: userRaw.PhoneNumber,
					}
				}(userRaw)
				mutex.Lock()
				users = append(users, user)
				mutex.Unlock()
			}
		}
		waitGroup.Done()
	}(workQueue, waitGroup)
}

waitGroup.Wait()

fmt.Printf("✅ Users loaded from the file: %d \n", len(users))
return context.WithValue(ctx, UserKey, users)
}