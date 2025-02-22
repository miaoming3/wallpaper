package dto

import "mime/multipart"

type UploadFile struct {
	File multipart.FileHeader `form:"file"`
}

type UploadFileMust struct {
	File []*UploadFile `form:"file"`
}
