// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"customlogs/models"
	"customlogs/restapi/operations"
	"customlogs/restapi/operations/log_preserve"
	"customlogs/restapi/operations/system"
	"customlogs/services"
	"customlogs/utility"
)

//go:generate swagger generate server --target ../../customlogs --name CutomLogs --spec ../swagger.yaml --principal interface{}

func configureFlags(api *operations.CutomLogsAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.CutomLogsAPI) http.Handler {
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

	// Applies when the "apikey" header is set
	api.APIKeyAuthAuth = func(token string) (interface{}, error) {
		return services.IsAuthentic(token)
	}

	if api.SystemGetChecklifeHandler == nil {
		api.SystemGetChecklifeHandler = system.GetChecklifeHandlerFunc(func(params system.GetChecklifeParams) middleware.Responder {
			return middleware.NotImplemented("operation system.GetChecklife has not yet been implemented")
		})
	} else {
		api.SystemGetChecklifeHandler = system.GetChecklifeHandlerFunc(
			func(params system.GetChecklifeParams) middleware.Responder {
				status := "Up"
				uptime := fmt.Sprintf("%vm", time.Since(utility.Configuration.StartTime).Minutes())
				response := models.System{
					// BuildVersion: &utility.Configuration.BuildVersion,
					Status: &status,
					Uptime: &uptime,
					// Version:      &utility.Configuration.Version,
				}

				return system.NewGetChecklifeOK().WithPayload(&response)
			})
	}
	if api.LogPreserveGetPreserveHandler == nil {
		api.LogPreserveGetPreserveHandler = log_preserve.GetPreserveHandlerFunc(func(params log_preserve.GetPreserveParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation log_preserve.GetPreserve has not yet been implemented")
		})
	} else {

	}

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
// scheme value will be set accordingly: "http", "https" or "unix".
func configureServer(s *http.Server, scheme, addr string) {
	config := utility.Configuration
	utility.InitLogger(config.Country)
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation.
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics.
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return customMiddleware(handler)
}

func customMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		enableCors(&w)

		// utility.Log.Infof("Received request: [%s] %v", r.Method, r.URL)
		next.ServeHTTP(w, r)
	})
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}
