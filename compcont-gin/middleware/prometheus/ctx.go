package prometheus

import (
	"context"
	"net/http"
)

type ctxKeyAPIName struct{}

func MarkAPINameToContext(req *http.Request, apiName string) {
	ctx := req.Context()
	if v, ok := ctx.Value(ctxKeyAPIName{}).(string); ok && v == apiName {
		return
	}
	*req = *req.WithContext(context.WithValue(ctx, ctxKeyAPIName{}, apiName))
}

func apiNameFromContext(ctx context.Context) (string, bool) {
	v, ok := ctx.Value(ctxKeyAPIName{}).(string)
	return v, ok
}
