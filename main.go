package main

import (
	adminModels "BlogGO/admin/models"
	"BlogGO/config"
	"fmt"
	"gopkg.in/gomail.v2"
	_ "gopkg.in/gomail.v2"
	"html"
	"net/http"
	"net/mail"
)

func main() {
	adminModels.Post{}.Migrate()
	adminModels.User{}.Migrate()
	adminModels.Category{}.Migrate()
	http.ListenAndServe("localhost:8080", config.Routes())

	http.HandleFunc("/contact", contactHandler)

}
func contactHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method.", http.StatusMethodNotAllowed)
		return
	}

	if err := r.ParseForm(); err != nil {
		http.Error(w, "Error parsing form data.", http.StatusBadRequest)
		return
	}

	name := html.EscapeString(r.PostFormValue("name"))
	email := html.EscapeString(r.PostFormValue("email"))
	phone := html.EscapeString(r.PostFormValue("phone"))
	message := html.EscapeString(r.PostFormValue("message"))

	if name == "" || email == "" || phone == "" || message == "" {
		http.Error(w, "No arguments provided!", http.StatusBadRequest)
		return
	}

	// Validate email format
	if _, err := mail.ParseAddress(email); err != nil {
		http.Error(w, "Invalid email format.", http.StatusBadRequest)
		return
	}

	if err := sendEmail(name, email, phone, message); err != nil {
		http.Error(w, "Failed to send email.", http.StatusInternalServerError)
		return
	}

	fmt.Fprintln(w, "Message sent successfully.")
}

func sendEmail(name, email, phone, message string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", "noreply@yourdomain.com")
	m.SetHeader("To", "yourname@yourdomain.com")
	m.SetHeader("Reply-To", email)
	m.SetHeader("Subject", fmt.Sprintf("Website Contact Form: %s", name))
	m.SetBody("text/plain", fmt.Sprintf("You have received a new message from your website contact form.\n\nHere are the details:\n\nName: %s\n\nEmail: %s\n\nPhone: %s\n\nMessage:\n%s", name, email, phone, message))

	d := gomail.NewDialer("smtp.yourdomain.com", 587, "your_smtp_username", "your_smtp_password")

	if err := d.DialAndSend(m); err != nil {
		return fmt.Errorf("could not send email: %v", err)
	}
	return nil
}

/*
	post := adminModels.Post{}.Get(1)
	post.Update("Title", "Gooo ile web programlama")
	fmt.Println(post.Title)
	//post.Delete(post.PictureURL)
	posts := adminModels.Post{}.GetAll("Description=?", "")
	fmt.Println(posts)

	posts = adminModels.Post{}.GetAll("Description=?", "deneme12")
	fmt.Println(posts)

	post = adminModels.Post{}.Get(1)

	post.Updates(adminModels.Post{Title: "Goooo ile web programlama", Content: "Go nedir ?", Slug: "go-nedir"})
*/
