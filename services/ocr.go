package services

import (
	"github.com/otiai10/gosseract/v2"
	"github.com/purylte/ocr-webui/stores"
)

type OCRService struct {
	clientStore *stores.OCRClientStore
}

func NewOCRService(clientStore *stores.OCRClientStore) *OCRService {
	return &OCRService{
		clientStore: clientStore,
	}
}

func (s *OCRService) GetLanguages(sessionId string) []string {
	client := s.clientStore.GetOrInitClient(sessionId)
	return client.CurrentLanguages
}
func (s *OCRService) GetPSM(sessionId string) gosseract.PageSegMode {
	client := s.clientStore.GetOrInitClient(sessionId)
	return client.CurrentPSM
}

func (s *OCRService) SetLanguages(sessionId string, languages []string) error {
	client := s.clientStore.GetOrInitClient(sessionId)
	client.Mutex.Lock()
	defer client.Mutex.Unlock()

	if err := client.Client.SetLanguage(languages...); err != nil {
		return err
	}
	return nil
}

func (s *OCRService) SetPSM(sessionId string, psm gosseract.PageSegMode) error {
	client := s.clientStore.GetOrInitClient(sessionId)
	client.Mutex.Lock()
	defer client.Mutex.Unlock()

	if err := client.Client.SetPageSegMode(psm); err != nil {
		return err
	}
	return nil
}

func (s *OCRService) OcrFromBytes(sessionId string, imageByte []byte) (string, error) {
	client := s.clientStore.GetOrInitClient(sessionId)
	client.Mutex.Lock()
	defer client.Mutex.Unlock()

	if err := client.Client.SetImageFromBytes(imageByte); err != nil {
		return "", err
	}

	text, err := client.Client.Text()
	if err != nil {
		return "", err
	}
	return text, nil
}
