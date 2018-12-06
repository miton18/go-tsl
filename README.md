GO-TSL Client
=============

[![Go Report Card](https://goreportcard.com/badge/github.com/miton18/go-tsl)](https://goreportcard.com/report/github.com/miton18/go-tsl)


Example:

```go
import (
	"fmt"
	"time"

	"github.com/miton18/go-tsl"
)

func main() {
	c := tsl.NewClient("https://tsl.gra1-ovh.metrics.ovh.net", "TOKEN", nil)
	q := c.NewQuery()

	q.
		Select("os.cpu").
		Where("host", tsl.Eq("serverA")).
		Last(5*time.Minute, tsl.NilTime)

	fmt.Println(q.Dump())

	r, err := q.Execute()
	if err != nil {
		panic(err.Error())
	}

	fmt.Println(fmt.Sprintf("%+v", r))
}
```
