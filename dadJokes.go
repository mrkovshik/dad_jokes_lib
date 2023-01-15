package dadJokes

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"text/template"
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

func makeJoke() string {
	const templ = `
	Today's amazing joke for you:

	- {{.Setup }}
	
	- {{.Punchline}}
	
	
	`
	var report = template.Must(template.New("jokeTempl").Parse(templ))
	url := "https://dad-jokes.p.rapidapi.com/random/joke"
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("X-RapidAPI-Key", (os.Getenv("RAPID_API_KEY")))
	req.Header.Add("X-RapidAPI-Host", "dad-jokes.p.rapidapi.com")
	res, err := http.DefaultClient.Do(req)
	if err != nil{
		fmt.Println("Error while building request", err)
	}
	defer res.Body.Close()
	var awesomeJoke joke
	if err := json.NewDecoder(res.Body).Decode(&awesomeJoke); err != nil {
		fmt.Println("Error decoding json", err)
	}
	fmt.Printf("%+v", awesomeJoke.Body[0])
	var buffer bytes.Buffer

	if err := report.Execute(&buffer, awesomeJoke.Body[0]); err != nil {
		fmt.Println(err)
	}
	return buffer.String()

}
