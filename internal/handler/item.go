package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"todoApp/internal/model"
)

func (h *Handler) createItem(context *gin.Context) {
	userId, err := getUserId(context)
	if err != nil {
		return
	}

	listId, err := strconv.ParseInt(context.Param("id"), 0, 64)
	if err != nil {
		newErrorResponse(context, http.StatusInternalServerError, err.Error())
		return
	}

	var input = model.ToDoItem{}
	if err := context.BindJSON(&input); err != nil {
		newErrorResponse(context, http.StatusBadRequest, err.Error())
		return
	}

	listId, err = h.service.ToDoItemService.CreateItemList(userId, listId, input)
	if err != nil {
		newErrorResponse(context, http.StatusInternalServerError, err.Error())
		return
	}
	context.JSON(http.StatusOK, map[string]int64{"id": listId})
}

func (h *Handler) updateItem(context *gin.Context) {

}

func (h *Handler) getItemById(context *gin.Context) {

}

func (h *Handler) getAllItems(context *gin.Context) {

}

func (h *Handler) deleteItem(context *gin.Context) {

}
