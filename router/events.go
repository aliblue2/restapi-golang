package router

import (
	"fmt"
	"net/http"
	"strconv"

	"azno-space.com/azno/models"
	"github.com/gin-gonic/gin"
)

func GetAllEventsHandler(context *gin.Context) {

	events, err := models.GetAllEvents()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "invalid token"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"events": events})
}

func GetEventByIdHandler(context *gin.Context) {
	idString := context.Param("eventId")

	id, err := strconv.ParseInt(idString, 10, 64)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "an error ocurred during to converting event id"})
		return
	}

	event, err := models.GetEventById(id)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "cant find event with this id"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"event": event})

}

func AddNewEventHandler(context *gin.Context) {

	var event models.Event

	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusConflict, gin.H{"message": "cant save this object as event object"})
		return
	}
	userId := context.GetInt64("userId")

	event.UserId = userId

	err = event.SaveEvent()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "cant save this event into database"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "successfully created"})

}

func DeleteEventHandler(context *gin.Context) {
	idString := context.Param("eventId")
	id, err := strconv.ParseInt(idString, 10, 64)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "cant convet id string to int"})
		return
	}

	event, err := models.GetEventById(id)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "cant find event with id"})
		return
	}

	userId := context.GetInt64("userId")

	if event.Id != userId {
		context.JSON(http.StatusForbidden, gin.H{"message": "only creator of event can delete it"})
		return
	}

	err = models.DeleteEventById(id)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}
	message := fmt.Sprintf("event with id :%v successfully deleted", id)
	context.JSON(http.StatusOK, gin.H{"message": message})

}
func EditEventHandler(context *gin.Context) {
	idString := context.Param("eventId")

	id, err := strconv.ParseInt(idString, 10, 64)

	if err != nil {
		context.JSON(http.StatusBadGateway, gin.H{"message": "an error occured during to converting to int"})
		return
	}

	tempEvent, err := models.GetEventById(id)

	if err != nil {
		context.JSON(http.StatusConflict, gin.H{"message": "cant find event with this id"})
		return
	}

	userId := context.GetInt64("userId")

	if tempEvent.Id != userId {
		context.JSON(http.StatusForbidden, gin.H{"message": "only creator of event can update it"})
		return
	}

	var event models.Event

	context.ShouldBindJSON(&event)
	event.UserId = userId

	err = models.UpdateEventById(id, event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Cant update event!."})
		return
	}

	message := fmt.Sprintf("event with id :%v , was successfully updated", id)
	context.JSON(http.StatusOK, gin.H{"message": message})

}
