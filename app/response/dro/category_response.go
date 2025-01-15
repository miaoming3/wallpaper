package dro

type CategoryListResponse struct {
	Name    string                  `json:"name"`
	ID      uint                    `json:"id"`
	Pid     uint                    `json:"pid"`
	Status  uint                    `json:"status"`
	Sort    uint                    `json:"sort"`
	Childer []*CategoryListResponse `json:"childer,omitempty"`
}
