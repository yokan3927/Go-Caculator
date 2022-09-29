package main

import (
	math "calculator-app/calculate"
	"fmt"
	// "fmt"
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
)

type Data struct {
	Ans float64
}

var temp float64 = 0
var operator string = ""
var ans float64

func index(c *gin.Context) {
	value := temp
	data := Data{value}
	c.HTML(http.StatusOK, "index.html", data)
}

func caculate(temp float64, a string, operator string) float64 {
	n2, _ := strconv.ParseFloat(a, 64)
	switch operator {
	case "+":
		ans = math.Add(temp, n2)
	case "-":
		ans = math.Minus(temp, n2)
	case "*":
		ans = math.Times(temp, n2)
	case "/":
		ans = math.Divide(temp, n2)
	case "%":
		ans = math.Mod(temp, n2)
	}
	return ans
}
func pow(temp float64) float64 {
	ans = math.Pow(temp)
	return ans
}

func add(c *gin.Context) {
	a, _ := c.GetPostForm("value")
	if operator != "" {
		ans = caculate(temp, a, operator)
		temp = 0
		operator = ""

	} else {
		operator_temp, _ := c.GetPostForm("submit")
		operator = operator_temp
		fmt.Println("operator:",operator)
		switch operator {
		case "^2":
			i, _ := strconv.ParseFloat(a, 64)
			ans = pow(i)
			temp = 0
			operator = ""
		default:
			i, _ := strconv.ParseFloat(a, 64)
			temp = i
		}
	}
	data := Data{ans}
	ans = 0

	c.HTML(http.StatusOK, "index.html", data)
}

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("template/html/*")
	r.Static("/assets", "./template/assets")
	r.GET("/", index)
	r.POST("/", add)
	r.Run(":8080")
}
