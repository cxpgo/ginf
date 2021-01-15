package model

import (
	"github.com/cxpgo/golib/model"
)

// file struct, 文件结构体
type ExaFile struct {
	model.GormModel
	FileName     string
	FileMd5      string
	FilePath     string
	ExaFileChunk []ExaFileChunk
	ChunkTotal   int
	IsFinish     bool
}

// file chunk struct, 切片结构体
type ExaFileChunk struct {
	model.GormModel
	ExaFileID       uint
	FileChunkNumber int
	FileChunkPath   string
}
