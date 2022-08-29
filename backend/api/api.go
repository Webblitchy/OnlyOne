package api

import (
	//import gin

	"backend/db"
	cookie "backend/utils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func Setup() *gin.Engine {
	router = gin.Default()
	user := router.Group("/user")
	{
		user.POST("/login", login)
		user.POST("/register", register)
		user.GET("/:username", showUser)             //.Use(isAuthenticated())
		user.GET("/:username/sock", listSocksOfUser) //.Use(isAuthenticated())
	}

	sock := router.Group("/sock")
	{
		sock.POST("/", addSock)                        //.Use(isAuthenticated())
		sock.GET("/:sockId/match", listMatchesOfSock)  //.Use(isAuthenticated())
		sock.PATCH("/:sockId/", patchAcceptListOfSock) //.Use(isAuthenticated())
		sock.GET("/:sockId", getSockInfo)
	}

	return router
}

func getSockInfo(c *gin.Context) {
	c.Next()
}

func patchAcceptListOfSock(c *gin.Context) {
	c.Next()
}
func showUser(c *gin.Context) {
	c.Next()
}

func listSocksOfUser(c *gin.Context) {
	c.Next()
}

func addSock(c *gin.Context) {
	c.Next()

}

func listMatchesOfSock(c *gin.Context) {
	c.Next()
}

func isAuthenticated() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

// get the next data for the user
func nextOne(c *gin.Context) {

	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func testLogin(c *gin.Context) {
	usr := c.Query("username")
	pwd := c.Query("password")
	db.VerifyLogin(usr, pwd)
	c.Next()

}

// create the login function
func login(c *gin.Context) {
	//retrieve the username and the password from the post data
	// username := c.PostForm("username")
	// pwd := c.PostForm("password")

	username := c.Query("username")
	pwd := c.Query("password")
	//check if the username and password are correct
	db.VerifyLogin(username, pwd)
	if username == "m" && pwd == "myPwd" {
		//if they are correct, return a success message
		//TODO - add a token to the response
		c.JSON(200, gin.H{
			"message": "login successful",
		})
	} else {
		//if they are not correct, return an error message
		c.JSON(401, gin.H{
			"message": "login failed",
		})
	}
}

// create the register function
func register(c *gin.Context) {
	c.Next()
}
