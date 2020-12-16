package dao

import (
	"context"
	"database/sql"
	"github.com/google/wire"

	"fmt"
	"github.com/pkg/errors"

	"github.com/ironboxer/week04/internal/model"
)

var NotFound = errors.New("Not Found")

type Dao interface {
	GetTag(ctx context.Context, id uint64) (*model.Tag, error)
}

type dao struct {
	db *sql.DB
}

func (d *dao) GetTag(ctx context.Context, id uint64) (*model.Tag, error) {
	var tag model.Tag
	err := d.db.QueryRowContext(ctx, "SELECT id, name FROM blog_tag where id = ?", id).Scan(&tag.ID, &tag.Name)
	if err != nil {
		return nil, errors.Wrap(NotFound, fmt.Sprintf("Tag %v", id))
	}
	return &tag, nil
}

func NewDao(db *sql.DB) Dao {
	return &dao{db: db}
}

func NewDB() (*sql.DB, func(), error) {
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/blog")
	cleanup := func() {
		if err == nil {
			db.Close()
		}
	}
	return db, cleanup, err
}

var Provider = wire.NewSet(NewDB, NewDao)