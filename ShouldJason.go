package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

type Person struct {
	Name string `form:"name"`
	Sex  string `form:"sex"`
	Age  string `form:"age"`
}

func main() {
	engine := gin.Default()
	engine.POST("/addstudent", func(context *gin.Context) {
		fmt.Println(context.FullPath())

		var person Person
		err := context.BindJSON(&person)
		if err != nil {
			log.Fatal(err.Error())
			return
		}

		fmt.Println(person.Name)
		fmt.Println(person.Age)

		context.Writer.Write([]byte("增加用户" + person.Name))
	})
	engine.Run()
}
