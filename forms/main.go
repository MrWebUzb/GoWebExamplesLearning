package main

import (
	"html/template"
	"net/http"
)

// ContactForm ...
type ContactForm struct {
	Firstname string
	Email     string
	Message   string
}

// ContactPage ...
type ContactPage struct {
	PageTitle string
	Success   bool
	Form      ContactForm
}

func main() {
	tpl := template.Must(template.ParseFiles("form.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		contactPage := &ContactPage{
			PageTitle: "Basic Contact Form",
			Success:   false,
			Form: ContactForm{
				Firstname: "",
				Email:     "",
				Message:   "",
			},
		}

		if r.Method == http.MethodPost {
			details := ContactForm{
				Firstname: r.PostFormValue("firstname"),
				Email:     r.PostFormValue("email"),
				Message:   r.PostFormValue("message"),
			}
			contactPage.Success = true
			contactPage.Form = details
		}

		tpl.Execute(w, contactPage)
	})

	http.ListenAndServe(":8080", nil)
}
