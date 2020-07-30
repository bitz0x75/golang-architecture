package session

import (
	"context"
)

type stringKey string
type intKey int

var userID stringKey
var admin intKey

func SetUserID(ctx context.Context, uID int) context.Context {
	return context.WithValue(ctx, userID, uID)
}

func SetAdmin(ctx context.Context, isAdmin bool) context.Context {
	return context.WithValue(ctx, admin, isAdmin)
}

func GetUserID(ctx context.Context) int {
	//value assertion
	if uid := ctx.Value(userID); uid != nil {
		if i, ok := uid.(int); ok {
			return i
		}
	}
	return 0
}

func GetAdmin(ctx context.Context) bool {
	b := ctx.Value(admin)
	if v, ok := b.(bool); ok {
		return v
	}
	return false
}

// ctx := context.WithValue(context.Background(), "userID", 12345)
// 	ctx = context.WithValue(ctx, 1, "admin")
// 	if v := ctx.Value("userID"); v != nil {
// 		fmt.Println(v)
// 	} else {
// 		fmt.Println("no value associated with that key")
// 	}

// 	if v := ctx.Value(1); v != nil {
// 		fmt.Println(v)
// 	} else {
// 		fmt.Println("no value associated with that key")
// 	}

// 	if v := ctx.Value(2); v != nil {
// 		fmt.Println(v)
// 	} else {
// 		fmt.Println("no value associated with that key")
// 	}
