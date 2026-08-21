package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

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
	router.GET("/:url/:path", gethttp)
	router.Run(fmt.Sprintf("%v:%v", config.Ips, config.Port))

}

func gethttp(c *gin.Context) {
	url := c.Param("url")
	path := c.Param("path")

	fmt.Print(fmt.Sprintf("\nhttp://%v/%v\n", url, path))

	resp, err := http.Get(fmt.Sprintf("http://%v/%v", url, path))
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	defer resp.Body.Close()

	/*body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("failed to read response body: %v", err)
	}

	var output map[string]any

	err = json.Unmarshal(body, &output)
	if err != nil {
		fmt.Print(err)
	}*/

	c.JSON(http.StatusOK, resp)
}


