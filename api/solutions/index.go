package handler

import (
	"encoding/json"
	"log"
	"net/http"
)

var Publishers = map[string]string{
  "cap": "capvideo",
  "hiz-ve-renk" : "hizrenkvideo",
  "fen-bilimleri" : "fenbilimlerivideocozum",
  "apotemi" : "apotemivideo",
}

type Solution struct {
	Id            string  `json:"_id"`
	Title         string  `json:"title"`
	Solution_id   string  `json:"solution_id"`
	Solution_type string  `json:"solution_type"`
	Audio         string  `json:"audio"`
	SWF           string  `json:"swf"`
	XAML          string  `json:"xaml"`
}

func Handler(w http.ResponseWriter, r *http.Request) {
  log.Println("[REQUEST]")

	publisher := Publishers[r.URL.Query().Get("publisher")]

	// Make Request
	resp, err := http.Get("http://" + publisher + ".frns.in/controller/get-source.php?action=content_list&id=" + r.URL.Query().Get("parent_id"))
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
	var solutions = make([]Solution, len(body["contents"].([]interface{})))
	for i := 0; i < len(solutions); i++ {
		// id -> _id ; pid -> parent_id ; parent -> is_parent etc.
		solution := body["contents"].([]interface{})[i].(map[string]interface{})
		solutions[i].Id = solution["id"].(string)
		solutions[i].Title = solution["nm"].(string)
		solutions[i].Solution_id = solution["solved_id"].(string)
		solutions[i].Solution_type = solution["solved_type"].(string)
		solutions[i].Audio = solution["audio"].(string)
		solutions[i].SWF = solution["swf"].(string)
		solutions[i].XAML = solution["video"].(string)
	}

	w.Header().Set("Content-Type", "application/json")
  json.NewEncoder(w).Encode(solutions)
}
