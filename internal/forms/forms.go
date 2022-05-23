package forms

import "github.com/cheeeasy2501/book-library/internal/model"

type FormInterface interface {
	LoadAndValidate() error
}

type Pagination struct {
	Page  uint64 `form:"page" json:"page" binding:"required,gte=1"`
	Limit uint64 `form:"limit" json:"limit" binding:"required,gte=1"`
}

func (pf *Pagination) LoadAndValidate() error {}

type BookRelationships struct {
	model.Relationships
}

func (br *BookRelationships) LoadAndValidate() {
	for index, value := range br.Relations {
		if value != "author" {
			br.Relations = append(br.Relations[:index], br.Relations[index+1])
		}
	}
}
