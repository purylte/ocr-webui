package cleaner

import (
	"time"

	"github.com/purylte/ocr-webui/stores"
)

type OCRClientCleaner struct {
	clientStore *stores.OCRClientStore
	interval    time.Duration
	expiration  time.Duration
}

func NewOCRClientCleaner(clientStore *stores.OCRClientStore, interval, expiration time.Duration) *OCRClientCleaner {
	return &OCRClientCleaner{
		clientStore: clientStore,
		interval:    interval,
		expiration:  expiration,
	}
}

func (c *OCRClientCleaner) Start() {
	ticker := time.NewTicker(c.interval)
	go func() {
		for range ticker.C {
			c.clientStore.Cleanup(c.expiration)
		}
	}()
}
