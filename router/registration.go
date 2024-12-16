package router

import (
	"net/http"
	"strconv"

	"azno-space.com/azno/models"
	"github.com/gin-gonic/gin"
)

func RegisterEventHandler(context *gin.Context) {

	userId := context.GetInt64("userId")

	eventId, err := strconv.ParseInt(context.Param("eventId"), 10, 64)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "an error occured"})
		return
	}

	event, err := models.GetEventById(eventId)

	if err != nil {
		context.JSON(http.StatusForbidden, gin.H{"message": "cant find event with this id"})
		return
	}

	_, err = event.Register(userId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "cant reserve registration for tis event"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "successfully reserved registration"})

}

func CancelEventHandler(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("eventId"), 10, 64)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "an error occured"})
		return
	}

	_, err = models.GetEventById(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "cant find event with this id"})
		return
	}

	var event models.Event
	event.Id = eventId

	rows, err := event.CancelRegisteration(userId)

	if err != nil {
		context.JSON(http.StatusForbidden, gin.H{"message": "an error ocurred during to canceling reservation"})
		return
	}

	if rows == 0 {
		context.JSON(http.StatusConflict, gin.H{"message": "cant cancel reserv with this creadentials"})
	} else {
		context.JSON(http.StatusOK, gin.H{"message": "reservation successfully canceled"})
	}

}
