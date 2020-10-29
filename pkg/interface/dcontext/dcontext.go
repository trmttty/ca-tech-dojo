package dcontext

import "context"

type contextKey string

const userContextKey contextKey = "user-id"

// func WithValue(ctx context.Context, id int) context.Context {
// 	return context.WithValue(ctx, userContextKey, id)
// }
func NewContext(ctx context.Context, id int) context.Context {
	return context.WithValue(ctx, userContextKey, id)
}

// func Extract(ctx context.Context) int {
// 	v := ctx.Value(userContextKey)
// 	if v == nil {
// 		return 0
// 	}
// 	return v.(int)
// }
func FromContext(ctx context.Context) (int, bool) {
	userID, ok := ctx.Value(userContextKey).(int)
	return userID, ok
}
