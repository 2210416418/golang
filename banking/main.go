package main

import (
	"banking/api"
	_ "banking/api"
	_ "encoding/json"
	_ "fmt"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	router.POST("/account", api.AddAccount)
	router.GET("/accounts", api.GetAccounts)
	router.DELETE("/accountdelete/accountNumber", api.DeleteAccount)
	router.Run()

	// var acc = api.Account_create{
	// 	AccountNumbr:   500,
	// 	AccountName:    "sai",
	// 	AccountBalance: 56.00,
	// 	Address: api.Address{
	// 		Apt:     567,
	// 		Street:  "yhgvggg",
	// 		City:    "frisco",
	// 		ZipCode: 64468,
	// 	},
	// }
	// x, _ := json.Marshal(acc)
	// fmt.Println(x)

}
