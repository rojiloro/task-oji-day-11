package main

import (
	"net/http"
	"strconv"
	"text/template"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	// static file from 'public' directory
	e.Static("/public", "public")

	e.GET("/hello", hai)
	e.GET("/", home)
	e.GET("/myproject", addProject)
	e.GET("/testimonial", testimonial)
	e.GET("/contact", contact)
	e.GET("/project-detail", projectDetail)
	e.POST("/saveProject", saveProject)

	e.Logger.Fatal(e.Start("localhost:5000"))
}

func hai(c echo.Context) error {
	return c.String(http.StatusOK, "haii dunia golang!!!, aku akan menaklukanmuuu....")
}

func home(c echo.Context) error {
	var tmpl, err = template.ParseFiles("views/index.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message":err.Error()})
	}

	return tmpl.Execute(c.Response(), nil)
}

func addProject(c echo.Context) error {
	var tmpl, err = template.ParseFiles("views/myproject.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message":err.Error()})
	}

	return tmpl.Execute(c.Response(), nil)
}

func testimonial(c echo.Context) error {
	var tmpl, err = template.ParseFiles("views/testimonial.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message":err.Error()})
	}

	return tmpl.Execute(c.Response(), nil)
}

func contact(c echo.Context) error {
	var tmpl, err = template.ParseFiles("views/contact.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message":err.Error()})
	}

	return tmpl.Execute(c.Response(), nil)
}

func projectDetail(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	data := map[string]interface{}{
		"Id": id,
		"judul": "Dumbways Mobile App - 2021",
		"durasi": "durasi : 3 bulan",
		"detail": "Lorem ipsum dolor sit amet consectetur adipisicing elit. Iste praesentium explicabo quae architecto aperiam at quis possimus voluptatibus ducimus minima.",
		"waktu": "17, week 3, jan, 2023",
	}

	var tmpl, err = template.ParseFiles("views/project-detail.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message":err.Error()})

	}
	return tmpl.Execute(c.Response(), data)
}

func saveProject(c echo.Context) error {
	name := c.FormValue("input-project-name")
	startDate := c.FormValue("input-start-date")
	endDate := c.FormValue("input-end-date")
	Playstore :=  c.FormValue("playstore")
    Android := c.FormValue("android")
    Java := c.FormValue("java")
    React := c.FormValue("react")
	descrition := c.FormValue("description")

	println("Project name : "+name)
	println("Star date : "+startDate)
	println("End date : "+endDate)
	println("Playstore :"+Playstore)
	println("Android : "+Android)
	println("Java : "+Java)
	println("React : "+React)
	println("description : "+descrition)
	return c.Redirect(http.StatusMovedPermanently,"/")
}