package common

type Paging struct {
	Limit int   `json:"limit" form:"limit"`
	Page  int   `json:"page" form:"page"`
	Total int64 `json:"total"`
}

func (paging *Paging) Preprocess() error {
	if paging.Limit <= 0 || paging.Limit > 100 {
		paging.Limit = 10
	}
	if paging.Page <= 0 {
		paging.Page = 1
	}
	return nil
}

func (paging *Paging) Offset() int {
	return (paging.Page - 1) * paging.Limit
}
