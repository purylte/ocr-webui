package types

import (
	"sync"

	"github.com/otiai10/gosseract/v2"
)

type TessOption struct {
	Languages []string
	PCMMode   gosseract.PageSegMode
}

type OCRClient struct {
	Client *gosseract.Client
	Mutex  sync.Mutex
}
