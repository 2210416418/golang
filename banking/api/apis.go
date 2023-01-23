package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var Accounts []Account_create

func AddAccount(c *gin.Context) {
	var req Account_create
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	Accounts = append(Accounts, req)
	c.JSON(http.StatusOK, req)
	fmt.Println(Accounts)

}

func GetAccounts(c *gin.Context) {
	fmt.Println(Accounts)

	c.JSON(http.StatusOK, Accounts)

}

func DeleteAccount(c *gin.Context) {
	fmt.Println(Accounts)
	del := c.Param("AccountNumbr")
	temp, _ := strconv.Atoi(del)
	fmt.Println("--------------*********____________------------", del)
	for i, a := range Accounts {
		if a.AccountNumbr == temp {
			Accounts = append(Accounts[:i], Accounts[i+1:]...)
		}

	}
	fmt.Println(Accounts)
}
