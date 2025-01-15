package dro

type ImagesResponse struct {
	Url         string               `json:"url"`
	Path        string               `json:"path"`
	Name        string               `json:"name"`
	Type        uint                 `json:"type"`
	IsRecommend uint                 `json:"is_recommend"`
	Cid         uint                 `json:"cid"`
	User        UserResponse         `json:"user"`
	Tags        *[]TagsResponse      `json:"tags"`
	Category    CategoryListResponse `json:"category"`
}
