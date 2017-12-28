package base

import (
	"SGMS/domain/config"
	"SGMS/domain/face"
	"fmt"
	"image"
	"io"
	"log"
	"os"
	"path"
	"strconv"
	"strings"

	"SGMS/domain/util"

	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"

	"path/filepath"

	"github.com/anthonynsimon/bild/imgio"
	"github.com/anthonynsimon/bild/transform"
)

func NewFileRepo() *FileRepo {
	return new(FileRepo)
}

type FileRepo struct {
}

func getSuffix(fileName string) string {
	lastIndexDot := strings.LastIndex(fileName, ".")
	if lastIndexDot <= 0 {
		return ""
	}
	return fileName[lastIndexDot+1 : len(fileName)]
}

func (this *FileRepo) Open(urlPath string) (*os.File, error) {
	log.Println("open image", this.GetAbsPath(urlPath))
	return os.OpenFile(this.GetAbsPath(urlPath), os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0660)
}
func (this *FileRepo) GenImagePath(fileKey string) (string, string) {
	if fileKey == "" {
		fileKey = util.RandomStr32()
	}
	dir, fullDir := this.GenFileDir(face.FILE_REPO_TYPE_IMAGE, fileKey)
	return path.Join(dir, fileKey+".png"), path.Join(fullDir, fileKey+".png")
}
func (this *FileRepo) GenFileDir(fileType, fileKey string) (string, string) {

	dir := path.Join("/", config.UploadDirName, "/", fileType, "/", fileKey[0:2], "/", fileKey[2:4])
	fullDir := path.Join(config.UploadRootDir, dir)
	stat, err := os.Stat(fullDir)
	if os.IsNotExist(err) || !stat.IsDir() {
		err := os.MkdirAll(fullDir, 0755)
		if err != nil {
			log.Println("mkdir failed dir:"+fullDir+".", err)
			return "", ""
		}
	}
	return dir, fullDir
}
func (this *FileRepo) Save(file face.FileSave, fileType, fileName string) (string, error) {
	log.Println("Save file", fileType, fileName)
	fileKey := util.Md5Reader(file)
	file.Seek(0, os.SEEK_SET)
	// dir := path.Join("/", config.UploadDirName, "/", fileType, "/", fileKey[0:2], "/", fileKey[2:4])
	// fullDir := path.Join(config.UploadRootDir, dir)
	dir, _ := this.GenFileDir(fileType, fileKey)

	suffix := getSuffix(fileName)
	urlPath := path.Join(dir, fileKey+"."+suffix)
	absPath := path.Join(config.UploadRootDir, urlPath)
	newFile, err := os.OpenFile(absPath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0660)
	if nil != err {
		log.Println("open file failed", err)
		return "", err
	}
	defer newFile.Close()
	io.Copy(newFile, file)
	return urlPath, nil

}

func (this *FileRepo) getImageFormat(filename string) imgio.Format {
	suffix := strings.ToLower(filepath.Ext(filename))
	switch suffix {
	case ".jgp":
		fallthrough
	case ".jpeg":
		return imgio.JPEG
	case ".png":
		return imgio.PNG
	}

	return imgio.JPEG
}

func (this *FileRepo) ResizeImage(rpath string, sizes []string, rect *image.Rectangle) (*face.FileSaveImageResult, error) {
	result := &face.FileSaveImageResult{}
	img, err := imgio.Open(this.GetAbsPath(rpath))
	if nil != err {
		log.Println("open image file failed", err)
		return nil, err
	}
	result.RawImage = rpath
	tmpImg := img
	if nil != rect {
		cropImage := transform.Crop(img, *rect)
		result.CroppedImage = SizePath(rpath, rect.Dx(), rect.Dy())
		if err := this.saveImage(this.GetAbsPath(result.CroppedImage), cropImage, this.getImageFormat(rpath)); err != nil {
			panic(err)
		}
		tmpImg = cropImage
	}
	result.ResizedImages = make([]string, len(sizes))
	if 0 < len(sizes) {
		for i, s := range sizes {
			w, h, err := ParseSize(s)
			if err != nil {
				panic(err)
			}
			f := SizePath(rpath, w, h)
			resized := transform.Resize(tmpImg, w, h, transform.Linear)
			log.Println("save size1", this.GetAbsPath(f))
			if err := this.saveImage(this.GetAbsPath(f), resized, imgio.JPEG); err != nil {
				panic(err)
			}
			result.ResizedImages[i] = f
		}
	}
	return result, nil
}

