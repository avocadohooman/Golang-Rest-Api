package postgres


import (
	//by using _ you say you won't declare any variables/function from the library, only import it	
	_"github.com/lib/pq" // standard postgresql GO library
	
	"github.com/go-pg/pg/v9" //PostgreSQL client and ORM for Golang
)

func New(opts *pg.Options) *pg.DB {
	db := pg.Connect(opts)

	return db
}