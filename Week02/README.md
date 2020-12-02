
### 作业

我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？


不应该wrap向上抛。

之所以要wrap,是要把错误/异常的上下文返回给上层,由上层决定如何处理。

如果没有wrap,上层不知道底层到底发生了什么,细节是什么。只知道有问题,但具体是什么不清楚,也就无法处理。

sql.ErrNoRows虽然是Err开头的,但其内涵并不表示为错误,或者异常,更多的是一种状态码/标识符,类似于Http Status Code。

---

下面是sql.ErrNoRows官网的定义, Scan方法返回sql.ErrNoRows表示返回的查询结果为空,或者有0条记录。

对于一个查询结果而言, 返回0条记录是完全合法的, 为什么要对0有偏见呢?如果返回0条需要设置一个特别的错误来表示, 那么返回1条记录是否要另外设置一个
新的错误sql.ErrOneRows?那么返回两条记录呢?sql.ErrTwoRows?

这里的sql.ErrNoRows只是为了方便判断Scan这个方法有没有操作成功,超过这个范畴,便不再有意义.

如果wrap之后向上层传递,反而暴露了底层的实现,可能会形成依赖。

另外, 对于sql.ErrNoRows这个错误, 其上下文信息不用通过wrap就可以判断。

sql语句编译期就可以确定,传入的参数是从上层来的,上层可以通过日志等方式记录这些参数,一旦报错,上层有完整的信息,不需要将dao层的信息再包裹一次。

---

https://golang.org/pkg/database/sql/#Row


ErrNoRows is returned by Scan when QueryRow doesn't return a row. In such a case, QueryRow returns a placeholder *Row value that defers this error until a Scan.

```golang
var ErrNoRows = errors.New("sql: no rows in result set")
```

### func (*DB) QueryRow

```golang
func (db *DB) QueryRow(query string, args ...interface{}) *Row
```

QueryRow executes a query that is expected to return at most one row. QueryRow always returns a non-nil value. Errors are deferred until Row's Scan method is called. If the query selects no rows, the *Row's Scan will return ErrNoRows. Otherwise, the *Row's Scan scans the first selected row and discards the rest.


### func (*Row) Scan

```golang
func (r *Row) Scan(dest ...interface{}) error
```

Scan copies the columns from the matched row into the values pointed at by dest. See the documentation on Rows.Scan for details. If more than one row matches the query, Scan uses the first row and discards the rest. If no row matches the query, Scan returns ErrNoRows.

---


对于"查询结果为空", 更好的方式为返回nil, 或者返回一个长度为0的切片, 或者dao层设置一个统一的表示查询结果为空的标识符,或者或者接口,
通过接口/方法来表示该查询结果是否为空。


代码:

依赖
```golang
main -> service -> dao
```


dao/models.go中通过对外暴露```IsDBConnectError```和```IsEmptyQueryResult```分别表示dao层的报错是否为数据库连接问题
还是查询结果为空的问题


service/service.go模拟具体的业务逻辑,对于dao层的报错向上层传递

main.go中```/tags/:id```handler中根据service层返回的异常通过dao层暴露的方法判断其为数据库连接报错还是查询结果为空。


dao/models.go
```go
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
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"strings"
)

type dbConnectError interface {
	DBConnectionFailed() bool
}

type DBConnectionError struct {
	errMsg string
}

func (db DBConnectionError) DBConnectionFailed() bool {
	return strings.Contains(db.errMsg, "DB")
}

func (db DBConnectionError) Error() string {
	return db.errMsg
}

func IsDBConnectError(err error) bool {
	te, ok := err.(dbConnectError)
	return ok && te.DBConnectionFailed()
}

type emptyQueryResult interface {
	EmptyQueryResult() bool
}

type EmptyQueryResultError struct {
	errMsg string
}

func (e EmptyQueryResultError) Error() string {
	return e.errMsg
}

func (e EmptyQueryResultError) EmptyQueryResult() bool {
	return strings.Contains(e.errMsg, "No")
}

func IsEmptyQueryResult(err error) bool {
	te, ok := err.(emptyQueryResult)
	return ok && te.EmptyQueryResult()
}

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
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/blog")
	if err != nil {
		//return nil, errors.New("DB Service Temporary Unavailable")
		return nil, DBConnectionError{errMsg: "DB Service Temporary Unavailable"}
	}
	defer db.Close()
	var tag Tag
	err = db.QueryRow("SELECT id, name FROM blog_tag where id = ?", id).Scan(&tag.ID, &tag.Name)
	if err != nil {
		return nil, EmptyQueryResultError{errMsg: fmt.Sprintf("No Such Tag(%d)", id)}
	}
	return &tag, nil
}
```



service/service.go
```go
package service

import "github.com/lttzzlll/week02/dao"

// TagService handles biz for Tag Model
type TagService struct {
}

// GetTagByID returns Tag by id
func (s *TagService) GetTag(id uint64) (*dao.Tag, error) {
	return dao.GetTag(id)
}

```

main.go

```go
package main

import (
	"github.com/gin-gonic/gin"
	"github.com/lttzzlll/week02/dao"
	"log"
	"net/http"
	"strconv"

	"github.com/lttzzlll/week02/service"
)

func main() {
	var s *service.TagService
	r := gin.Default()
	r.GET("/tags/:id", func(c *gin.Context) {
		id := c.Param("id")
		tagID, err := strconv.Atoi(id)
		if err != nil {
			c.String(http.StatusBadRequest, "Invalid Tag ID %s", id)
			return
		}

		tag, err := s.GetTag(uint64(tagID))
		if err != nil {
			var errMsg = err.Error()
			var returnCode int
			if dao.IsDBConnectError(err) {
				returnCode = http.StatusServiceUnavailable
			} else if dao.IsEmptyQueryResult(err) {
				returnCode = http.StatusNotFound
			} else {
				returnCode = http.StatusInternalServerError
			}
			c.String(returnCode, errMsg)
			log.Printf("Error: %s, Status: %d", errMsg, returnCode)
			return
		}
		c.JSON(http.StatusOK, gin.H{"id": tag.ID, "name": tag.Name})
	})
	r.Run()
}

```