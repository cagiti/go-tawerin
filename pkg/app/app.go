package app

import (
	"fmt"
	"github.com/cagiti/go-tawerin/pkg/util"
	"github.com/gorilla/mux"
	"github.com/magiconair/properties"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"strings"
)

type App struct {
	Router *mux.Router
}

func (a *App) Initialize() {
	a.Router = mux.NewRouter()
	a.initializeRoutes()
}

func (a *App) Run(addr string) {
	log.Printf("Starting... on %s", addr)
	log.Fatal(http.ListenAndServe(addr, a.Router))
}

func (a *App) getProperties(host string) map[string]string {
	var p map[string]string
	if strings.Contains(host, "wales") {
		p = properties.MustLoadFile(filepath.Join(util.StaticDir(), "wales.properties"), properties.UTF8).Map()
	} else {
		p = properties.MustLoadFile(filepath.Join(util.StaticDir(), "cymru.properties"), properties.UTF8).Map()
	}
	return p
}

func (a *App) handler(w http.ResponseWriter, r *http.Request) {
	m := a.getProperties(r.Host)
	if r.URL.Path == "/" {
		t, _ := template.ParseFiles(filepath.Join(util.TemplatesDir(), "index.tmpl"))
		err := t.Execute(w, m)
		if err != nil {
			log.Print("Unable to parse template: ", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else if r.URL.Path == "/ping" {
		fmt.Fprintf(w, "OK")
	} else {
		t, _ := template.ParseFiles(filepath.Join(util.TemplatesDir(), fmt.Sprintf("%s.tmpl", r.URL.Path)))
		err := t.Execute(w, m)
		if err != nil {
			log.Print("Unable to parse template: ", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func (a *App) initializeRoutes() {
	a.Router.HandleFunc("/", a.handler)
	a.Router.HandleFunc("/ytim", a.handler)
	a.Router.HandleFunc("/band", a.handler)
	a.Router.HandleFunc("/dawnsfeydd", a.handler)
	a.Router.HandleFunc("/rolau", a.handler)
	a.Router.HandleFunc("/oriel", a.handler)
	a.Router.HandleFunc("/perfformiadau", a.handler)
	a.Router.HandleFunc("/cysylltu", a.handler)
	a.Router.HandleFunc("/error", a.handler)
	a.Router.HandleFunc("/result", a.handler)
	a.Router.HandleFunc("/ping", a.handler)
	a.Router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(util.StaticDir()))))
}
