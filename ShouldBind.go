package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

type Register struct {
	Username string `form:"username"`
	Password string `form:"password"`
	Mail     string `form:"mail"`
}

func main02() {
	engine := gin.Default()
	engine.POST("/register", func(context *gin.Context) {
		fmt.Println(context.FullPath())

		var register Register
		err := context.ShouldBind(&register)
		if err != nil {
			log.Fatal(err.Error())
			return
		}
		fmt.Println(register.Username)
		fmt.Println(register.Mail)

		context.Writer.Write([]byte(register.Username + "  register"))
	})

	engine.Run()
}
