package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"oji/connection"

	// "os/user"
	"strconv"
	"text/template"
	"time"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
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

type User struct {
	Id int
	Name string
	Email string
	Password string
}

type SessionData struct {
	IsLogin bool
	Name string
}

var userData = SessionData{}


func main() {
	connection.DatabaseConnect()

	e := echo.New()

	// static file from 'public' directory
	e.Static("/public", "public")

	// to use session using echo
	e.Use(session.Middleware(sessions.NewCookieStore([]byte("session"))))

	e.GET("/hello", hai)
	e.GET("/", home)
	e.GET("/myproject", addProject)
	e.GET("/testimonial", testimonial)
	e.GET("/contact", contact)
	e.GET("/project-detail/:id", projectDetail)
	e.GET("/project-edit/:id", editProject)
	
	// login
	e.GET("/form-login", formLogin)
	e.POST("/login", subLogin)
	e.POST("logout", logout)

	// Register
	e.GET("/form-register", formRegister)
	e.POST("/register", register)
	
	// routing post
	e.POST("/saveproject", saveProject)
	e.POST("/deleteProject/:id", deleteProject)
	e.POST("/updateProject/:id", updateProject)

	e.Logger.Fatal(e.Start("localhost:5000"))
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
	
	sess, _ := session.Get("session", c)

	if sess.Values["isLogin"] != true {
		userData.IsLogin = false
	} else {
		userData.IsLogin = sess.Values["isLogin"].(bool)
		userData.Name = sess.Values["name"].(string)
	}
	
	Projects := map[string]interface{}{
		"Projects" : result,
		"FlashStatus" : sess.Values["status"],
		"FlashMessage" : sess.Values["message"],
		"DataSession" : userData,
	}

	delete(sess.Values, "message")
	delete(sess.Values, "status")
	sess.Save(c.Request(), c.Response())

	var tmpl, err = template.ParseFiles("views/index.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message":err.Error()})
	}
	// connection.Conn.Close(context.Background())
	data.Close()
	
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

	// connection.Conn.Close(context.Background())

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

	err := connection.Conn.QueryRow(context.Background(), "SELECT * FROM tb_project WHERE id=$1", id).Scan(
		&isiProject.Id, &isiProject.Name, &isiProject.StartDateTime, &isiProject.EndDateTime, &isiProject.Duration, &isiProject.Detail, &isiProject.Playstore, &isiProject.Android, &isiProject.Java, &isiProject.React,
	)
	
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message":err.Error()})
	}

	data := map[string]interface{}{
		"Project": isiProject,
		"StarDate" : isiProject.StartDateTime.Format("01-02-2006"),
		"EndDate" : isiProject.EndDateTime.Format("01-02-2006"),
	}

	var tmpl, errTemp = template.ParseFiles("views/update.html")
	
	if errTemp != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"messagek": errTemp.Error()})
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
	
	_, err := connection.Conn.Exec(context.Background(),
			`UPDATE tb_project SET name=$1, star_date=$2, end_date=$3, duration=$4, detail=$5, playstore=$6, android=$7, java=$8, react=$9 WHERE id=$10`,
			name, starDate, endDate, diffUse, detail, playstore, android, java, react, id,
			)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message":err.Error()})
	}

	return c.Redirect(http.StatusMovedPermanently, "/")
}

func formRegister (c echo.Context) error {
	var tmpl, err = template.ParseFiles("views/register.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return tmpl.Execute(c.Response(), nil)
}

func register (c echo.Context) error {
	err := c.Request().ParseForm()
	if err != nil {
		log.Fatal(err)
	}

	name := c.FormValue("input-username")
	email := c.FormValue("input-email")
	password := c.FormValue("input-password")

	passwordHash, _ := bcrypt.GenerateFromPassword([]byte(password), 10)

	_, err = connection.Conn.Exec(context.Background(), "INSERT INTO tb_users(username, email, password) VALUES ($1, $2, $3)", name, email, passwordHash)

	if err != nil {
		redirectWithMessage(c, "Register failed, please try again.", false, "/form-register")
	}

	return redirectWithMessage(c, "Register success!", true, "/form-login")
}

func formLogin (c echo.Context) error {
	sess, _ := session.Get("session", c)
	
	flash := map[string]interface{}{
		"FlashStatus" : sess.Values["status"],
		"FlashMessage" : sess.Values["message"],
	}

	delete(sess.Values, "message")
	delete(sess.Values, "status")
	sess.Save(c.Request(), c.Response())

	var tmpl, err = template.ParseFiles("views/login.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message":err.Error()})
	}

	return tmpl.Execute(c.Response(), flash)	
}

func subLogin (c echo.Context) error {
	err := c.Request().ParseForm()
	if err != nil {
		log.Fatal(err)
	}

	email := c.FormValue("input-email")
	password := c.FormValue("input-password")

	user := User{}
	err = connection.Conn.QueryRow(context.Background(), "SELECT * FROM tb_users WHERE email=$1", email).Scan(&user.Id,  &user.Name, &user.Email, &user.Password)

	if err != nil {
		return redirectWithMessage(c, "Email Incorrect!", false, "/form-login")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return redirectWithMessage(c,"Password Incorrect!", false, "/form-login")
	}
	
	sess, _ := session.Get("session", c)
	sess.Options.MaxAge = 10800 //3 jam
	sess.Values["message"] = "Login success!"
	sess.Values["status"] = true
	sess.Values["name"] = user.Name
	sess.Values["email"] = user.Email
	sess.Values["id"] = user.Id
	sess.Values["isLogin"] = true
	sess.Save(c.Request(), c.Response())

	return c.Redirect(http.StatusMovedPermanently, "/")
	
}

func logout(c echo.Context) error {
	sess, _ := session.Get("session", c)
	sess.Options.MaxAge = -1
	sess.Save(c.Request(), c.Response())

	return c.Redirect(http.StatusMovedPermanently, "/")
}

func redirectWithMessage(c echo.Context, message string, status bool, path string) error {
	sess, _ := session.Get("session", c)
	sess.Values["message"] = message
	sess.Values["status"] = status
	sess.Save(c.Request(), c.Response())
	return c.Redirect(http.StatusMovedPermanently, path)
}