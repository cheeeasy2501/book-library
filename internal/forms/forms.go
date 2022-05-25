package forms

type FormInterface interface {
	LoadAndValidate() error
}

type Pagination struct {
	Page  uint64 `form:"page" json:"page" binding:"required,gte=1"`
	Limit uint64 `form:"limit" json:"limit" binding:"required,gte=1"`
}

func (pf *Pagination) LoadAndValidate() error {
	return nil
}

func (pf Pagination) GetOffset() uint64 {
	return (pf.Page - 1) * pf.Limit
}
