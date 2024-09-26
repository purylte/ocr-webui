package types

import (
	"sync"
	"time"

	"github.com/otiai10/gosseract/v2"
)

type TessOption struct {
	Languages []string
	PCMMode   gosseract.PageSegMode
}

type OCRClient struct {
	Client       *gosseract.Client
	Mutex        sync.Mutex
	LastAccessed time.Time
}
