package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/seew0/DoubtBuddy/api"
	"github.com/seew0/DoubtBuddy/util"
)

type Request struct {
	Query string `json:"query"`
}

func Getanswer(c *gin.Context) {
	c.Writer.Header().Set("Content-Type", "application/json")
	req := new(Request)
	c.BindJSON(&req)

	rawData := api.GetData(req.Query)
	data := util.GetCode(rawData)

	if rawData == "error" {
		c.JSON(400, gin.H{"error": "some error occurred try again"})
		return
	}
	
	if data == "Code extraction failed."{
		c.JSON(400, gin.H{"error":"some error occurred try again"})
		return
	}

	c.JSON(200, gin.H{"status": "succesful", "answer": data})
}
