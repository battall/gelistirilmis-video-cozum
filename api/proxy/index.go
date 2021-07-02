package handler

import (
	"net/http"
	"io"
	"log"
)

var Publishers = map[string]string{
  "cap": "capyayinlari",
  "hiz-ve-renk" : "hizrenk",
  "fen-bilimleri" : "fenbilimleri",
  "apotemi" : "apotemi",
}

func Handler(w http.ResponseWriter, r *http.Request) {
  log.Println("[REQUEST]")

	publisher := Publishers[r.URL.Query().Get("publisher")]
  path := r.URL.Query().Get("path")

	log.Println("http://" + publisher + ".frns.in/" + path)

	// Make Request
	resp, err := http.Get("http://" + publisher + ".frns.in/" + path)
	if err != nil {
		http.Error(w, err.Error(), resp.StatusCode)
		return
  }
	defer resp.Body.Close()

	// Proxy Headers
  for name, values := range resp.Header {
    // Loop over all values for the name.
    for _, value := range values {
			w.Header().Set(name, value)
    }
  }

	// Proxy Body
	io.Copy(w, resp.Body)
}
