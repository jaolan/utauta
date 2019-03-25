package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/gorilla/handlers"
	"log"
	"math/rand"
	"net/http"
	"strconv"
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
	AlbumID	string      `json:"album_id"`
	Name	string 	    `json:"name"`
	Year 	int   		`json:"year"`
	Picture	string 	    `json:picture`
	Artist		   	    `json:artist_id`
}

//init artist var as slice artist struct
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


//Album CRUD operations
//init album var as slice album struct
var albums []Album

//get all albums
func  getAlbums(w http.ResponseWriter, r *http.Request)  {
	//setting header of content to type json instead of plaintext
	w.Header().Set("Content-Type", "application/json")
	//sending the test struct to json encoder
	json.NewEncoder(w).Encode(albums)
}

//get one album
func  getAlbum(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	//get the params
	params := mux.Vars(r)
	for _, item := range albums{
		if item.AlbumID == params["album_id"]{
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Album{})
}

//create one album
func  createAlbums(w http.ResponseWriter, r *http.Request)  {
	//setting header of content to type json instead of plaintext
	w.Header().Set("Content-Type", "application/json")
	var album Album
	_ = json.NewDecoder(r.Body).Decode(&album)
	//@todo for example ONLY, MOCK ID
	album.AlbumID = strconv.Itoa(rand.Intn(30))
	albums = append(albums,album)
	json.NewEncoder(w).Encode(album)
}

//update an album
func  updateAlbum(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range albums{
		//checks for album_id, uses slices to remove an album via append
		if item.AlbumID == params["album_id"] {
			albums = append(albums[:index],  albums[index+1:]...)
			var album Album
			_ = json.NewDecoder(r.Body).Decode(&album)
			//@todo for example ONLY, MOCK ID
			album.AlbumID = params["album_id"]
			albums = append(albums,album)
			json.NewEncoder(w).Encode(album)
			return
		}
	}
	json.NewEncoder(w).Encode(albums)
}

//remove an album
func  deleteAlbum(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range albums{
		//checks for album_id, uses slices to remove an album via append
		if item.AlbumID == params["album_id"] {
			albums = append(albums[:index],  albums[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(albums)
}

func main(){
	//init mux router
	r:= mux.NewRouter()

	//testing data @todo implem db, for album add Album: &Album
	artists = append(artists, Artist{ArtistID: "1", Name:"frank", Desc:"frank makes jpop.", Picture:"http://www.trueactivist.com/wp-content/uploads/2015/07/hotdog-1024x768.jpg"})
	artists = append(artists, Artist{ArtistID: "2", Name:"まゆこ", Desc:"J-POPが好きです", Picture:"https://aprilrose0404.files.wordpress.com/2010/10/onigiri-1.jpg"})

	//testing data for Albums
	albums = append(albums, Album{AlbumID: "1", Name:"franks album", Year:1982, Picture:"http://www.trueactivist.com/wp-content/uploads/2015/07/hotdog-1024x768.jpg"})
	albums = append(albums, Album{AlbumID: "2", Name:"まゆこ", Year:1990, Picture:"https://aprilrose0404.files.wordpress.com/2010/10/onigiri-1.jpg"})

	// Route handlers / Endpts
	r.HandleFunc("/artists", getArtists).Methods("GET")
	r.HandleFunc("/artists/{artist_id}", getArtist).Methods("GET")
	r.HandleFunc("/artists", createArtists).Methods("POST")
	r.HandleFunc("/artists/{artist_id}", updateArtist).Methods("PUT")
	r.HandleFunc("/artists/{artist_id}", deleteArtist).Methods("DELETE")

	// Route handlers / Endpts for Albums
	r.HandleFunc("/albums", getAlbums).Methods("GET")
	r.HandleFunc("/albums/{album_id}", getAlbum).Methods("GET")
	r.HandleFunc("/albums", createAlbums).Methods("POST")
	r.HandleFunc("/albums/{album_id}", updateAlbum).Methods("PUT")
	r.HandleFunc("/albums/{album_id}", deleteAlbum).Methods("DELETE")

	//todo implement matching with artist foreign key
	//Route handlers / Endpts for an artists' albums
	r.HandleFunc("/artists/{artist_id}/albums", getAlbums).Methods("GET")
	r.HandleFunc("/artists/{artist_id}/{album_id}", getAlbum).Methods("GET")
	r.HandleFunc("/artists/{artist_id}/albums", createAlbums).Methods("POST")
	r.HandleFunc("/artists/{artist_id}/{album_id}", updateAlbum).Methods("PUT")
	r.HandleFunc("/artists/{artist_id}/{album_id}", deleteAlbum).Methods("DELETE")

	//listens on port, log.fatal throws err if something fails
	log.Fatal(http.ListenAndServe(":9000",  handlers.CORS()(r)))

}