package config

import (
	admin "BlogGO/admin/controllers"
	site "BlogGO/site/controllers"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func Routes() *httprouter.Router {
	r := httprouter.New()
	// ADMIN
	// Blog Posts
	r.GET("/admin", admin.Dashboard{}.Index)
	r.GET("/admin/yeni-ekle", admin.Dashboard{}.NewItem)
	r.POST("/admin/add", admin.Dashboard{}.Add)
	r.GET("/admin/delete/:id", admin.Dashboard{}.Delete)
	r.GET("/admin/edit/:id", admin.Dashboard{}.Edit)
	r.POST("/admin/update/:id", admin.Dashboard{}.Update)

	//USEROPS
	r.GET("/admin/login", admin.Userops{}.Index)
	r.POST("/admin/do_login", admin.Userops{}.Login)
	r.GET("/admin/logout", admin.Userops{}.Logout)

	//CATEGORIES
	r.GET("/admin/kategoriler", admin.Categories{}.Index)
	r.POST("/admin/kategoriler/add", admin.Categories{}.Add)
	r.GET("/admin/kategoriler/delete/:id", admin.Categories{}.Delete)

	//HOMEPAGE
	r.GET("/", site.Homepage{}.Index)
	r.GET("/yazilar/:slug", site.Homepage{}.Detail)

	// SERVE FILES
	r.ServeFiles("/admin/assets/*filepath", http.Dir("admin/assets/"))
	r.ServeFiles("/assets/*filepath", http.Dir("site/assets"))
	r.ServeFiles("/uploads/*filepath", http.Dir("uploads"))
	return r
}
