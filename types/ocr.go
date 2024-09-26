package types

import (
	"sync"
	"time"

	"github.com/otiai10/gosseract/v2"
)

type OCRClient struct {
	Client           gosseract.Client
	Mutex            sync.Mutex
	CurrentPSM       gosseract.PageSegMode
	CurrentLanguages []string
	LastAccessed     time.Time
}
