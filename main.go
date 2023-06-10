package main

import (
	"fmt"
	"net/http"
	"strconv"
	"text/template"
	"time"

	"github.com/labstack/echo/v4"
)

type Project struct {
	Name string
	StarDate string
	EndDate string
	Duration string
	Detail string
	Playstore bool
	Android bool
	Java bool
	React bool	
}

var dataProject = []Project{
	{
		Name: "Project 1",
		StarDate: "15-05-2023",
		EndDate: "15-06-2023",
		Duration: "1 bulan",
		Detail: "Bootcamp sebulan gaes",
		Playstore: true,
		Android: true,
		Java: true,
		React: true,
	},
	
	{
		Name: "Project 2",
		StarDate: "15-05-2023",
		EndDate: "15-06-2023",
		Duration: "1 bulan",
		Detail: "Bootcamp sebulan gaes hehe",
		Playstore: true,
		Android: true,
		Java: true,
		React: true,
	},

	
}

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
	
	// routing post
	e.POST("/saveproject", saveProject)
	e.POST("/deleteProject/:id", deleteProject)

	e.Logger.Fatal(e.Start("localhost:5001"))
}

func hai(c echo.Context) error {
	return c.String(http.StatusOK, "haii dunia golang!!!, aku akan menaklukanmuuu....")
}

func home(c echo.Context) error {
	var tmpl, err = template.ParseFiles("views/index.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message":err.Error()})
	}
	
	
	Projects := map[string]interface{}{
		"Projects" : dataProject,
	}
	return tmpl.Execute(c.Response(), Projects)
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
	detail := c.FormValue("description")
	
	// ambil date input
	date1 := c.FormValue("input-start-date")
	date2 := c.FormValue("input-end-date")
	// parse date input dan formatting
	uDate1, _ := time.Parse("2006-01-02", date1)
	starDate := uDate1.Format("2 Jan 2006")

	uDate2, _ := time.Parse("2006-01-02", date2)
	endDate := uDate2.Format("2 Jan 2006")

	// perhitungan selisih
	var diffUse string
	timeDiff := uDate2.Sub(uDate1)

	if timeDiff.Hours()/24 < 30 {
		tampil := strconv.FormatFloat(timeDiff.Hours()/24, 'f', 0, 64)
		diffUse = "Duration : " +tampil+" hari"
	}else if timeDiff.Hours()/24/30 < 12 {
		tampil := strconv.FormatFloat(timeDiff.Hours()/24/30, 'f', 0, 64)
		diffUse = "Duration : " +tampil+ " Bulan"
	}else {

	}
	// checkbox
	var playstore bool
	if c.FormValue("playstore") == "checked"{
		playstore = true
	}
	
	var android bool
	if c.FormValue("android") == "checked"{
		android = true
	}
	
	var java bool
	if c.FormValue("java") == "checked"{
		java = true
	}
	
	var react bool
	if c.FormValue("react") == "checked"{
		react = true
	}

	
	
	var newProject = Project{
		Name: name,
		StarDate: starDate,
		EndDate: endDate,
		Duration: diffUse,
		Detail: detail,
		Playstore: playstore,
		Android: android,
		Java: java,
		React: react,

	}

	dataProject = append(dataProject, newProject)

	fmt.Println(dataProject)

	return c.Redirect(http.StatusMovedPermanently,"/")
}

func deleteProject (c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	fmt.Println("index: ", id)

	dataProject = append(dataProject[:id], dataProject[:id+1]...)

	return c.Redirect(http.StatusMovedPermanently, "/")
}

