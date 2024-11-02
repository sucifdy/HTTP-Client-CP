package main

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// Struct untuk menampung respons dari API Quotable (GET)
type Quotes struct {
	Tags   []string `json:"tags"`
	Author string   `json:"author"`
	Quote  string   `json:"content"`
}

// Fungsi untuk melakukan GET request ke API Quotable
func ClientGet() ([]Quotes, error) {
	// Mengatur client dengan pengabaian verifikasi sertifikat
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr, Timeout: 10 * time.Second}

	// Mengirim permintaan GET
	resp, err := client.Get("https://api.quotable.io/quotes/random?limit=3")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Dekode JSON respons ke dalam struct Quotes
	var quotes []Quotes
	if err := json.NewDecoder(resp.Body).Decode(&quotes); err != nil {
		return nil, err
	}

	return quotes, nil
}

// Struct untuk body JSON yang dikirim pada POST request
type data struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

// Struct untuk menampung respons dari API Postman Echo (POST)
type Postman struct {
	Data data   `json:"data"`
	Url  string `json:"url"`
}

// Fungsi untuk melakukan POST request ke API Postman Echo
func ClientPost() (Postman, error) {
	// Menyiapkan body JSON untuk POST request
	postBody, _ := json.Marshal(map[string]string{
		"name":  "Dion",
		"email": "dionbe2022@gmail.com",
	})
	responseBody := bytes.NewBuffer(postBody)

	client := &http.Client{Timeout: 10 * time.Second}

	// Mengirim permintaan POST
	resp, err := client.Post("https://postman-echo.com/post", "application/json", responseBody)
	if err != nil {
		return Postman{}, err
	}
	defer resp.Body.Close()

	// Dekode JSON respons ke dalam struct Postman
	var postmanResp Postman
	if err := json.NewDecoder(resp.Body).Decode(&postmanResp); err != nil {
		return Postman{}, err
	}

	return postmanResp, nil
}

func main() {
	// Menjalankan fungsi ClientGet dan menampilkan hasilnya
	quotes, err := ClientGet()
	if err != nil {
		fmt.Println("Error in GET request:", err)
	} else {
		fmt.Println("Quotes from API:")
		for _, quote := range quotes {
			fmt.Printf("Quote: %s\nAuthor: %s\nTags: %v\n\n", quote.Quote, quote.Author, quote.Tags)
		}
	}

	// Menjalankan fungsi ClientPost dan menampilkan hasilnya
	postmanResp, err := ClientPost()
	if err != nil {
		fmt.Println("Error in POST request:", err)
	} else {
		fmt.Printf("Postman Response:\nData: %v\nURL: %s\n", postmanResp.Data, postmanResp.Url)
	}
}
