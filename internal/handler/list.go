package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
	"todoApp/internal/model"
)

func (h *Handler) createList(context *gin.Context) {
	id, err := getUserId(context)
	if err != nil {
		return
	}

	var input = model.ToDoList{}
	if err := context.BindJSON(&input); err != nil {
		newErrorResponse(context, http.StatusBadRequest, err.Error())
		return
	}
	listId, err := h.service.TodoListService.CreateList(id, input)
	if err != nil {
		newErrorResponse(context, http.StatusInternalServerError, err.Error())
		return
	}

	context.JSON(http.StatusOK, map[string]int64{"listId": listId})
}

func (h *Handler) updateList(context *gin.Context) {

}

func (h *Handler) getListById(context *gin.Context) {
	userId, err := getUserId(context)
	if err != nil {
		return
	}

	listId, err := strconv.ParseInt(context.Param("id"), 0, 64)
	if err != nil {
		logrus.Error("1")
		newErrorResponse(context, http.StatusInternalServerError, err.Error())
		return
	}

	list, err := h.service.TodoListService.GetById(userId, listId)
	if err != nil {
		logrus.Error("2")
		newErrorResponse(context, http.StatusInternalServerError, err.Error())
		return
	}

	context.JSON(http.StatusOK, list)
}

type getAllListResponse struct {
	Lists []model.ToDoList `json:"data"`
}

func (h *Handler) getAllLists(context *gin.Context) {

	userId, err := getUserId(context)
	if err != nil {
		return
	}

	lists, err := h.service.TodoListService.GetAll(userId)
	if err != nil {
		newErrorResponse(context, http.StatusInternalServerError, err.Error())
		return
	}

	context.JSON(http.StatusOK, getAllListResponse{Lists: lists})

}

func (h *Handler) deleteList(context *gin.Context) {

}
