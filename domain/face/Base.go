package face

import (
	"image"
	"io"
	"os"
)

type FileSaveImageResult struct {
	RawImage, CroppedImage string
	ResizedImages          []string
}

type FileSave interface {
	io.Reader
	io.Seeker
	// io.Closer
}
type IFileRepo interface {
	Save(file FileSave, fileType, fileName string) (string, error)
	// 保存图片
	SaveImage(file FileSave, fileName string, sizes []string, rect *image.Rectangle) (*FileSaveImageResult, error)
	// 将图片转换为三种尺寸
	ResizeImage(rpath string, sizes []string, rect *image.Rectangle) (*FileSaveImageResult, error)
	GetAbsPath(urlPath string) string
	GetRelPath(absPath string) string
	Open(urlPath string) (*os.File, error)
}

const (
	FILE_REPO_TYPE_IMAGE = "image"
	FILE_REPO_TYPE_BIN   = "bin"
	FILE_REPO_TYPE_DOC   = "doc"
)
