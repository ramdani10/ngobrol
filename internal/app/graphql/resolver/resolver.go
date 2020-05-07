package resolver

import (
	"sync"

	"github.com/imansuparman/ngobrol/internal/app/service"
)

var (
	resolverInst *Resolver
	once         = new(sync.Once)
)

// Resolver struct for resolver instance
type Resolver struct {
	svc *service.Services
}

func init() {
	once.Do(func() {
		resolverInst = new(Resolver)
	})
}

func Init(opts ...Options) {
	for _, opt := range opts {
		opt(resolverInst)
	}
}
