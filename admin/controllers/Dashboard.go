package controllers

import (
	"BlogGO/admin/helpers"
	"BlogGO/admin/models"
	"fmt"
	"github.com/gosimple/slug"
	"github.com/julienschmidt/httprouter"
	"html/template"
	"io"
	"net/http"
	"os"
	"strconv"
)

type Dashboard struct{}

func (dashboard Dashboard) Index(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	//fmt.Println(helpers.SetAlert(w, r, "message"))
	//fmt.Println(helpers.GetAlert(w, r))

	if !helpers.CheckUser(w, r) {
		return
	}

	view, err := template.New("index").Funcs(template.FuncMap{
		"getCategory": func(categoryID int) string {
			return models.Category{}.Get(categoryID).Title
		},
	}).ParseFiles(helpers.INclude("dashboard/list")...)

	if err != nil {
		fmt.Println(err)
		return
	}
	data := make(map[string]interface{})
	data["Posts"] = models.Post{}.GetAll()
	data["Alert"] = helpers.GetAlert(w, r)
	view.ExecuteTemplate(w, "index", data)
}

func (dashboard Dashboard) NewItem(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	if !helpers.CheckUser(w, r) {
		return
	}
	view, err := template.ParseFiles(helpers.INclude("dashboard/add")...)
	if err != nil {
		fmt.Println(err)
		return
	}
	data := make(map[string]interface{})
	data["Categories"] = models.Category{}.GetAll()
	view.ExecuteTemplate(w, "index", data)
}

func (dashboard Dashboard) Add(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	if !helpers.CheckUser(w, r) {
		return
	}
	title := r.FormValue("blog-title")
	slug := slug.Make(title)
	description := r.FormValue("blog-desc")
	categoryID, err := strconv.Atoi(r.FormValue("blog-category"))
	if err != nil {
		fmt.Println(err)
		return
	}
	content := r.FormValue("blog-content")
	//UPLOAD
	pictureURL := ""
	r.ParseMultipartForm(10 << 20)
	file, header, _ := r.FormFile("blog-picture")
	if file != nil {
		f, err := os.OpenFile("uploads/"+header.Filename, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println(err)
			return
		}
		_, err = io.Copy(f, file)
		if err != nil {
			fmt.Println(err)
			return
		}
		pictureURL = "uploads/" + header.Filename
	}

	//UPLOAD END
	models.Post{
		Title:       title,
		Slug:        slug,
		Description: description,
		CategoryID:  categoryID,
		Content:     content,
		PictureURL:  pictureURL,
	}.Add()

	err = helpers.SetAlert(w, r, "Kaydedildi")
	if err != nil {
		fmt.Println(err)
	}
	http.Redirect(w, r, "/admin", http.StatusSeeOther)
	//TODO alert
}

func (dashboard Dashboard) Delete(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	if !helpers.CheckUser(w, r) {
		return
	}
	post := models.Post{}.Get(params.ByName("id"))
	post.Delete()
	http.Redirect(w, r, "/admin", http.StatusSeeOther)
}
func (dashboard Dashboard) Edit(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	if !helpers.CheckUser(w, r) {
		return
	}
	view, err := template.ParseFiles(helpers.INclude("dashboard/edit")...)

	if err != nil {
		fmt.Println(err)
		return
	}
	data := make(map[string]interface{})
	data["Post"] = models.Post{}.Get(params.ByName("id"))
	data["Categories"] = models.Category{}.GetAll()
	view.ExecuteTemplate(w, "index", data)
	//http.Redirect(w, r, "/admin", http.StatusSeeOther)
}
func (dashboard Dashboard) Update(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	if !helpers.CheckUser(w, r) {
		return
	}
	post := models.Post{}.Get(params.ByName("id"))
	title := r.FormValue("blog-title")
	slug := slug.Make(title)
	description := r.FormValue("blog-desc")
	categoryID, err := strconv.Atoi(r.FormValue("blog-category"))
	if err != nil {
		fmt.Println(err)
		return
	}
	content := r.FormValue("blog-content")
	isSelected := r.FormValue("isSelected")
	var picture_url string

	if isSelected == "1" {
		//UPLOAD
		r.ParseMultipartForm(10 << 20)
		file, header, err := r.FormFile("blog-picture")
		if err != nil {
			fmt.Println(err)
			return
		}
		f, err := os.OpenFile("uploads/"+header.Filename, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println(err)
			return
		}
		_, err = io.Copy(f, file)
		if err != nil {
			fmt.Println(err)
			return
		}
		//UPLOAD END
		picture_url = "uploads/" + header.Filename
		os.Remove(post.PictureURL)
	} else {
		picture_url = post.PictureURL
	}
	post.Updates(models.Post{
		Title:       title,
		Slug:        slug,
		Description: description,
		CategoryID:  categoryID,
		Content:     content,
		PictureURL:  picture_url,
	})
	http.Redirect(w, r, "/admin", http.StatusSeeOther)
}
