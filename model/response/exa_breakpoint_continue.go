package response

import "github.com/cxpgo/ginf/model"

type FilePathResponse struct {
	FilePath string `json:"filePath"`
}

type FileResponse struct {
	File model.ExaFile `json:"file"`
}
