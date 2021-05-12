package controllers

import (
	"fmt"
	"github.com/gosimple/slug"
	"github.com/julienschmidt/httprouter"
	"goblog/admin/helpers"
	"goblog/admin/models"
	"html/template"
	"io"
	"net/http"
	"os"
	"strconv"
)

type Dashboard struct {}

func (dashboard Dashboard) Index(w http.ResponseWriter,r *http.Request,params httprouter.Params) {
	if !helpers.CheckUser(w,r) {
		return
	}
	view,err := template.New("index").Funcs(template.FuncMap{
		"getCategory" : func(categoryID int) string {
			return models.Category{}.Get(categoryID).Title
		},
	}).ParseFiles(helpers.Include("dashboard/list")...)
	if err != nil {
		fmt.Println(err)
		return
	}
	data := make(map[string]interface{})
	data["Posts"] = models.Post{}.GetAll()
	data["Alert"] = helpers.GetAlert(w,r)
	view.ExecuteTemplate(w,"index",data)
}

func (dashboard Dashboard) NewItem(w http.ResponseWriter,r *http.Request,params httprouter.Params)  {
	if !helpers.CheckUser(w,r) {
		return
	}
	view,err := template.ParseFiles(helpers.Include("dashboard/add")...)
	if err != nil {
		fmt.Println(err)
		return
	}
	data := make(map[string]interface{})
	data["Categories"] = models.Category{}.GetAll()
	view.ExecuteTemplate(w,"index",data)
}

func (dashboard Dashboard) Add(w http.ResponseWriter,r *http.Request,params httprouter.Params)  {
	if !helpers.CheckUser(w,r) {
		return
	}
	title := r.FormValue("blog-title")
	slug := slug.Make(title)
	description := r.FormValue("blog-desc")
	categoryID,_ := strconv.Atoi(r.FormValue("blog-category"))
	content := r.FormValue("blog-content")

	//Upload
	r.ParseMultipartForm(10 << 20)
	file,header,err := r.FormFile("blog-picture")
	if err != nil {
		fmt.Println(err)
		return
	}
	f,err := os.OpenFile("uploads/"+header.Filename,os.O_WRONLY|os.O_CREATE,0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	_,err = io.Copy(f,file)
	// Upload End
	if err != nil {
		fmt.Println(err)
		return
	}
	models.Post{
		Title: title,
		Slug: slug,
		Description: description,
		CategoryID: categoryID,
		Content: content,
		Picture_url: "uploads/" + header.Filename,
	}.Add()
	helpers.SetAlert(w,r,"Kayıt Başarıyla Eklendi")
	http.Redirect(w,r,"/admin",http.StatusSeeOther)
}

func (dashboard Dashboard) Delete(w http.ResponseWriter,r *http.Request,params httprouter.Params)  {
	if !helpers.CheckUser(w,r) {
		return
	}
	post := models.Post{}.Get(params.ByName("id"))
	post.Delete()
	http.Redirect(w,r,"/admin",http.StatusSeeOther)
}

func (dashboard Dashboard) Edit(w http.ResponseWriter,r *http.Request,params httprouter.Params)  {
	if !helpers.CheckUser(w,r) {
		return
	}
	view,err := template.ParseFiles(helpers.Include("dashboard/edit")...)
	if err != nil {
		fmt.Println(err)
		return
	}
	data := make(map[string]interface{})
	data["Post"] = models.Post{}.Get(params.ByName("id"))
	data["Categories"] = models.Category{}.GetAll()
	view.ExecuteTemplate(w,"index",data)
}

func (dashboard Dashboard) Update(w http.ResponseWriter,r *http.Request,params httprouter.Params)  {
	if !helpers.CheckUser(w,r) {
		return
	}
	post := models.Post{}.Get(params.ByName("id"))
	title := r.FormValue("blog-title")
	slug := slug.Make(title)
	description := r.FormValue("blog-desc")
	categoryID,_ := strconv.Atoi(r.FormValue("blog-category"))
	content := r.FormValue("blog-content")
	is_selected := r.FormValue("is_selected")
	var picture_url string

	if is_selected == "1" {
		//Upload
		r.ParseMultipartForm(10 << 20)
		file,header,err := r.FormFile("blog-picture")
		if err != nil {
			fmt.Println(err)
			return
		}
		f,err := os.OpenFile("uploads/"+header.Filename,os.O_WRONLY|os.O_CREATE,0666)
		if err != nil {
			fmt.Println(err)
			return
		}
		_,err = io.Copy(f,file)
		// Upload End
		picture_url = "uploads/" + header.Filename
		os.Remove(post.Picture_url)
	}else{
		picture_url = post.Picture_url
	}

	post.Updates(models.Post{
		Title: title,
		Slug: slug,
		Description: description,
		CategoryID: categoryID,
		Content: content,
		Picture_url: picture_url,
	})
	http.Redirect(w,r,"/admin/edit/"+params.ByName("id"),http.StatusSeeOther)
}