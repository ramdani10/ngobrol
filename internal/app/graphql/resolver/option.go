package resolver

import (
	"github.com/imansuparman/ngobrol/internal/app/service"
)

type Options func(*Resolver)

func WithServices(svc *service.Services) Options {
	return func(r *Resolver) {
		r.svc = svc
	}
}