func (this *FileRepo) SaveImage(file face.FileSave, fileName string, sizes []string, rect *image.Rectangle) (*face.FileSaveImageResult, error) {
	rpath, err := this.Save(file, face.FILE_REPO_TYPE_IMAGE, fileName)
	if nil != err {
		log.Println("save file failed", err)
		return nil, err
	}
	return this.ResizeImage(rpath, sizes, rect)
	// log.Println("SaveImage", fileName, sizes, rect)
	// if nil != err {
	// 	log.Println("save file failed", err)
	// 	return nil, err
	// }
	// img, err := imgio.Open(this.GetAbsPath(rpath))
	// if nil != err {
	// 	log.Println("open image file failed", err)
	// 	return nil, err
	// }
	// tmpImg := img
	// result := &face.FileSaveImageResult{}
	// result.RawImage = rpath
	// if nil != rect {
	// 	cropImage := transform.Crop(img, *rect)
	// 	result.CroppedImage = SizePath(rpath, rect.Dx(), rect.Dy())
	// 	log.Println("save rect", result.CroppedImage)
	// 	if err := this.saveImage(this.GetAbsPath(result.CroppedImage), cropImage, imgio.JPEG); err != nil {
	// 		panic(err)
	// 	}
	// 	tmpImg = cropImage
	// }
	// result.ResizedImages = make([]string, len(sizes))
	// if 0 < len(sizes) {
	// 	for i, s := range sizes {
	// 		w, h, err := ParseSize(s)
	// 		if err != nil {
	// 			panic(err)
	// 		}
	// 		f := SizePath(rpath, w, h)
	// 		resized := transform.Resize(tmpImg, w, h, transform.Linear)
	// 		log.Println("save size", f)
	// 		if this.saveImage(this.GetAbsPath(f), resized, imgio.JPEG) != nil {
	// 			panic(err)
	// 		}
	// 		result.ResizedImages[i] = f
	// 	}
	// }

	// return result, nil
}

func (this *FileRepo) GetAbsPath(urlPath string) string {
	return path.Join(config.UploadRootDir, urlPath)
}

func (this *FileRepo) GetRelPath(urlPath string) string {
	return path.Join("/", path.Base(config.UploadDirName), urlPath)
}

func (this *FileRepo) saveImage(filename string, img image.Image, format imgio.Format) error {
	// filename = strings.TrimSuffix(filename, filepath.Ext(filename))

	// switch format {
	// case imgio.PNG:
	// 	filename += ".png"
	// case imgio.JPEG:
	// 	filename += ".jpg"
	// }

	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	return imgio.Encode(f, img, format)
}

func SizePath(path string, width, height int) string {
	dotIndex := strings.LastIndex(path, ".")
	if -1 == dotIndex {
		return fmt.Sprintf("%s.%dx%d", width, height)
	}
	return fmt.Sprintf("%s.%dx%d%s", path[0:dotIndex], width, height, path[dotIndex:])
}

func ParseSize(size string) (int, int, error) {
	ss := strings.Split(size, "x")
	if 0 >= len(ss) {
		return 0, 0, nil
	}
	if 1 == len(ss) {
		i, err := strconv.Atoi(ss[0])
		if err != nil {
			return 0, 0, err
		}
		return i, i, nil
	}
	if 2 <= len(ss) {
		i, err := strconv.Atoi(ss[0])
		if err != nil {
			return 0, 0, err
		}
		j, err := strconv.Atoi(ss[1])
		if err != nil {
			return 0, 0, err
		}
		return i, j, nil
	}
	return 0, 0, nil
}
