package main

import (
    "github.com/magiconair/properties"
    "github.com/gorilla/mux"
    "github.com/newrelic/go-agent"
    "html/template"
    "net/http"
    "strings"
    "fmt"
    "log"
    "os"
)

func getProperties(host string) map[string]string {

    var p map[string]string

    if strings.Contains(host, "wales") {
        p = properties.MustLoadFile("static/wales.properties", properties.UTF8).Map()
    } else {
        p = properties.MustLoadFile("static/cymru.properties", properties.UTF8).Map()
    }

    return p
}

func handler(w http.ResponseWriter, r *http.Request) {
    t, _ := template.ParseFiles("templates/index.tmpl")
    m := getProperties(r.Host)
    t.Execute(w, m)
}

func main() {
    newRelicLicenseKey := os.Getenv("NEWRELIC_LICENSE_KEY")
    newRelicApplicationName := os.Getenv("NEWRELIC_APP_NAME")

    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    config := newrelic.NewConfig(newRelicApplicationName, newRelicLicenseKey)
    app, err := newrelic.NewApplication(config)

    if err != nil {
        log.Fatal("Unable to create new relic application: ", err)
        os.Exit(1)
    }

    r := mux.NewRouter()
    r.HandleFunc(newrelic.WrapHandleFunc(app,"/", handler))
    r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
    http.ListenAndServe(fmt.Sprintf(":%s",port), r)
}
