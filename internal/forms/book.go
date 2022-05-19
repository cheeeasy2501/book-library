package forms

type GetBook struct {
	Id uint64 `uri:"id" binding:"required"`
}
