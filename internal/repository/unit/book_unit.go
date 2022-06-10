package unit

import (
	"database/sql"
	"github.com/cheeeasy2501/book-library/internal/model"
)

Нужно написать обертку вокруг nat.db
func (c *Client) Session(ctx context.Context) (context.Context, func(error), error) {
	if value := ctx.Value(SessionTxKey{}); value != nil {
		if v, ok := value.(*sql.Tx); ok {
			return ctx, customFinisher(v), nil
		}
	}
	tx, err := c.db.BeginTx(ctx, nil)
	if err != nil {
		return ctx, customFinisher(nil), err
	}

	ctx = context.WithValue(ctx, SessionTxKey{}, tx)
	return ctx, customFinisher(tx), nil
}

func customFinisher(tx *sql.Tx) func(error) {
	return func(err error) {
		if tx == nil {
			return
		}
		switch err {
		case nil:
			tx.Commit()
		default:
			tx.Rollback()
		}
	}
}


func (s *ImagesAggregateService) Create(ctx context.Context, user *model.User, image *model.Image) (err error) {
	ctx, finish, err = s.db.Session(ctx)
	if err != nil {
		return err
	}
	defer func() {
		finish(err)
	}()
	err = s.users.Create(ctx, user)
	if err != nil {
		return
	}
	err = s.images.Create(ctx, image)
	if err != nil {
		return
	}
	return
}
//type UnitOfWork interface {
//	Save() *UnitError
//}
//
//type BookUnit struct {
//	errs             []error
//	tx               *sql.Tx
//	BookRepository   *repository.BookRepository // change to interface
//	AuthorRepository repository.AuthorRepoInterface
//}
//
//func NewBookUnit(repos *repository.Repository) *BookUnit {
//	return &BookUnit{
//		//BookRepository: repository.NewBookRepository(),
//	}
//}
//
//func (u *BookUnit) Save() *UnitError {
//	unitError := NewUnitError()
//
//	if len(u.errs) == 0 {
//		err := u.tx.Commit()
//		if err != nil {
//			unitError.Message = CommitErrorMessage
//			unitError.errs = append(unitError.errs, err)
//			return unitError
//		}
//
//		return nil
//	}
//
//	unitError.Message = UnitErrorMessage
//	unitError.errs = append(unitError.errs, u.errs...)
//	err := u.tx.Rollback()
//	if err != nil {
//		unitError.Message = RollbackErrorMessage
//		unitError.errs = append(unitError.errs, err)
//		return unitError
//	}
//
//	return unitError
//}
//
//const (
//	UnitErrorMessage     = "Unit Error"
//	RollbackErrorMessage = "Rollback Error"
//	CommitErrorMessage   = "Commit Error"
//)
//
//type UnitError struct {
//	errs    []error
//	Message string
//}
//
//func NewUnitError() *UnitError {
//	return &UnitError{}
//}
