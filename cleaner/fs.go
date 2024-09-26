package cleaner

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"syscall"
	"time"
)

type FSCleaner struct {
	directory string
	maxAge    time.Duration
	interval  time.Duration
}

func NewFSCleaner(directory string, maxAge time.Duration, interval time.Duration) *FSCleaner {
	return &FSCleaner{
		directory: directory,
		maxAge:    maxAge,
		interval:  interval,
	}
}

func (c *FSCleaner) Clean() error {
	return filepath.Walk(c.directory, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		fileTime := getAccessTime(info)

		if time.Since(fileTime) > c.maxAge {
			if err := os.Remove(path); err != nil {
				return fmt.Errorf("failed to remove file %s: %w", path, err)
			}
		}

		return nil
	})
}

func getAccessTime(info fs.FileInfo) time.Time {
	stat := info.Sys().(*syscall.Stat_t)
	return time.Unix(int64(stat.Atim.Sec), int64(stat.Atim.Nsec))
}

func (c *FSCleaner) Start() {
	ticker := time.NewTicker(c.interval)
	go func() {
		for range ticker.C {
			if err := c.Clean(); err != nil {
				log.Printf("Error during cleanup: %v\n", err)
			}
		}
	}()
}
