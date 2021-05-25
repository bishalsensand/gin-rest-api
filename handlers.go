package main

import (
	"net/http"

	"github.com/bishalsensand/gin-rest-api/firebase"
	"github.com/gin-gonic/gin"
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

	if entry == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": "Could not find a matching entry",
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
	success, err := firebase.DeleteEntryById(entryId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": success,
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"success": success,
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
