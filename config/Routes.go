package config

import (
	"github.com/julienschmidt/httprouter"
	admin "goblog/admin/controllers"
	site "goblog/site/controllers"
	"net/http"
)

func Routes() *httprouter.Router {
	r := httprouter.New()
	//ADMIN
	//Blog Posts
	r.GET("/admin",admin.Dashboard{}.Index)
	r.GET("/admin/yeni-ekle",admin.Dashboard{}.NewItem)
	r.POST("/admin/add",admin.Dashboard{}.Add)
	r.GET("/admin/delete/:id",admin.Dashboard{}.Delete)
	r.GET("/admin/edit/:id",admin.Dashboard{}.Edit)
	r.POST("/admin/update/:id",admin.Dashboard{}.Update)

	//Categories
	r.GET("/admin/kategoriler",admin.Categories{}.Index)
	r.POST("/admin/kategoriler/add",admin.Categories{}.Add)
	r.GET("/admin/kategoriler/delete/:id",admin.Categories{}.Delete)

	//Userops
	r.GET("/admin/login",admin.Userops{}.Index)
	r.POST("/admin/do_login",admin.Userops{}.Login)
	r.GET("/admin/logout",admin.Userops{}.Logout)

	//SITE
	//Homepage
	r.GET("/",site.Homepage{}.Index)
	r.GET("/yazilar/:slug",site.Homepage{}.Detail)

	// SERVE FILES
	r.ServeFiles("/admin/assets/*filepath",http.Dir("admin/assets"))
	r.ServeFiles("/assets/*filepath",http.Dir("site/assets"))
	r.ServeFiles("/uploads/*filepath",http.Dir("uploads"))
	return r
}
