# in-memory-cache-go
Golang tool for storing data in the system RAM

My first experience as a development package for storing data of the in-memory-cache type

## Example

```go
package main

import (
  "fmt"
  "github.com/d1mpi/in-memory-cache-go"
)

func main() {
	newCache := cache.New()
	newCache.Set("Alina", 12)

	userId := newCache.Get("Alina")
	fmt.Println(userId)

	newCache.Delete("Alina")

	userId = newCache.Get("Alina")
	fmt.Println(userId)
}
```

## Install
Use go get -u github.com/d1mpi/in-memory-cache-go
