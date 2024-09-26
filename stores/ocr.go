package stores

import (
	"sync"
	"time"

	"github.com/otiai10/gosseract/v2"
	"github.com/purylte/ocr-webui/types"
)

type OCRClientStore struct {
	clients map[string]*types.OCRClient
	mutex   *sync.Mutex
}

func NewOCRClientStore() *OCRClientStore {
	return &OCRClientStore{
		clients: make(map[string]*types.OCRClient),
	}
}

func (s *OCRClientStore) GetOrInitClient(sessionID string) *types.OCRClient {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if client, exists := s.clients[sessionID]; exists {
		client.LastAccessed = time.Now()
		return client
	}

	newClient := &types.OCRClient{
		Client:       gosseract.NewClient(),
		Mutex:        sync.Mutex{},
		LastAccessed: time.Now(),
	}
	s.clients[sessionID] = newClient
	return newClient
}

func (s *OCRClientStore) Cleanup(expiration time.Duration) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	now := time.Now()
	for sessionID, client := range s.clients {
		if now.Sub(client.LastAccessed) > expiration {
			client.Client.Close()
			delete(s.clients, sessionID)
		}
	}
}
