package lamda

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// App is the application to run a HTTP service with AWS Lambda and AWS APIGateway.
type App struct {
}

// New creates and returns a new instance of application.
func New() *App {
	app := new(App)

	return app
}

// Start runs the Lambda function.
func (app *App) Start() {
	lambda.StartWithOptions(app.entry)
}

// entry is the handler of the application for Lambda.
func (app *App) entry(
	ctx context.Context,
	req *events.APIGatewayProxyRequest,
) (*events.APIGatewayProxyResponse, error) {
	// TODO

	return nil, nil
}
