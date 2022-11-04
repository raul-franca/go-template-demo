package main

import (
	"html/template"
	"net/http"
)

type Contato struct {
	ID    string
	Name  string
	Email string
	Ativo bool
}

type Contatos []Contato

//
//func NewContato(name, email string) *Contato {
//	return &Contato{
//		ID:    uuid.New().String(),
//		Name:  name,
//		Email: email,
//		Ativo: true,
//	}
//}

func (contato Contato) GetContato() string {
	return "#-> Contato id: " + contato.ID + " " +
		"\n Nome: " + contato.Name + ", " +
		"e-mail: " + contato.Email
}

func main() {

	var contato Contato
	contato.Name = "Tereza"
	contato.Email = "tereza@email.com"
	contato.Ativo = true

	contato2 := Contato{ID: "033030302", Name: "Aurora", Email: "email@gmail.com", Ativo: true}

	contatos := Contatos{
		{"123", "Raul", "email@gmail.com", true},
		{"13", "Lula", "Lula@gmail.com", true},
		{"432", "Maria", "Maria@gmail.com", true},
	}
	contatos = append(
		contatos,
		contato,
		contato2,
	)
	templates := []string{ // Arquivos que vão compor o template
		"01/templates/header.html",
		"01/templates/content.html",
		"01/templates/footer.html",
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		t := template.New("content.html")
		t, err := t.ParseFiles(templates...)
		if err != nil {
			panic(err)
		}

		temp := template.Must(t, nil)
		err = temp.Execute(w, contatos)
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
