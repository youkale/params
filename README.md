Params
--- 
[![Badge](https://img.shields.io/badge/link-996.icu-%23FF4D5B.svg?style=flat-square)](https://996.icu/#/en_US)
[![LICENSE](https://img.shields.io/badge/license-Anti%20996-blue.svg?style=flat-square)](https://github.com/996icu/996.ICU/blob/master/LICENSE)

Params is Go library for convert url.Values to struct, support Http request query-parameters, form-parameters.

----

### Installation
Use `go get` to install:

```go
 go get github.com/youkale/params
```

Usage
1. Define a struct with tags that match the query parameter names:

```go
type User struct {
    UserId  int64   `param:"user_id,100"`
    StoreId int     `param:"store_id"`
    Page    float32 `param:"page"`
    Name    string  `param:"name"`
    Age     uint8   `param:"age,18"`
    Enable  bool    `param:"enable,false"`
}

```

2. In your HTTP handler, parse the query parameters into an instance of the struct:

```go
import "github.com/youkale/params"

func MyHandler(w http.ResponseWriter, r *http.Request) {
    // or convert request.Form
    var user User
    if err := params.Convert(r.URL.Query(), &user); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
    return
    }
        // Do something with the user struct...
}
```

Params will automatically parse the query parameters and set the values in the struct fields.


### Tag Options

- `param:"store_id,store_01"`: The following options can be included in the `param` tag.
- `store_id`: Indicates the key to get the content of the field.
- `store_01`: specifies a default value for the field if the query parameter is not present.


### Performance ###

```
    goos: linux
    goarch: amd64
    pkg: github.com/youkale/params
    2000000000	         0.00 ns/op
    PASS
```

## License ##

[![LICENSE](https://img.shields.io/badge/license-Anti%20996-blue.svg)](https://github.com/996icu/996.ICU/blob/master/LICENSE)
