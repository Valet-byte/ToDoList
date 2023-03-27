package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"todoApp/internal/model"
)

func (h *Handler) signIn(con *gin.Context) {
}

func (h *Handler) signUp(con *gin.Context) {

	var input = &model.User{}

	if err := con.BindJSON(input); err != nil {
		NewErrorResponse(con, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.service.AuthorizationService.CreateUser(*input)
	if err != nil {
		NewErrorResponse(con, http.StatusInternalServerError, err.Error())
	}

	con.JSON(http.StatusOK, map[string]interface{}{"id": id})
}
