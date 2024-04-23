package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"text/template"
	"time"
)

type Character struct {
	Info struct {
		Count int    `json:"count"`
		Pages int    `json:"pages"`
		Next  string `json:"next"`
		Prev  any    `json:"prev"`
	} `json:"info"`
	Results []struct {
		ID      int    `json:"id"`
		Name    string `json:"name"`
		Status  string `json:"status"`
		Species string `json:"species"`
		Type    string `json:"type"`
		Gender  string `json:"gender"`
		Origin  struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"origin"`
		Location struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"location"`
		Image   string    `json:"image"`
		Episode []string  `json:"episode"`
		URL     string    `json:"url"`
		Created time.Time `json:"created"`
	} `json:"results"`
	Data *Character
	Req  *Character
}

func main() {
	//mise en place des templates
	tmpl1 := template.Must(template.ParseFiles("./pages/index.html"))
	tmpl2 := template.Must(template.ParseFiles("./pages/request.html"))
	//tmpl3 := template.Must(template.ParseFiles("./pages/character.html"))
	//handle func
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			errorHandler(w, r, http.StatusNotFound)
			return

		}
		if r.Method != http.MethodPost {
			tmpl1.Execute(w, nil)
			return
		}
	})
	http.HandleFunc("/request", func(w http.ResponseWriter, r *http.Request) {

		Send := Character{
			Data: UseInput(r.FormValue("input")),
		}

		// Affichage des personnages dans la template request.html
		tmpl2.Execute(w, Send)
	})
	//lié le css
	fs := http.FileServer(http.Dir("./style/"))
	http.Handle("/style/", http.StripPrefix("/style/", fs))
	//accéder plus facilement au site, cliquer sur le lien dans le terminal
	fmt.Println("running on http://localhost/")
	// sur le port 80
	http.ListenAndServe(":80", nil)

}

// fonction pour gérer les erreurs 404
func errorHandler(w http.ResponseWriter, r *http.Request, status int) {
	w.WriteHeader(status)
	if status == http.StatusNotFound {
		fmt.Fprint(w, "ERROR 404")
	}
}

func characterData() *Character {
	url := "https://rickandmortyapi.com/api/character"

	timeClient := http.Client{
		Timeout: time.Second * 2,
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		panic(err)
	}

	req.Header.Set("User-Agent", "random-user-agent")
	res, getErr := timeClient.Do(req)
	if getErr != nil {
		fmt.Println(getErr)
	}
	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		fmt.Println(readErr)
	}
	NewStruct := Character{}
	jsonErr := json.Unmarshal(body, &NewStruct)
	if jsonErr != nil {
		fmt.Println(jsonErr)
	}

	return &NewStruct

}

func UseInput(input string) *Character {
	url := "https://rickandmortyapi.com/api/character/?name=" + input

	timeClient := http.Client{
		Timeout: time.Second * 2,
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		panic(err)
	}

	req.Header.Set("User-Agent", "random-user-agent")
	res, getErr := timeClient.Do(req)
	if getErr != nil {
		fmt.Println(getErr)
	}
	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		fmt.Println(readErr)
	}
	NewStruct := Character{}
	jsonErr := json.Unmarshal(body, &NewStruct)
	if jsonErr != nil {
		fmt.Println(jsonErr)
	}

	return &NewStruct

}
