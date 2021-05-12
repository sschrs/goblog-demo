package controllers

import (
	"crypto/sha256"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"goblog/admin/helpers"
	"goblog/admin/models"
	"html/template"
	"net/http"
)

type Userops struct {}

func (userops Userops) Index(w http.ResponseWriter,r *http.Request,params httprouter.Params){
	view,err := template.ParseFiles(helpers.Include("userops/login")...)
	if err != nil {
		fmt.Println(err)
		return
	}
	data := make(map[string]interface{})
	data["Alert"] = helpers.GetAlert(w,r)
	view.ExecuteTemplate(w,"index",data)
}

func (userops Userops) Login(w http.ResponseWriter,r *http.Request,params httprouter.Params){
	username := r.FormValue("username")
	password := fmt.Sprintf("%x",sha256.Sum256([]byte(r.FormValue("password"))))

	user := models.User{}.Get("username = ? AND password = ?",username,password)
	if (user.Username == username && user.Password == password){
		helpers.SetUser(w,r,username,password)
		helpers.SetAlert(w,r,"Hoşgeldiniz")
		http.Redirect(w,r,"/admin",http.StatusSeeOther)
	}else{
		helpers.SetAlert(w,r,"Yanlış Kullanıcı Adı veya Şifre")
		http.Redirect(w,r,"/admin/login",http.StatusSeeOther)
	}
}

func (userops Userops) Logout(w http.ResponseWriter,r *http.Request,params httprouter.Params)  {
	helpers.RemoveUser(w,r)
	helpers.SetAlert(w,r,"Hoşçakalın")
	http.Redirect(w,r,"/admin/login",http.StatusSeeOther)
}