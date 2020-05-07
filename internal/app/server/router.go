package server

import (
	"fmt"
	"github.com/imansuparman/ngobrol/internal/app/handler/api"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	cmiddleware "github.com/go-chi/chi/middleware"
	gqlhandler "github.com/graphql-go/graphql-go-handler"
	"github.com/imansuparman/ngobrol/internal/app/commons"
	"github.com/imansuparman/ngobrol/internal/app/handler"
	"github.com/imansuparman/ngobrol/version"
	phttp "github.com/kitabisa/perkakas/v2/http"
	pmiddleware "github.com/kitabisa/perkakas/v2/middleware"
	pstructs "github.com/kitabisa/perkakas/v2/structs"
)

// Router a chi mux
func Router(opt handler.HandlerOption) *chi.Mux {
	handlerCtx := phttp.NewContextHandler(pstructs.Meta{
		Version: version.Version,
		Status:  "stable", //TODO: ask infra if this is used
		APIEnv:  version.Environment,
	})
	commons.InjectErrors(&handlerCtx)

	logMiddleware := pmiddleware.NewHttpRequestLogger(opt.Logger)
	// headerCheckMiddleware := pmiddleware.NewHeaderCheck(handlerCtx, cfg.GetString("app.secret"))

	r := chi.NewRouter()
	// A good base middleware stack (from chi) + middleware from perkakas
	r.Use(cmiddleware.RequestID)
	r.Use(cmiddleware.RealIP)
	r.Use(logMiddleware)
	// r.Use(headerCheckMiddleware) //use this if you want to use default kitabisa's header
	r.Use(cmiddleware.Recoverer)

	// the handler
	phandler := phttp.NewHttpHandler(handlerCtx)

	healthCheckHandler := api.HealthCheckHandler{}
	healthCheckHandler.HandlerOption = opt

	messageHandler := api.MessageHandler{}
	messageHandler.HandlerOption = opt

	indexHandler := handler.IndexHandler{}
	indexHandler.HandlerOption = opt

	// Setup your routing here
	r.Method(http.MethodGet, "/health_check", phandler(healthCheckHandler.HealthCheck))

	// Home
	r.HandleFunc("/", indexHandler.Index)

	// Assets
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	fs := http.FileServer(http.Dir(dir + "/public"))
	r.Handle("/assets/*", http.StripPrefix("/assets/", fs))

	r.Method(http.MethodGet, "/api/v1/messages", phandler(messageHandler.Index))
	r.Method(http.MethodPost, "/api/v1/messages", phandler(messageHandler.Create))

	if opt.Config.GetBool("graphql.is_enabled") {
		route := opt.Config.GetString("graphql.route")
		gqlHandler := gqlhandler.New(&gqlhandler.Config{
			Schema: &opt.GraphqlSchema,
		})
		r.Handle(fmt.Sprintf("/%s", route), gqlHandler)
	}

	return r
}

// TODO: func authRouter()
