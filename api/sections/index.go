package handler

import (
	"encoding/json"
	"log"
	"strconv"
	"net/http"
)

var Publishers = map[string]string{
  "cap": "capvideo",
  "hiz-ve-renk" : "hizrenkvideo",
  "fen-bilimleri" : "fenbilimlerivideocozum",
  "apotemi" : "apotemivideo",
}

type Section struct {
	Id        string  `json:"_id"`
	Title     string  `json:"title"`
	Parent_id string  `json:"parent_id"`
  Is_parent bool    `json:"is_parent"` // false: it is not a parent to another section, it contains just solutions 
}

func Handler(w http.ResponseWriter, r *http.Request) {
  log.Println("[REQUEST]")

	publisher := Publishers[r.URL.Query().Get("publisher")]

	// Make Request
	resp, err := http.Get("http://" + publisher + ".frns.in/controller/get-source.php?action=source_list&id=" + r.URL.Query().Get("parent_id"))
	if err != nil {
		http.Error(w, err.Error(), resp.StatusCode)
		return
  }
	defer resp.Body.Close()

	// Parse Request
	var body map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&body); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
  }
	var sections = make([]Section, len(body["sources"].([]interface{})))
	for i := 0; i < len(sections); i++ {
		// id -> _id ; pid -> parent_id ; parent -> is_parent etc.
		sect := body["sources"].([]interface{})[i].(map[string]interface{})
		sections[i].Id = sect["id"].(string)
		sections[i].Parent_id = sect["pid"].(string)
		sections[i].Is_parent, _ = strconv.ParseBool((sect["parent"].(string)))
		sections[i].Title = sect["nm"].(string)
	}

	w.Header().Set("Content-Type", "application/json")
  json.NewEncoder(w).Encode(sections)
}
