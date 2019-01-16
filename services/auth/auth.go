package authservice

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type credential struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

/*
Login ...
*/
func Login(c *gin.Context) {

	var cred credential
	c.BindJSON(&cred)
	if cred.Username == "test" && cred.Password == "test1234" {
		c.JSON(http.StatusOK, gin.H{
			"message": "Logged In successfully",
			"token":   "poiuytrewqasdfghjklmnbvcxze2bnjde76yh2nj3",
		})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Bad credentials",
		})
	}
	return

}

/*
Register ...
*/
func Register(c *gin.Context) {
	c.JSON(http.StatusOK, "register")
}
