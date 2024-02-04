package handler

import (
	"net/http"

	"github.com/ViktorG13/gopportunities/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Summary List all openings
// @Description List all job openings created
// @Tags Openings
// @Accept json
// @Produce json
// @Failure 500 {object} ErrorResponse
// @Router /openings [get]
func ListOpeningHandler(ctx *gin.Context) {
	openings := []schemas.Opening{}

	if err := db.Find(&openings).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "error: listening openings.")
		return
	}

	sendSuccess(ctx, "list-openings", openings)
}
