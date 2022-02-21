# go-program-tuning
Simple tuning work for go programs in high concurrency scenarios.

### Installation

Run the following command under your project:

> go get -u github.com/NICEXAI/go-program-tuning

### Usage

Just import it directly, such as:

```go
package main

import (
	_ "github.com/NICEXAI/go-program-tuning"
	"log"
)

func main() {
	log.Println("server is running")
}
```