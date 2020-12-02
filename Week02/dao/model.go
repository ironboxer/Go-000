/*
mysql> desc blog_tag;
+-------------+------------------+------+-----+---------+----------------+
| Field       | Type             | Null | Key | Default | Extra          |
+-------------+------------------+------+-----+---------+----------------+
| id          | int unsigned     | NO   | PRI | NULL    | auto_increment |
| name        | varchar(100)     | YES  |     |         |                |
| created_on  | int unsigned     | YES  |     | 0       |                |
| created_by  | varchar(100)     | YES  |     |         |                |
| modified_on | int unsigned     | YES  |     | 0       |                |
| modified_by | varchar(100)     | YES  |     |         |                |
| deleted_on  | int unsigned     | YES  |     | 0       |                |
| state       | tinyint unsigned | YES  |     | 1       |                |
+-------------+------------------+------+-----+---------+----------------+
8 rows in set (0.01 sec)
*/

package dao

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

// Tag represents Blog Tag
type Tag struct {
	ID         uint64
	Name       string
	CreatedOn  uint64
	ModifiedOn uint64
	ModifiedBy string
	DeletedOn  uint64
	State      bool
}

// GetTag returns Tag by ID
func GetTag(id uint64) (*Tag, error) {
	db, err := sql.Open("mysql", "r")
	if err != nil {
		return nil, errors.New("DB Service Temporary Unavailable")
	}
	defer db.Close()
	var tag Tag
	err = db.QueryRow("SELECT id, name FROM blog_tag where id = ?", id).Scan(&tag.ID, &tag.Name)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("No Such Tag(%d)", id))
	}
	return &tag, nil
}
