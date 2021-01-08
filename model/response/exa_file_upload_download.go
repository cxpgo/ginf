package response

import "github.com/cxpgo/ginf/model"

type ExaFileResponse struct {
	File model.ExaFileUploadAndDownload `json:"file"`
}
