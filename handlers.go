package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sensand/quick-rest-api/firebase"
)

func getEntries(c *gin.Context) {
	entries, err := firebase.GetEntries(10)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    entries,
	})
}

func getEntry(c *gin.Context) {
	entryId := c.Param("id")

	entry, err := firebase.GetEntryById(entryId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    entry,
	})
}

func deleteEntry(c *gin.Context) {
	entryId := c.Param("id")

	c.JSON(http.StatusNotImplemented, gin.H{
		"success": false,
		"entryId": entryId,
	})
}

func createEntry(c *gin.Context) {
	var entry firebase.Entry

	if err := c.ShouldBindJSON(&entry); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	entryId, err := entry.CreateNew()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"id":      entryId,
	})
}

func updateEntry(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{
		"success": false,
	})
}
