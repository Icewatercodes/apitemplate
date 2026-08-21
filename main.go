package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"io"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

func main() {

	conffile, err := os.Open("config.json")
	if err != nil {
		log.Fatalf("failed to open: %v", err)
	}
	defer conffile.Close()

	var config struct {
		Ips  string `json:"ips"`
		Port uint16 `json:"port"`
	}

	if err := json.NewDecoder(conffile).Decode(&config); err != nil {
		log.Fatalf("decode json erro: %v", err)
	}

	fmt.Print(config)

	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", gettAlbumsByID)
	resp, err := http.Get("http://thinker:8080/albums/")
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("failed to read response body: %v", err)
	}

	fmt.Print("here:    ", resp, "\n")
	router.Run(fmt.Sprintf("%v:%v", config.Ips, config.Port))

}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func gettAlbumsByID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "album not found"})
}


