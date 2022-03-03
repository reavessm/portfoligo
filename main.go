/**
 * File: main.go
 * Written by: Stephen M. Reaves
 * Created on: Thu, 03 Mar 2022
 */

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"myStaticWebsite/config"
	"net/http"
	"text/template" // not html/template due to rendering of comments

	"gopkg.in/yaml.v2"
)

const (
	portNumber   string = "8080"
	templateFile string = "template/home.tmpl"
	configFile   string = "config/user.yml"
)

func readYaml(path string) config.Config {
	yfile, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalln(err)
	}

	config := config.Config{}
	err = yaml.Unmarshal(yfile, &config)
	if err != nil {
		log.Fatalln(err)
	}

	return config
}

func main() {
	mux := http.NewServeMux()
	fmt.Println(fmt.Sprintf("Starting server on port %s", portNumber))

	tmpl := template.Must(template.ParseFiles(templateFile))

	// Parse yaml config for main landing page
	mux.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		tmpl.Execute(rw, readYaml(configFile))
	})

	// CSS and Resumes
	fs := http.FileServer(http.Dir("./public"))
	mux.Handle("/static/", fs)

	log.Fatalln(http.ListenAndServe(fmt.Sprintf(":%s", portNumber), mux))
}
