package repository

import "github.com/imansuparman/ngobrol/internal/app/commons"

// Option anything any repo object needed
type Option struct {
	commons.Options
}

// Repository all repo object injected here
type Repository struct {
	Message IMessageRepo
	Cache   ICacheRepo
}
