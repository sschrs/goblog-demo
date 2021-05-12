package controllers

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"goblog/site/helpers"
	"goblog/site/models"
	"html/template"
	"net/http"
	"time"
)

type Homepage struct {}

func (homepage Homepage) Index(w http.ResponseWriter,r *http.Request,params httprouter.Params)  {
	view,err := template.New("index").Funcs(template.FuncMap{
		"getCategory" : func(categoryID int) string {
			return models.Category{}.Get(categoryID).Title
		},
		"getDate" : func(t time.Time) string {
			return fmt.Sprintf("%02d.%02d.%d",t.Day(),int(t.Month()),t.Year())
		},
	}).ParseFiles(helpers.Include("homepage/list")...)
	if err != nil {
		fmt.Println(err)
		return
	}
	data := make(map[string]interface{})
	data["Posts"] = models.Post{}.GetAll()
	view.ExecuteTemplate(w,"index",data)
}

func (homepage Homepage) Detail(w http.ResponseWriter,r *http.Request,params httprouter.Params)  {
	view,err := template.ParseFiles(helpers.Include("homepage/detail")...)
	if err != nil {
		fmt.Println(err)
		return
	}
	data := make(map[string]interface{})
	data["Post"] = models.Post{}.Get("slug = ?",params.ByName("slug"))
	view.ExecuteTemplate(w,"index",data)
}