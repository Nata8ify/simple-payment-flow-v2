package main

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	echoadapter "github.com/awslabs/aws-lambda-go-api-proxy/echo"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"os"
	"strconv"
)

var echoLambda *echoadapter.EchoLambda
var e *echo.Echo

func init() {
	e = echo.New()
	server := Server{}
	server.mount(e)
	echoLambda = echoadapter.New(e)
}

type Server struct{}
type GenerateAccountResponse struct {
	Address    string `json:"address"`
	PrivateKey string `json:"privateKey"`
	PublicKey  string `json:"publicKey"`
}

func (server *Server) mount(e *echo.Echo) {
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World! Ya")
	})
	e.GET("/generate", func(c echo.Context) error {

		/* Initial input parameter */
		change, err := strconv.Atoi(c.QueryParam("change"))
		if err != nil {
			panic("Unable to retrieving 'change' value")
		}
		index, err := strconv.Atoi(c.QueryParam("index"))
		if err != nil {
			panic("Unable to retrieving 'index' value")
		}

		/* Generate account */
		account := Generate(change, index)

		/* Generate keypair */
		pvk, pub := GetKeypair(account)

		return c.JSON(http.StatusOK, GenerateAccountResponse{
			Address:    account.Address.String(),
			PrivateKey: pvk,
			PublicKey:  pub,
		})

	})
}

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return echoLambda.ProxyWithContext(ctx, req)
}

func main() {
	_, isLambdaMode := os.LookupEnv("LAMBDA_TASK_ROOT")
	if isLambdaMode {
		lambda.Start(Handler)
	} else {
		log.Fatal(e.Start(":1234"))
	}
}
