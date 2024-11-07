package dto

type ImageSearch struct {
	Name        string `json:"name" form:"name" query:"name" `
	IsRecommend uint   `json:"is_recommend" form:"is_recommend" query:"is_recommend"`
	Username    string `json:"username" form:"username" query:"username"`
	TagsID      uint   `json:"tags_id" form:"tags_id" query:"tags_id" url:"tags_id"`
	CID         uint   `json:"cid" form:"cid" query:"cid"`
}

type ImageSave struct {
	Name        string `json:"name" form:"name" binding:"required"`
	Url         string `json:"url" form:"url"  `
	Path        string `json:"path" form:"path" binding:"required"`
	Type        uint8  `json:"type" form:"type" `
	CID         uint   `json:"cid" form:"cid" binding:"required"`
	Tags        []uint `json:"tags_id" form:"tags_id"`
	IsRecommend uint8  `json:"is_recommend" form:"is_recommend"`
	IsShow      uint8  `json:"is_show" form:"is_show"`
}

type ImageUpdate struct {
	ID uint `json:"id" from:"id" binding:"required"`
	ImageSave
}

type ImageDelete struct {
	ID uint `json:"id" from:"id" query:"id"`
}
