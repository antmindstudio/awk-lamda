package lamda

import (
	"context"

	"github.com/aws/aws-lambda-go/lambda"
)

type Config struct {
	// Context is the base context of the application.
	Context context.Context
	// ResponseEscapeHTML indicates whether the HTML in the response is escaped or not.
	ResponseEscapeHTML bool
}

// setupByConfig setup the application by the specified config
func (app *App) setupByConfig(config Config) {
	if config.Context != nil {
		app.lambdaOptions = append(app.lambdaOptions, lambda.WithContext(config.Context))
	}
	if config.ResponseEscapeHTML {
		app.lambdaOptions = append(app.lambdaOptions, lambda.WithSetEscapeHTML(true))
	}
}
