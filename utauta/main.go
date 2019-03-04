package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

//Artist model
type Artist struct {
	ArtistID string		`json:"artist_id"`
	Name	 string 	`json:"name"`
	Desc	 string 	`json:"description"`
	Picture	 string 	`json:picture`
}

//Album model
type Album struct {
	AlbumID	string         `json:"album_id"`
	Name	string 	    `json:"name"`
	Year 	time.Time   `json:"description"`
	Picture	string 	    `json:picture`
	Artist		   	    `json:artist_id`
}

//init books var as slice artist struct
var artists []Artist

//get all artists
func  getArtists(w http.ResponseWriter, r *http.Request)  {
	//setting header of content to type json instead of plaintext
	w.Header().Set("Content-Type", "application/json")
	//sending the test struct to json encoder
	json.NewEncoder(w).Encode(artists)
}

//get one artist
func  getArtist(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	//get the params
	params := mux.Vars(r)
	for _, item := range artists{
		if item.ArtistID == params["artist_id"]{
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Artist{})
}

//create one artist
func  createArtists(w http.ResponseWriter, r *http.Request)  {
	//setting header of content to type json instead of plaintext
	w.Header().Set("Content-Type", "application/json")
	var artist Artist
	_ = json.NewDecoder(r.Body).Decode(&artist)
	//@todo for example ONLY, MOCK ID
	artist.ArtistID = strconv.Itoa(rand.Intn(30))
	artists = append(artists,artist)
	json.NewEncoder(w).Encode(artist)
}

//update an artist
func  updateArtist(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range artists{
		//checks for artist_id, uses slices to remove an artist via append
		if item.ArtistID == params["artist_id"] {
			artists = append(artists[:index],  artists[index+1:]...)
			var artist Artist
			_ = json.NewDecoder(r.Body).Decode(&artist)
			//@todo for example ONLY, MOCK ID
			artist.ArtistID = params["artist_id"]
			artists = append(artists,artist)
			json.NewEncoder(w).Encode(artist)
			return
		}
	}
	json.NewEncoder(w).Encode(artists)
}

//remove an artist
func  deleteArtist(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range artists{
		//checks for artist_id, uses slices to remove an artist via append
		if item.ArtistID == params["artist_id"] {
			artists = append(artists[:index],  artists[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(artists)
}


func main(){
	//init mux router
	r:= mux.NewRouter()

	//testing data @todo implem db, for album add Album: &Album
	artists = append(artists, Artist{ArtistID: "1", Name:"frank", Desc:"frank makes jpop.", Picture:"http://www.trueactivist.com/wp-content/uploads/2015/07/hotdog-1024x768.jpg"})
	artists = append(artists, Artist{ArtistID: "2", Name:"まゆこ", Desc:"J-POPが好きです", Picture:"https://aprilrose0404.files.wordpress.com/2010/10/onigiri-1.jpg"})


	// Route handlers / Endpts
	r.HandleFunc("/artists", getArtists).Methods("GET")
	r.HandleFunc("/artists/{artist_id}", getArtist).Methods("GET")
	r.HandleFunc("/artists", createArtists).Methods("POST")
	r.HandleFunc("/artists/{artist_id}", updateArtist).Methods("PUT")
	r.HandleFunc("/artists/{artist_id}", deleteArtist).Methods("DELETE")

	//listens on port, log.fatal throws err if something fails
	log.Fatal(http.ListenAndServe(":8000", r))


}