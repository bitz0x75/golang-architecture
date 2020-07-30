package main

import (
	"context"
	"fmt"

	"github.com/bitz0x75/golang-architecture/session"
)

func main() {
	ctx := context.Background()
	ctx =session.SetUserID(ctx, 1)
	ctx =session.SetAdmin(ctx, true)
	i := session.GetUserID(ctx)
	b := session.GetAdmin(ctx)

	fmt.Println("User", i, "is admin?", b)

}
