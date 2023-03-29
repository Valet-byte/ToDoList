package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"todoApp/internal/service"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) InitHandler() *gin.Engine {

	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-in", h.signIn)
		auth.POST("/sign-up", h.signUp)
	}

	api := router.Group("/api", h.userIdentity)
	{
		lists := api.Group("/lists")
		{
			lists.POST("/", h.createList)
			lists.GET("/", h.getAllLists)
			lists.GET("/:id", h.getListById)
			lists.PUT("/:id", h.updateList)
			lists.DELETE("/:id", h.deleteList)
		}

		items := lists.Group(":id/items")
		{
			items.POST("/", h.createItem)
			items.GET("/", h.getAllItems)
			items.GET("/:item-id", h.getItemById)
			items.PUT("/:item-id", h.updateItem)
			items.DELETE("/:item-id", h.deleteItem)
		}
	}

	return router
}

func getUserId(con *gin.Context) (int64, error) {
	id, ok := con.Get(userId)
	if !ok {
		newErrorResponse(con, http.StatusUnauthorized, "User id not found!")
		return -1, errors.New("user id not found")
	}

	intId, ok := id.(int64)
	if !ok {
		newErrorResponse(con, http.StatusUnauthorized, "Incorrect user id!")
		return -1, errors.New("incorrect user id")
	}

	return intId, nil
}
