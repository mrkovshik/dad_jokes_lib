package main

import (
	"encoding/json"
	"fmt"
	"text/template"
	"net/http"
	"os"
)

type joke struct {
	
	Success bool `json:"success"`
		Body    []struct {
			ID        string        `json:"_id"`
			Setup     string        `json:"setup"`
			Punchline string        `json:"punchline"`
			Type      string        `json:"type"`
			Likes     []interface{} `json:"likes"`
			Author    struct {
				Name string      `json:"name"`
				ID   interface{} `json:"id"`
			} `json:"author"`
			Approved      bool   `json:"approved"`
			Date          int    `json:"date"`
			Nsfw          bool   `json:"NSFW"`
			ShareableLink string `json:"shareableLink"`
		} `json:"body"`
	}

func main() {

	const templ = `
	Today's amazing joke for you:

	- {{.Setup }}
	
	- {{.Punchline}}
	
	
	`
	var report = template.Must(template.New("jokeTempl").Parse(templ))

	url := "https://dad-jokes.p.rapidapi.com/random/joke"

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("X-RapidAPI-Key",(os.Getenv("RAPID_API_KEY")))
	req.Header.Add("X-RapidAPI-Host", "dad-jokes.p.rapidapi.com")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()


	var awesomeJoke joke
	fmt.Println("res.Body closed:", res.Body == nil)


	if err := json.NewDecoder(res.Body).Decode(&awesomeJoke); err != nil {
		fmt.Println("oops %v",err)
		}
		fmt.Printf("%+v", awesomeJoke.Body[0])
	if err := report.Execute(os.Stdout, awesomeJoke.Body[0]); err != nil {
		fmt.Println(err)
		}

}