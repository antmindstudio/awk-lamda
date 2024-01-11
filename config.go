package lamda

import (
	"context"

	"github.com/aws/aws-lambda-go/lambda"
)

type Config struct {
	// Context is the base context of the application.
	Context context.Context
	// ContextValues are the values in the handler context. If a base context was set, the context is
	// used as the parent.
	ContextValues map[any]any
	// ResponseEscapeHTML indicates whether the HTML in the response is escaped or not.
	ResponseEscapeHTML bool
	// SIGTERMCallbacks are the callback functions to run on SIGTERM, and it will enable the Lambda
	// SIGTERM behavior if the callbacks are not empty.
	SIGTERMCallbacks []func()
}

// setupByConfig setup the application by the specified config
func (app *App) setupByConfig(config Config) {
	if config.Context != nil {
		app.lambdaOptions = append(app.lambdaOptions, lambda.WithContext(config.Context))
	}
	if len(config.ContextValues) > 0 {
		for k, v := range config.ContextValues {
			app.lambdaOptions = append(app.lambdaOptions, lambda.WithContextValue(k, v))
		}
	}
	if config.ResponseEscapeHTML {
		app.lambdaOptions = append(app.lambdaOptions, lambda.WithSetEscapeHTML(true))
	}
	if len(config.SIGTERMCallbacks) > 0 {
		app.lambdaOptions = append(app.lambdaOptions, lambda.WithEnableSIGTERM(config.SIGTERMCallbacks...))
	}
}
