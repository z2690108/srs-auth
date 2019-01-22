package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	// Success resp for Callback
	Success = iota
	// Failed resp for Callback
	Failed
)

// Publish publish
func Publish(c *gin.Context) {
	req := make(map[string]interface{})
	err := c.Bind(&req)
	if err != nil {
		fmt.Println("Publish Failed")
		c.AbortWithError(http.StatusOK, err)
		return
	}

	params, ok := req["param"].(string)
	if !ok || !authCheck(params) {
		fmt.Println("Auth check Failed")
		c.AbortWithStatus(http.StatusOK)
		return
	}

	fmt.Printf("Req: %+v\n", req)
	c.JSON(http.StatusOK, Success)
	return
}

// UnPublish unpublish
func UnPublish(c *gin.Context) {
	req := make(map[string]interface{})
	err := c.Bind(&req)
	if err != nil {
		fmt.Println("UnPublish Failed")
		c.AbortWithError(http.StatusOK, err)
		return
	}
	fmt.Printf("Req: %+v\n", req)
	c.JSON(http.StatusOK, Success)
	return
}

// Connect connect
func Connect(c *gin.Context) {
	req := make(map[string]interface{})
	err := c.Bind(&req)
	if err != nil {
		fmt.Println("Connect Failed")
		c.AbortWithError(http.StatusOK, err)
		return
	}
	fmt.Printf("Req: %+v\n", req)
	c.JSON(http.StatusOK, Success)
	return
}

// Close close
func Close(c *gin.Context) {
	req := make(map[string]interface{})
	err := c.Bind(&req)
	if err != nil {
		fmt.Println("Close Failed")
		c.AbortWithError(http.StatusOK, err)
		return
	}
	fmt.Printf("Req: %+v\n", req)
	c.JSON(http.StatusOK, Success)
	return
}

func authCheck(paramsStr string) bool {
	params := getParams(paramsStr)
	fmt.Println("Params:", params)

	if params["spUserID"] != "admin" {
		return false
	}
	return true
}

func getParams(paramsStr string) map[string]string {
	if paramsStr == "" {
		return nil
	}

	paramsStr = strings.TrimPrefix(paramsStr, "?")

	paramsSlice := strings.Split(paramsStr, "&")
	params := make(map[string]string, len(paramsSlice))

	kv := make([]string, 2, 2)
	for _, value := range paramsSlice {
		kv = strings.Split(value, "=")
		params[kv[0]] = kv[1]
	}

	return params
}
