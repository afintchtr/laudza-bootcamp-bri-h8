package funcs

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func ShowPost() {
	res, err := http.Get("http://jsonplaceholder.typicode.com/posts/1")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res.Body)
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	sb := string(body)
	log.Println(sb)
}

func StorePost() {
	data := map[string]interface{}{
		"title": "Laudza",
		"body": "Jordan",
		"userId": 1, 
	}

	requestJson, err := json.Marshal(data)
	client := &http.Client{}

	if err != nil {
		log.Fatal(err)
	}

	req, err := http.NewRequest("POST", "http://jsonplaceholder.typicode.com/posts", bytes.NewBuffer(requestJson))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		log.Fatal(err)
	}

	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(string(body))
}