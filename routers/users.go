package routers

import (
	"fmt"
	"github.com/Ivanhahanov/LDAPManager/ldap"
	"github.com/gin-gonic/gin"
	"net/http"
)

type User struct {
}

func AddUserHandler(c *gin.Context) {
	//var user *User
	//err := c.BindJSON(&user); if err != nil{
	//	log.Println("add user:", err)
	//}
	//log.Println(user)
	ldap.AddUser()
}

func ChangeUserHandler(c *gin.Context) {
	c.String(http.StatusOK, "")
}

func GetUserHandler(c *gin.Context) {
	username := c.Param("username")
	user := ldap.SearchUser(username)
	c.String(http.StatusOK, fmt.Sprintf("%v", user[0].DN))
}

func UsersListHandler(c *gin.Context) {
	c.String(http.StatusOK, "")
}

func RemoveUserHandler(c *gin.Context) {
	c.String(http.StatusOK, "The available groups are [...]")
}
