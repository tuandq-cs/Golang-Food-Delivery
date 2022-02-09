package uploadgin

import (
	"Golang_Edu/common"
	"Golang_Edu/component/appctx"
	uploadbusiness "Golang_Edu/modules/upload/business"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UploadImage(appCtx appctx.AppContext) func(ctx *gin.Context) {
	return func(context *gin.Context) {
		fileHeader, err := context.FormFile("file")
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		folder := context.DefaultPostForm("folder", "img")

		file, err := fileHeader.Open()
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		defer file.Close() // Close file

		dataBytes := make([]byte, fileHeader.Size)
		if _, err := file.Read(dataBytes); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		biz := uploadbusiness.NewUploadImageBiz(appCtx.UploadProvider())
		img, err := biz.UploadImage(context, dataBytes, folder, fileHeader.Filename)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		img.Fulfill(appCtx.UploadProvider().GetDomain())
		context.JSON(http.StatusOK, common.SimpleSuccessResponse(img))
	}
}
