package services

import (
	"context"
	"errors"

	"github.com/alexedwards/scs/v2"
	"github.com/purylte/ocr-webui/types"
)

type ImageService struct {
	sessionManager scs.SessionManager
}

func NewImageService(sm scs.SessionManager) *ImageService {
	return &ImageService{
		sessionManager: sm,
	}
}

func (s *ImageService) ImageIsAllowed(ctx context.Context, imageName string) bool {
	images := s.sessionManager.Get(ctx, "images").([]*types.ImageData)
	for _, image := range images {
		if image.Name == imageName {
			return true
		}
	}
	return false
}

func (s *ImageService) AddAllowedImage(ctx context.Context, image *types.ImageData) error {
	images, ok := s.sessionManager.Get(ctx, "images").([]*types.ImageData)
	if !ok {
		images = []*types.ImageData{image}
	} else {
		images = append(images, image)
	}
	s.sessionManager.Put(ctx, "images", images)
	return nil
}

func (s *ImageService) SetCurrentImage(ctx context.Context, image *types.ImageData) {
	s.sessionManager.Put(ctx, "currentImage", *image)
}

func (s *ImageService) GetCurrentImage(ctx context.Context) (*types.ImageData, error) {
	val, ok := s.sessionManager.Get(ctx, "currentImage").(types.ImageData)
	if !ok {
		return nil, errors.New("type assertion to ImageData failed")
	}
	return &val, nil
}
