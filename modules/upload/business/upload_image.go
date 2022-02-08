package uploadbusiness

import (
	"Golang_Edu/common"
	"Golang_Edu/component/uploadprovider"
	"bytes"
	"context"
	"errors"
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"path/filepath"
	"strings"
	"time"
)

type uploadImageBiz struct {
	provider uploadprovider.UploadProvider
}

func NewUploadImageBiz(provider uploadprovider.UploadProvider) *uploadImageBiz {
	return &uploadImageBiz{provider: provider}
}

func (biz *uploadImageBiz) UploadImage(ctx context.Context,
	data []byte, folder, fileName string) (*common.Image, error) {

	fileBytes := bytes.NewBuffer(data)
	// Check whether data is an image
	w, h, err := getImageDimension(fileBytes)
	if err != nil {
		return nil, err
	}

	// Manipulate folder string
	folder = handleFolderString(folder)
	// Manipulate fileName
	fileName, fileExt := handleFileNameString(fileName)
	img, err := biz.provider.SaveFileUploaded(ctx, data, fmt.Sprintf("%s/%s", folder, fileName))
	if err != nil {
		return nil, errors.New("can not save image")
	}
	// Pass value for Image instance
	img.Width = w
	img.Height = h
	img.Extension = fileExt
	return img, nil
}

func handleFileNameString(fileName string) (string, string) {
	fileExt := filepath.Ext(fileName)                                   // .jpg, .png, ...
	return fmt.Sprintf("%d%s", time.Now().UnixNano(), fileExt), fileExt // 9232894234.png
}

func handleFolderString(folder string) string {
	if strings.TrimSpace(folder) == "" {
		folder = "img"
	}
	return folder
}

func getImageDimension(reader io.Reader) (int, int, error) {
	img, _, err := image.DecodeConfig(reader)
	if err != nil {
		return 0, 0, err
	}
	return img.Width, img.Height, nil
}
