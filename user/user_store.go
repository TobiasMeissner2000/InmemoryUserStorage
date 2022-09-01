package user

import (
	"errors"
	"sync"
)

type UserStore struct {
	users      map[int]User
	lastUserID int
	mutex      sync.RWMutex
}

func NewUserStore() UserStore {
	return UserStore{
		users:      make(map[int]User),
		lastUserID: 1000,
		mutex:      sync.RWMutex{},
	}
}

func (us *UserStore) Add(u *User) {
	us.mutex.Lock()
	defer us.mutex.Unlock()

	us.lastUserID++
	u.ID = us.lastUserID
	us.users[u.ID] = *u
}

func (us *UserStore) Get(id int) (User, error) {
	us.mutex.RLock()
	defer us.mutex.RUnlock()

	u, found := us.users[id]
	if found {
		return u, nil
	}
	return User{}, errors.New("user not found")
}

func (us *UserStore) GetAll() map[int]User {
	us.mutex.RLock()
	defer us.mutex.RUnlock()

	return us.users
}

func (us *UserStore) Delete(id int) error {
	us.mutex.Lock()
	defer us.mutex.Unlock()

	_, found := us.users[id]
	if found {
		delete(us.users, id)
	}
	return errors.New("user not found")

}
