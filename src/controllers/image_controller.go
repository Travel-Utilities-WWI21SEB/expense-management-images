package controllers

import "context"

type ImageCtl interface {
	UploadImage(ctx context.Context)
	GetImage(ctx context.Context)
}

type ImageController struct {
}

func (ctl *ImageController) UploadImage(ctx context.Context) {

}

func (ctl *ImageController) GetImage(ctx context.Context) {

}
