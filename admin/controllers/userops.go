package controllers

import (
	"BlogGO/admin/helpers"
	"BlogGO/admin/models"
	"crypto/sha256"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"html/template"
	"net/http"
)

type Userops struct{}

func (userops Userops) Index(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	view, err := template.ParseFiles(helpers.INclude("userops/login")...)
	if err != nil {
		fmt.Println(err)
		return
	}
	data := make(map[string]interface{})
	data["Alert"] = helpers.GetAlert(w, r)
	view.ExecuteTemplate(w, "index", data)
}

func (userops Userops) Login(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	username := r.FormValue("username")
	password := fmt.Sprintf("%x", sha256.Sum256([]byte(r.FormValue("password"))))
	user := models.User{}.Get("Username= ?  AND password= ?", username, password)
	fmt.Println(user)
	if user.Username == username && user.Password == password {
		//login
		helpers.SetUser(w, r, username, password)
		helpers.SetAlert(w, r, "Hoşgeldiniz")
		http.Redirect(w, r, "/admin", http.StatusSeeOther)
	} else {
		//denied
		helpers.SetAlert(w, r, "Yanlış Kullanıcı Adı veya Şifre!")
		http.Redirect(w, r, "/admin/login", http.StatusSeeOther)
	}
}

func (userops Userops) Logout(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	helpers.RemoveUser(w, r)
	helpers.SetAlert(w, r, "Çıkış Yapıldı...")
	http.Redirect(w, r, "/admin/login", http.StatusSeeOther)
}
