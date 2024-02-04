package handler

import (
	"net/http"

	"github.com/ViktorG13/gopportunities/schemas"
	"github.com/gin-gonic/gin"
)

func CreateOpeningHandler(ctx *gin.Context) {
	request := CreateOpeningRequest{}

	ctx.BindJSON(&request)

	logger.Infof("Request Received: %+v", request)

	if err := request.Validate(); err != nil {
		logger.Errorf("Validation Error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	opening := schemas.Opening{
		Role:     request.Role,
		Company:  request.Company,
		Location: request.Location,
		Remote:   *request.Remote,
		Link:     request.Link,
		Salary:   request.Salary,
	}

	if err := db.Create(&opening).Error; err != nil {
		logger.Errorf("Error Creating Opening: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "Error Creating Opening on Database.")
		return
	}

	sendSuccess(ctx, "create-opening", opening)
}
