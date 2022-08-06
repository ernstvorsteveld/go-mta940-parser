package sta

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"mime/multipart"
	"mta-service/common"
	"net/http"
	"os"
	"path/filepath"
)

var ch chan common.FilenameMessage

func Register(router *gin.RouterGroup, c chan common.FilenameMessage) {
	FileUploadRegister(router)
	ch = c
}
func FileUploadRegister(router *gin.RouterGroup) {
	router.POST("/upload", FileUpload)
}

type BindFile struct {
	Name  string                `form:"name" binding:"required"`
	File  *multipart.FileHeader `form:"file" binding:"required"`
}

func FileUpload(c *gin.Context) {
	var bindFile BindFile

	// Bind file
	if err := c.ShouldBind(&bindFile); err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("err: %s", err.Error()))
		return
	}

	// Save uploaded file
	file := bindFile.File
	dst := os.Getenv("HOME") + "/" + filepath.Base(file.Filename)
	if err := c.SaveUploadedFile(file, dst); err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
		return
	}

	c.String(http.StatusOK,
		fmt.Sprintf("File %s uploaded successfully with name=%s.", file.Filename, bindFile.Name))

	ch <- common.FilenameMessage{
		Filename : bindFile.Name,
		Dst : dst,
	}
}