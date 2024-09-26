package services

import (
	"context"
	"crypto/rand"
	"encoding/hex"

	"github.com/alexedwards/scs/v2"
)

type SessionService struct {
	sessionManager scs.SessionManager
}

func NewSessionService(sm scs.SessionManager) *SessionService {
	return &SessionService{
		sessionManager: sm,
	}
}

func (s *SessionService) GetOrGenerateId(ctx context.Context) string {
	if id, exists := s.sessionManager.Get(ctx, "id").(string); exists {
		return id
	}

	newId := generateRandomId()
	s.sessionManager.Put(ctx, "id", newId)
	return newId
}

func generateRandomId() string {
	bytes := make([]byte, 16)
	if _, err := rand.Read(bytes); err != nil {
		panic(err)
	}
	return hex.EncodeToString(bytes)
}
