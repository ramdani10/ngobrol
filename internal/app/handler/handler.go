package handler

import (
	"github.com/graphql-go/graphql"
	"github.com/imansuparman/ngobrol/internal/app/commons"
	"github.com/imansuparman/ngobrol/internal/app/service"
)

// HandlerOption option for handler, including all service
type HandlerOption struct {
	commons.Options
	*service.Services
	GraphqlSchema graphql.Schema
}
