package main

import (
	"github.com/gin-gonic/gin"
)

func tsProxyHandler(c *gin.Context) {
	remoteURL := c.Query("url")
	resp, err := getHTTPClient().Get(remoteURL)
	if err != nil {
		c.AbortWithError(500, err)
		return
	}
	defer resp.Body.Close()
	c.DataFromReader(200, resp.ContentLength, resp.Header.Get("Content-Type"), resp.Body, nil)
}
