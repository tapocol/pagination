Pagination (Martini Middleware)
===============================

A Martini Middleware in Go that will keep handle limit and page parameters for list pages.

Usage
-----
```go
m.Get("/", pagination.Handler, func(pagi *pagination.Pagination) string {
  limit := pagi.Limit // uint64
  page := pagi.Page // uint64
  offset := pagi.Offset() // uint64
})
```

In mgo:

```go
m.Get("/", pagination.Handler, func(pagi *pagination.Pagination) string {
  ...
  err := query.Skip(int(pagi.Offset())).Limit(int(pagi.Limit)).All(&records)
  ...
})
```

Setup
-----

Download:

```sh
$ go get github.com/craigjackson/pagination
```

Change the config:

```go
pagination.DefaultLimit = uint64(10) // Default limit when parameter missing/invalid
pagination.MinLimit     = uint64(1) // Minimum value for limit
pagination.MaxLimit     = uint64(20) // Maximum value for limit
pagination.LimitParam   = "l" // Expected limit parameter name in http.Request.FormValue
pagination.DefaultPage  = uint64(0) // Default page when parameter missing/invalid
pagination.PageParam    = "p" // Expected page parameter name in http.Request.FormValue
```

