
### 作业

我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？


不应该wrap向上抛。

之所以要wrap,是要把错误/异常的上下文返回给上层,由上层决定如何处理。

如果没有wrap,上层不知道底层到底发生了什么,细节是什么。只知道有问题,但具体是什么不清楚,也就无法处理。

sql.ErrNoRows虽然是Err开头的,但其内涵并不表示为错误,或者异常,更多的是一种状态码/标识符,类似于Http Status Code。




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
