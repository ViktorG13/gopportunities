package handler

import (
	"fmt"
	"net/http"

	"github.com/ViktorG13/gopportunities/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Summary Show a opening
// @Description Show a job opening
// @Tags Openings
// @Accept json
// @Produce json
// @Param id query string true "Opening identification"
// @Success 200 {object} ShowOpeningResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /opening [get]
func ShowOpeningHandler(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id",
			"queryParameter").Error())
		return
	}

	opening := schemas.Opening{}

	// Find Opening
	if err := db.First(&opening, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, fmt.Sprintf("Opening With id: %s Not Found.", id))
		return
	}

	sendSuccess(ctx, "show-opening", opening)
}
