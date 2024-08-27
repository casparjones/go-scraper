package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/chromedp/chromedp"
	"io"
	"net/http"
)

// RequestBody definiert die Struktur des erwarteten JSON im POST-Body.
type RequestBody struct {
	Method string `json:"method"`
	URL    string `json:"url"`
	Body   string `json:"body"` // Für GET-Anfragen wird dieser Teil ignoriert.
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Server läuft auf :8080")
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Nur POST ist erlaubt", http.StatusMethodNotAllowed)
		println("Nur POST ist erlaubt")
		return
	}

	var reqBody RequestBody
	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Error reading body:", err)
		return
	}

	err = json.Unmarshal(bodyBytes, &reqBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		println("error decoding JSON: " + err.Error())
		return
	}

	println("request JSON: " + string(bodyBytes))

	// Puppeteer-Logik (chromedp)
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	var siteContent string
	err = chromedp.Run(ctx,
		chromedp.Navigate(reqBody.URL),
		chromedp.OuterHTML("html", &siteContent),
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte(siteContent))
}
