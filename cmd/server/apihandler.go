package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func handleNewJob(datastore Database) gin.HandlerFunc {

	return func(c *gin.Context) {
		request := JobSchedule{}
		err := c.ShouldBindJSON(&request)
		if err != nil {
			c.JSON(http.StatusBadRequest, "bad request")
			return
		}

		if request.ID != 0 {
			c.JSON(http.StatusBadRequest, "id has to be zero")
			return
		}
	}


}
