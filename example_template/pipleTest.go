package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"text/template"
	"time"
)

func double(x int) int {
	return x + x
}
func square(x int) float64 {
	return math.Pow(float64(x), 2)
}
func sqRoot(x float64) float64 {
	return math.Sqrt(x)
}

var tpl *template.Template
var fm = template.FuncMap{
	"fdbl":  double,
	"fsq":   square,
	"fsqrt": sqRoot,
}

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("index2.html"))
}

func main() {
	t := time.Now()
	fmt.Println(t)
	tf := t.Format(time.Kitchen)
	fmt.Println(tf)

	err := tpl.ExecuteTemplate(os.Stdout, "index2.html", 3)
	if err != nil {
		log.Fatalln(err)
	}
}
