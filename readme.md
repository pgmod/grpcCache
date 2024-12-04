Сервер работает на порту 50051


### пример клиента

``` golang
package main

import (
	"context"
	"os"
	"time"

	"github.com/pgmod/grpcCache/client"
)

func main() {

	target := os.Getenv("GRPC_ADDR")
	if target == "" {
		target = "localhost:50051"
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	client := client.CreateClient(target)
	client.Clear(ctx, "100000")

}
```
