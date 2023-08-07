package appctx

import (
	"shopbee/component/uploadprovider"
)

type AppContext interface {
	UploadProvider() uploadprovider.UploadProvider
}

type appCtx struct {
	uploadProvider uploadprovider.UploadProvider
}

func NewAppContext(uploadProvider uploadprovider.UploadProvider) *appCtx {
	return &appCtx{
		uploadProvider: uploadProvider,
	}
}

func (ctx *appCtx) UploadProvider() uploadprovider.UploadProvider {
	return ctx.uploadProvider
}
