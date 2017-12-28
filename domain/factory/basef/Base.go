package basef

import (
	"SGMS/domain/base"
	"SGMS/domain/face"
)

func NewFileRepo() face.IFileRepo {
	return new(base.FileRepo)
}
