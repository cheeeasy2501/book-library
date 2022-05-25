package model

type BookAggregate struct {
	Book
	Authors []Author `json:"authors,omitempty"`
}

func (a *BookAggregate) Columns() string {
	author := Author{}
	columns := a.Book.Columns() + ", " + author.Columns()
	//+ a.Authors[0].Columns()
	return columns
}

func (a *BookAggregate) Fields() []interface{} {
	//author := Author{}
	//fields := []interface{}{&a.Id, &a.Title, &a.Description, &a.Link, &a.InStock, &a.CreatedAt, a.UpdatedAt}
	//return append(fields, author.Fields()...)
	return a.Book.Fields()
}

func (a *BookAggregate) Scan(src interface{}) error {
	var err error
	switch src.(type) {

	}

	return err
}
