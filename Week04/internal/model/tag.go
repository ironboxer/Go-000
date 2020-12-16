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

package model

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
