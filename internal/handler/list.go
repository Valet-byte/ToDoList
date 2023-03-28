package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
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

}

func (h *Handler) getAllLists(context *gin.Context) {

}

func (h *Handler) deleteList(context *gin.Context) {

}
