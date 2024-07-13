package main

import (
	"fmt"
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

	passwordLength := len(request.Password)
	minLowerCase := 1
	minUpperCase := 1
	minDigit := 1
	
	minLength := 6
	maxLength := 20

	otherLength :=  minLength - minLowerCase - minUpperCase - minDigit;

	if passwordLength >= maxLength {
		var response model.ErrorResponse
		response.TimeStamp = time.Now().Format(time.RFC3339)
		response.Status = http.StatusBadRequest
		response.Error = fmt.Sprintf("Password length should be greater than %v but less than %v", minLength, maxLength)
		response.Path = ctx.Request.URL.Path
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	if match := MatchRepeatCharacter(request.Password, 3); match {
		var response model.ErrorResponse
		response.TimeStamp = time.Now().Format(time.RFC3339)
		response.Status = http.StatusBadRequest
		response.Error = "Password shouldn't contain 3 repeating characters in a row"
		response.Path = ctx.Request.URL.Path
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	if match, _ := regexp.MatchString("[a-z]", request.Password); match {
		minLowerCase--
		passwordLength--
	}

	if match, _ := regexp.MatchString("[A-Z]", request.Password); match {
		minUpperCase--
		passwordLength--
	}

	if match, _ := regexp.MatchString("[\\d]", request.Password); match {
		minDigit--
		passwordLength--
	}

	var response model.StorngPasswordResponse

	if passwordLength < otherLength {
		response.Steps = otherLength - passwordLength + minLowerCase + minUpperCase + minDigit;
		ctx.JSON(200, response)

	} else {
		response.Steps = minLowerCase + minUpperCase + minDigit;
		ctx.JSON(200, response)
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
				} else {
					break
				}
			}
			if counter == repeat {
				return true
			}
		}
	}

	return false
}
