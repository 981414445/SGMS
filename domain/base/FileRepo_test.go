package base

import (
	"os"
	"testing"
)

func TestSizePath(t *testing.T) {
	if "hello.1x10.jpg" != SizePath("hello.jpg", 1, 10) {
		t.Error("SizePath error")
	}

}

func TestSaveImage(t *testing.T) {
	repo := new(FileRepo)
	of, _ := os.Open("../../static_resources/image/ab/a1/aba13c58a5a68c54d97ddcef0775fc64.jpg")
	repo.SaveImage(of, "hello.jpg", nil, nil)

}
