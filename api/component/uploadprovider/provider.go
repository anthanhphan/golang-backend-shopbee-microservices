package uploadprovider

import (
	"context"
	"shopbee/common"
)

type UploadProvider interface {
	SaveFileUploaded(context context.Context, data []byte, dst string) (*common.Image, error)
}
