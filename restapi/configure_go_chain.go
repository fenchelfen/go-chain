// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"gochain/models"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"gochain/restapi/operations"
)

//go:generate swagger generate server --target ../../gochain --name GoChain --spec ../swagger/swagger.yml --principal interface{}

func configureFlags(api *operations.GoChainAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.GoChainAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.UseSwaggerUI()
	// To continue using redoc as your UI, uncomment the following line
	// api.UseRedoc()

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	if api.MineHandler == nil {
		api.MineHandler = operations.MineHandlerFunc(func(params operations.MineParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.Mine has not yet been implemented")
		})
	}
	if api.AddTransactionHandler == nil {
		api.AddTransactionHandler = operations.AddTransactionHandlerFunc(func(params operations.AddTransactionParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.AddTransaction has not yet been implemented")
		})
	}
	if api.RegisterNodeHandler == nil {
		api.RegisterNodeHandler = operations.RegisterNodeHandlerFunc(func(params operations.RegisterNodeParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.RegisterNode has not yet been implemented")
		})
	}
	if api.RegisterWithNodeHandler == nil {
		api.RegisterWithNodeHandler = operations.RegisterWithNodeHandlerFunc(func(params operations.RegisterWithNodeParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.RegisterWithNode has not yet been implemented")
		})
	}

	api.GetChainHandler = operations.GetChainHandlerFunc(
		func(params operations.GetChainParams) middleware.Responder {

			var length int64 = 4

			var _peers = []*models.Peer{{}}
			var chain = []*models.Block{{}}

			body := models.Chain{
				Chain:  chain,
				Length: &length,
				Peers:  _peers,
			}

			return operations.NewGetChainOK().WithPayload(&body)
		},
	)

	api.PreServerShutdown = func() {}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
