package lamda

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// App is the application to run a HTTP service with AWS Lambda and AWS APIGateway.
type App struct {
	lambdaOptions []lambda.Option
}

// New creates and returns a new instance of application with the specified config.
func New(config ...Config) *App {
	app := Default()

	if len(config) > 0 {
		app.setupByConfig(config[0])
	}

	return app
}

// Default creates and returns a new instance with the default config.
func Default() *App {
	app := new(App)

	app.lambdaOptions = make([]lambda.Option, 0)

	return app
}

// Start runs the Lambda function.
func (app *App) Start() {
	lambda.StartWithOptions(app.entry, app.lambdaOptions...)
}

// entry is the handler of the application for Lambda.
func (app *App) entry(
	ctx context.Context,
	req *events.APIGatewayProxyRequest,
) (*events.APIGatewayProxyResponse, error) {
	// TODO

	return nil, nil
}
