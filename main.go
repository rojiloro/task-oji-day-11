package main

import (
	"context"
	"fmt"
	"net/http"
	"oji/connection"
	"strconv"
	"text/template"
	"time"

	"github.com/labstack/echo/v4"
)

type Project struct {
	Id int
	Name string
	StarDate string
	EndDate string
	Duration string
	Detail string
	Playstore bool
	Android bool
	Java bool
	React bool
	StartDateTime time.Time
	EndDateTime time.Time	
}

var dataProject = []Project{
	{
		Id: 0,
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
		Id: 1,
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
	connection.DatabaseConnect()

	e := echo.New()

	// static file from 'public' directory
	e.Static("/public", "public")

	e.GET("/hello", hai)
	e.GET("/", home)
	e.GET("/myproject", addProject)
	e.GET("/testimonial", testimonial)
	e.GET("/contact", contact)
	e.GET("/project-detail/:id", projectDetail)
	e.GET("/project-edit/:id", editProject)
	
	// routing post
	e.POST("/saveproject", saveProject)
	e.POST("/deleteProject/:id", deleteProject)
	e.POST("/updateProject/:id", updateProject)

	e.Logger.Fatal(e.Start("localhost:5001"))
}

func hai(c echo.Context) error {
	return c.String(http.StatusOK, "haii dunia golang!!!, aku akan menaklukanmuuu....")
}

func home(c echo.Context) error {
	data, _ := connection.Conn.Query(context.Background(), "SELECT id, name, star_date, end_date, duration, detail, playstore, android, java, react FROM tb_project")

	var result []Project
	for data.Next() {
		var each = Project{}

		err := data.Scan(&each.Id, &each.Name, &each.StartDateTime, &each.EndDateTime, &each.Duration, &each.Detail, &each.Playstore, &each.Android, &each.Java, &each.React)

		if err != nil {
			fmt.Println(err.Error())
			return c.JSON(http.StatusInternalServerError, map[string]string{"message":err.Error()})
		}

		each.StarDate = each.StartDateTime.Format("01-02-2006")
		each.EndDate = each.EndDateTime.Format("01-02-2006")

		result = append(result, each)
	}
	
	Projects := map[string]interface{}{
		"Projects" : result,
	}

	var tmpl, err = template.ParseFiles("views/index.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message":err.Error()})
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

	var projectDetails = Project{}

	err := connection.Conn.QueryRow(context.Background(), "SELECT * FROM tb_project WHERE id=$1", id).Scan(
		&projectDetails.Id, &projectDetails.Name, &projectDetails.StartDateTime, &projectDetails.EndDateTime, &projectDetails.Duration, &projectDetails.Detail, &projectDetails.Playstore, &projectDetails.Android, &projectDetails.Java, &projectDetails.React,
	)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string {"message":err.Error()})
	}

	data := map[string]interface{}{
		"Project": projectDetails,
		"StarDate" : projectDetails.StartDateTime.Format("01-02-2006"),
		"EndDate" : projectDetails.EndDateTime.Format("01-02-2006"),
	}

	var tmpl, errTemp = template.ParseFiles("views/project-detail.html")

	if errTemp != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": errTemp.Error()})
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

	_, err := connection.Conn.Exec(
		context.Background(),`INSERT INTO tb_project (name, star_date, end_date, duration, detail, playstore, android, java, react)
		Values ($1, $2, $3, $4, $5, $6, $7, $8, $9)`,
		name, starDate, endDate, diffUse, detail, playstore, android, java, react,
	)
	
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message":err.Error()})
	}

	// var id int
	// err = connection.Conn.QueryRow(context.Background(), "SELECT id FROM tb_project WHERE id=(SELECT max(id) FROM tb_project)").Scan(&id)

	return c.Redirect(http.StatusMovedPermanently,"/")
}

func deleteProject (c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	_, err := connection.Conn.Exec(context.Background(), "DELETE FROM tb_project WHERE id=$1", id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	return c.Redirect(http.StatusMovedPermanently, "/")
}


func editProject (c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	var isiProject = Project{}

	for i, data := range dataProject{
		if id == i {
			isiProject = Project{
				Id : id,
				Name: data.Name,
				StarDate: data.StarDate,
				EndDate: data.EndDate,
				Duration: data.Duration,
				Detail: data.Detail,
				Playstore: data.Playstore,
				Android: data.Android,
				Java: data.Java,
				React: data.React,
			}
		}
	}

	data := map[string]interface{}{
		"Project": isiProject,
	}

	var tmpl, err = template.ParseFiles("views/update.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return tmpl.Execute(c.Response(), data)
}

func updateProject (c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))	

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
	
	var updateProject = Project {
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

	dataProject[id] = updateProject

	return c.Redirect(http.StatusMovedPermanently, "/")
}