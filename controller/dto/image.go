package dto

type ImageSearch struct {
	Name        string `json:"name" from:"name" query:"name"`
	IsRecommend uint   `json:"is_recommend" from:"is_recommend" query:"is_recommend"`
	Username    string `json:"username" from:"username" query:"username"`
	TagsID      uint   `json:"tags_id" from:"tags_id" query:"tags_id"`
	CID         uint   `json:"cid" from:"cid" query:"cid"`
}

type ImageSave struct {
	Name        string           `json:"name" from:"name"`
	Url         string           `json:"url" from:"url" `
	Type        string           `json:"type" from:"type"`
	CID         uint             `json:"cid" from:"cid"`
	UID         uint             `json:"uid" from:"uid"`
	TagsID      []ImagesTagsData `json:"tags_id" from:"tags_id"`
	IsRecommend uint             `json:"is_recommend" from:"is_recommend"`
	IsShow      uint8            `json:"is_show" from:"is_show"`
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
