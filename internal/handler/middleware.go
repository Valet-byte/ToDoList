package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const (
	authorizationHeader = "Authorization"
	userId              = "userId"
)

func (h *Handler) userIdentity(con *gin.Context) {
	header := con.GetHeader(authorizationHeader)

	if header == "" {
		newErrorResponse(con, http.StatusUnauthorized, "empty auth header")
		return
	}

	headerParts := strings.Split(header, " ")

	if len(headerParts) != 2 {
		newErrorResponse(con, http.StatusUnauthorized, "invalid auth header")
		return
	}

	uId, err := h.service.ParseToken(headerParts[1])

	if err != nil {
		newErrorResponse(con, http.StatusUnauthorized, err.Error())
		return
	}

	con.Set(userId, uId)
}
