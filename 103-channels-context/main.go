/*
		In Go servers, each incoming request is handled in its own goroutine. Request handlers often
	start additional goroutines to access backends such as databases and RPC services. The set of
	goroutines working on a request typically needs access to request-specific values such as the
	identity of the end user, authorization tokens, and the request's deadline. When a request is
	canceled or times out, all the goroutines working on that request should exit quickly so the
	system can reclaim any resources they are using. At Google, we developed a context package
	that makes it easy to pass request-scoped values, cancellation signals, and deadlines across
	API boundaries to all the goroutines involved in handling a request. The package is publicly
	available as context. This article describes how to use the package and provides a complete
	working example.
*/
// https://peter.bourgon.org/blog/2016/07/11/context.html
// https://medium.com/@matryer/context-has-arrived-per-request-state-in-go-1-7-4d095be83bd8

package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()

	fmt.Println("context:\t", ctx)
	fmt.Println("context err:\t", ctx.Err())
	fmt.Printf("context type:\t%T\n", ctx)
	fmt.Println("----------")

	ctx, cancel := context.WithCancel(ctx)

	fmt.Println("context:\t", ctx)
	fmt.Println("context err:\t", ctx.Err())
	fmt.Printf("context type:\t%T\n", ctx)
	fmt.Println("----------")

	cancel()

	fmt.Println("context:\t", ctx)
	fmt.Println("context err:\t", ctx.Err())
	fmt.Printf("context type:\t%T\n", ctx)
	fmt.Println("cancel:\t\t", cancel)
	fmt.Printf("cancel type:\t%T\n", cancel)
	fmt.Println("----------")
}
