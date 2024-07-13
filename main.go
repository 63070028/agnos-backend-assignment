package main

import (
	"net/http"
	"regexp"
	"time"

	"github.com/63070028/agnos-backend-assignment/model"
	"github.com/gin-gonic/gin"
)

func main() {

	route := gin.Default()

	route.GET("/api", func(ctx *gin.Context) {
		ctx.JSON(200, "Hello Api")
	})

	route.POST("/api/strong_password_steps", strongPasswordStep)

	route.Run(":8000")
}

func strongPasswordStep(ctx *gin.Context) {
	var request model.StorngPasswordRequest

	if err := ctx.ShouldBindJSON(&request); err != nil {
		var response model.ErrorResponse
		response.TimeStamp = time.Now().Format(time.RFC3339)
		response.Status = http.StatusBadRequest
		response.Error = http.StatusText(400)
		response.Path = ctx.Request.URL.Path
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	if request == (model.StorngPasswordRequest{}) {
		var response model.ErrorResponse
		response.TimeStamp = time.Now().Format(time.RFC3339)
		response.Status = http.StatusBadRequest
		response.Error = "Request is empty"
		response.Path = ctx.Request.URL.Path
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	length := len(request.Password)
	var response model.StorngPasswordResponse
	if length < 6 {
		response.Steps = 6 - length
		ctx.JSON(200, response)
	} else {
		response.Steps = 0
		ctx.JSON(200, response)
	}

	if match, _ := regexp.MatchString("[a-z]", request.Password); !match {
		var response model.ErrorResponse
		response.TimeStamp = time.Now().Format(time.RFC3339)
		response.Status = http.StatusBadRequest
		response.Error = "Password should least 1 lowercase letter"
		response.Path = ctx.Request.URL.Path
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	if match, _ := regexp.MatchString("[A-Z]", request.Password); !match {
		var response model.ErrorResponse
		response.TimeStamp = time.Now().Format(time.RFC3339)
		response.Status = http.StatusBadRequest
		response.Error = "Password should least 1 uppercase letter"
		response.Path = ctx.Request.URL.Path
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	if match, _ := regexp.MatchString(".*[A-Z].*", request.Password); !match {
		var response model.ErrorResponse
		response.TimeStamp = time.Now().Format(time.RFC3339)
		response.Status = http.StatusBadRequest
		response.Error = "Password should least 1 uppercase letter"
		response.Path = ctx.Request.URL.Path
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	if match := MatchRepeatCharacter(request.Password, 3); !match {
		var response model.ErrorResponse
		response.TimeStamp = time.Now().Format(time.RFC3339)
		response.Status = http.StatusBadRequest
		response.Error = "Password shouldn't contain 3 repeating characters in a row"
		response.Path = ctx.Request.URL.Path
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
}

func MatchRepeatCharacter(text string, repeat int) bool {
	length := len(text)
	if length >= repeat {
		for i := 0; i < length-repeat+1; i++ {
			curr := text[i]
			counter := 1
			for shift := 1; shift < repeat; shift++ {
				if curr == text[i+shift] {
					counter++
				}
			}
			if counter == repeat {
				return true
			}
		}
	}

	return false
}
