# rakuten-api-client

## Usage

```
go get github.com/mutao-net/rakuten-api-client
```

```go
import 	"github.com/mutao-net/rakuten-api-client"
params := rakuten.QueryParameters{
	ApplicationID: "${ApplicationID}",
	Keyword:       "golang",
	Sort:          "+reviewAverage",
}
result := rakuten.GetRakutenItems(params)
```