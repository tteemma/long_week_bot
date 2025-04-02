package storage

import (
	"encoding/json"
	"github.com/robfig/cron/v3"
	"go.etcd.io/bbolt"
	"strconv"
	"time"
)

type Storage struct {
	db *bbolt.DB
}

func NewStorage(db *bbolt.DB) *Storage {
	return &Storage{db: db}
}

type UserState struct {
	Name       string
	BirthDate  time.Time
	CronID     cron.EntryID
	IsWaitDate bool
}

func (s *Storage) SaveUser(chatID int64, user *UserState) error {
	return s.db.Update(func(tx *bbolt.Tx) error {
		bkt, err := tx.CreateBucketIfNotExists([]byte("users"))
		if err != nil {
			return err
		}
		data, err := json.Marshal(user)
		return bkt.Put(intToBytes(chatID), data)
	})
}
func (s *Storage) GetUser(chatID int64) (*UserState, error) {
	var userStates UserState
	err := s.db.View(func(tx *bbolt.Tx) error {
		bkt := tx.Bucket([]byte("users"))
		if bkt == nil {
			return nil
		}
		data := bkt.Get(intToBytes(chatID))
		return json.Unmarshal(data, &userStates)
	})
	if err != nil {
		return nil, err
	}
	return &userStates, nil
}

func (s *Storage) GetAllUsersAfterRefresh() (map[int64]*UserState, error) {
	users := make(map[int64]*UserState)
	err := s.db.View(func(tx *bbolt.Tx) error {
		bkt := tx.Bucket([]byte("users"))
		if bkt == nil {
			return nil
		}
		return bkt.ForEach(func(k, v []byte) error {
			var userState UserState

			if err := json.Unmarshal(v, &userState); err != nil {
				return err
			}
			chatID, err := strconv.ParseInt(string(k), 10, 64)
			if err != nil {
				return err
			}

			users[chatID] = &userState

			return nil
		})
	})
	if err != nil {
		return nil, err
	}
	return users, nil
}

func intToBytes(i int64) []byte {
	return []byte(strconv.FormatInt(int64(i), 10))
}
