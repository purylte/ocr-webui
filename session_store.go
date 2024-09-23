package main

import (
	"context"
	"errors"
)

func canAccessImage(ctx context.Context, imageName string) bool {
	images := sessionManager.Get(ctx, "images").([]*ImageData)
	for _, image := range images {
		if image.Name == imageName {
			return true
		}
	}
	return false
}

func putAllowedImage(ctx context.Context, image *ImageData) error {
	images, ok := sessionManager.Get(ctx, "images").([]*ImageData)
	if !ok {
		images = []*ImageData{image}
	} else {
		images = append(images, image)
	}
	sessionManager.Put(ctx, "images", images)
	return nil
}

func setCurrentImage(ctx context.Context, image *ImageData) {
	sessionManager.Put(ctx, "currentImage", *image)
}

func getCurrentImage(ctx context.Context) (*ImageData, error) {
	val, ok := sessionManager.Get(ctx, "currentImage").(ImageData)
	if !ok {
		return nil, errors.New("type assertion to ImageData failed")
	}
	return &val, nil
}
