package helpers

import (
	"time"

	"github.com/yunjuiot/tuyacloud"
)

// StaticStore for token.
type StaticStore struct {
	token string
}

// Token for StaticStore.Token()
func (s *StaticStore) Token() string {
	return s.token
}

func (s *StaticStore) IsExpiresAt(d time.Duration) bool {
	return false
}

// Refresh for StaticStore.Refresh()
func (s *StaticStore) Refresh(c *tuyacloud.Client) error {
	return nil
}

// NewStaticStore returns new static store.
func NewStaticStore(t string) *StaticStore {
	return &StaticStore{token: t}
}
