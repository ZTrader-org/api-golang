package main

import (
	"github.com/ZTrader-org/api-golang/internal/platform"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"time"
)

var ztrader platform.Platformer

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err.Error())
	}

	transport := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
	}

	client := &http.Client{Transport: transport}

	baseUrl := os.Getenv("BASE_URL")
	apiKey := os.Getenv("API_KEY")
	apiSecret := os.Getenv("API_SECRET")
	port := os.Getenv("PORT")

	ztrader = platform.NewZtrader(client, baseUrl, apiKey, apiSecret)

	r := gin.Default()

	r.POST("/create-account", createAccount)

	if err = r.Run(port); err != nil {
		return
	}
}

func createAccount(c *gin.Context) {

	var jsonReq platform.CreateAccountRequest
	err := c.BindJSON(&jsonReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	res, err := ztrader.CreateAccount(c, jsonReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, res)
}
