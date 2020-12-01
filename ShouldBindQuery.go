package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

type Student struct {
	Name    string `form:"name"`
	CLasses string `form:"classes"`
}

func main01() {
	engine := gin.Default()
	//http://localhost:8080/hello?name=mike&classes=class01
	engine.GET("/hello", func(context *gin.Context) {
		fmt.Println(context.FullPath())
		var student Student
		err := context.ShouldBindQuery(&student)
		if err != nil {
			log.Fatal(err.Error())
			return
		}

		fmt.Println(student.Name)
		fmt.Println(student.CLasses)

		context.Writer.Write([]byte(student.Name + " hello"))
	})
	engine.Run()
}
