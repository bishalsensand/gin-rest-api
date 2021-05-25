package main

import (
	"net/http"

	"github.com/bishalsensand/gin-rest-api/firebase"
	"github.com/gin-gonic/gin"
)

func getEntries(c *gin.Context) {
	entries, err := firebase.GetEntries(10)

	if err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"success": false,
			"error":   err.Error(),
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
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"success": false,
			"error":   err.Error(),
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
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"success": success,
			"error":   err.Error(),
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

	entry.Id = nil // disallow setting Id through request body

	entryId, err := entry.CreateNew()

	if err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"id":      entryId,
	})
}

func putEntry(c *gin.Context) {
	entryId := c.Param("id")
	var entry firebase.Entry

	if err := c.ShouldBindJSON(&entry); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	matchedEntry, err := firebase.GetEntryById(entryId)

	if err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	if matchedEntry == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Could not find an entry with given id",
		})
		return
	}

	entry.Id = nil // disallow setting Id from request body
	success, err := entry.UpdateById(entryId)

	if err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"success": success,
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": success,
		"id":      entryId,
	})
}
