package stores

import (
	"sync"

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
		return client
	}

	newClient := &types.OCRClient{
		Client: gosseract.NewClient(),
		Mutex:  sync.Mutex{},
	}
	s.clients[sessionID] = newClient
	return newClient
}
