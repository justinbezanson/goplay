# Go + Gin + Templ Webapp Setup

## Prerequisites

- Go 1.22+
- [templ](https://github.com/a-h/templ) CLI: `go install github.com/a-h/templ/cmd/templ@latest`

## Setup Steps

### 1. Initialize the module

```bash
go mod init goplay
```

### 2. Install dependencies

```bash
go get github.com/gin-gonic/gin
go get github.com/a-h/templ
```

### 3. Create the templ component

**`views/hello.templ`**:

```
package views

templ Hello() {
	<h1>Hello, World!</h1>
}
```

Generate Go code from templ:

```bash
templ generate
```

### 4. Create the main server

**`main.go`**:

```go
package main

import (
	"github.com/gin-gonic/gin"
	"goplay/views"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "", nil)
		c.Render(200, views.Hello())
	})

	r.Run(":8080")
}
```

### 5. Run the app

```bash
templ generate && go run .
```

Visit `http://localhost:8080` — you should see "Hello, World!" rendered.
