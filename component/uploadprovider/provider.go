package uploadprovider

import (
	"Golang_Edu/common"
	"context"
)

type UploadProvider interface {
	SaveFileUploaded(ctx context.Context, data []byte, dst string) (*common.Image, error)
	GetDomain() string
}
