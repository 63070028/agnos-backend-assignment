package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/63070028/agnos-backend-assignment/model"
	"github.com/63070028/agnos-backend-assignment/service"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {

	route := gin.Default()
	godotenv.Load(".env")

	config := model.ConfigStrongPassword{
		MinLowerCase: 1,
		MinUpperCase: 1,
		MinDigit:     1,
		MinLength:    6,
		MaxLength:    19,
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s",
		os.Getenv("DB_HOST"), os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_PORT"))
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&model.StorngPasswordLog{})

	route.POST("/api/strong_password_steps", func(ctx *gin.Context) {
		var request model.StorngPasswordRequest

		log := &model.StorngPasswordLog{}
		log.Ip = ctx.ClientIP()

		if err := ctx.ShouldBindJSON(&request); err != nil || len(request.Password) == 0 {
			var response model.ErrorResponse
			response.TimeStamp = time.Now().Format(time.RFC3339)
			response.Status = http.StatusBadRequest
			response.Error = http.StatusText(400)
			response.Path = ctx.Request.URL.Path
			log.Request = JsonMarshal(request)
			log.Response = JsonMarshal(response)
			db.Create(log)

			ctx.JSON(http.StatusBadRequest, response)
			return
		}

		log.Request = JsonMarshal(request)

		var response model.StorngPasswordResponse
		response.Steps = service.MiminimumActions(request.Password, config)
		log.Response = JsonMarshal(response)

		db.Create(log)
		ctx.JSON(200, response)
	})

	route.Run(":" + os.Getenv("PORT"))

}

func JsonMarshal(obj any) string {
	b, err := json.Marshal(obj)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return ""
	}
	return string(b)
}
