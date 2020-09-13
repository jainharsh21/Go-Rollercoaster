package main
import (
	"net/http"
)

// Coaster type
type Coaster struct {
	Name         string `json:"name"`
	Manufacturer string `json:"manufacturer"`
	ID           string `json:"id"`
	InPark       string `json:"in_park"`
	Height       int    `json:"height"`
}



func coastersHandler(w http.ResponseWriter, r *http.Request){

}

func main()  {
	http.HandleFunc("/coasters",coastersHandler)
	err := http.ListenAndServe(":8080",nil)
	if err != nil {
		panic(err)
	}
}