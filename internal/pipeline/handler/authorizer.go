package handler

import (
	"context"

	"github.com/dadrus/heimdall/internal/heimdall"
)

type Authorizer interface {
	Authorize(context.Context, RequestContext, *heimdall.SubjectContext) error
	WithConfig(config map[string]any) (Authorizer, error)
}
