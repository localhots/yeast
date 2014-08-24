package sleep

import (
	"fmt"
	"time"

	"code.google.com/p/go.net/context"
)

func Call(ctx context.Context) context.Context {
	fmt.Println("Going into sleep")
	time.Sleep(5 * time.Second)
	fmt.Println("Woke up")
	return ctx
}
