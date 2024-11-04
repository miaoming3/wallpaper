package dto

type ImageSearch struct {
	Name        string `json:"name" form:"name" query:"name" `
	IsRecommend uint   `json:"is_recommend" form:"is_recommend" query:"is_recommend"`
	Username    string `json:"username" form:"username" query:"username"`
	TagsID      uint   `json:"tags_id" form:"tags_id" query:"tags_id" url:"tags_id"`
	CID         uint   `json:"cid" form:"cid" query:"cid"`
}

type ImageSave struct {
	Name        string           `json:"name" form:"name"`
	Url         string           `json:"url" form:"url" `
	Type        string           `json:"type" form:"type"`
	CID         uint             `json:"cid" form:"cid"`
	UID         uint             `json:"uid" form:"uid"`
	TagsID      []ImagesTagsData `json:"tags_id" form:"tags_id"`
	IsRecommend uint             `json:"is_recommend" form:"is_recommend"`
	IsShow      uint8            `json:"is_show" form:"is_show"`
}

type ImagesTagsData struct {
	ID uint `json:"id" from:"ID"`
}

type ImageUpdate struct {
	ID uint `json:"id" from:"id"`
	ImageSave
}

type ImageDelete struct {
	ID uint `json:"id" from:"id" query:"id"`
}
