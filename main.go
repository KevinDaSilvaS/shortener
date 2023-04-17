package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"shortener/app/links"
	"shortener/customtypes"
	"shortener/repository"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	dbConn := repository.Conn()
	router.LoadHTMLGlob("templates/*")

	router.POST("/s", func(c *gin.Context) {
		newlink := customtypes.NewLink{Url: "google.com", Alias: "g"}
		rawBody, _ := ioutil.ReadAll(c.Request.Body)

		json.Unmarshal([]byte(rawBody), &newlink)

		code, err := links.AddLink(newlink, dbConn)

		out := gin.H{
			"error": fmt.Sprint(err),
		}

		if err == nil {
			out = gin.H{
				"alias": newlink.Alias,
				"path":  "/s/{" + newlink.Alias + "}",
			}
		}

		c.JSON(code, out)
	})

	router.GET("/s/:linkalias", func(c *gin.Context) {
		alias := c.Param("linkalias")

		code, url, err := links.GetLink(alias, dbConn)

		if err != nil {
			c.JSON(code, gin.H{
				"error": fmt.Sprint(err),
			})

			return
		}

		c.HTML(code, "link.tmpl", gin.H{"url": url})
	})

	http.ListenAndServe(":4578", router)
}
