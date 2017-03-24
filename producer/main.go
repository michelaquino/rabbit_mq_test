package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/michelaquino/rabbit_mq_test/producer/context"
	"github.com/michelaquino/rabbit_mq_test/producer/model"
)

func main() {
	apiContext := context.GetAPIContext()
	logger := apiContext.GetLogger()

	e := echo.New()
	e.GET("/sendMessage", func(c echo.Context) error {
		model.SendMessageToQueue()

		return c.String(http.StatusOK, "Message sended")
	})

	logger.Info("Started at port 8888!")
	e.Logger.Fatal(e.Start(":8888"))
}
