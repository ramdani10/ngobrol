package service

import (
	"github.com/imansuparman/ngobrol/internal/app/commons"
	"github.com/imansuparman/ngobrol/internal/app/repository"
)

// Option anything any service object needed
type Option struct {
	commons.Options
	*repository.Repository
}

// Services all service object injected here
type Services struct {
	HealthCheck IHealthCheck
	Message     IMessage
}
