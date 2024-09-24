package main

import (
	"context"
	"errors"

	"github.com/purylte/ocr-webui/types"
)

func canAccessImage(ctx context.Context, imageName string) bool {
	images := sessionManager.Get(ctx, "images").([]*types.ImageData)
	for _, image := range images {
		if image.Name == imageName {
			return true
		}
	}
	return false
}

func putAllowedImage(ctx context.Context, image *types.ImageData) error {
	images, ok := sessionManager.Get(ctx, "images").([]*types.ImageData)
	if !ok {
		images = []*types.ImageData{image}
	} else {
		images = append(images, image)
	}
	sessionManager.Put(ctx, "images", images)
	return nil
}

func setCurrentImage(ctx context.Context, image *types.ImageData) {
	sessionManager.Put(ctx, "currentImage", *image)
}

func getCurrentImage(ctx context.Context) (*types.ImageData, error) {
	val, ok := sessionManager.Get(ctx, "currentImage").(types.ImageData)
	if !ok {
		return nil, errors.New("type assertion to ImageData failed")
	}
	return &val, nil
}
