package apis

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	. "github.com/testGo/models"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

func IndexApi(c *gin.Context) {
	c.String(http.StatusOK, "It works")
}

func AddPersonApi(c *gin.Context) {
	body, _ := ioutil.ReadAll(c.Request.Body)
	var dat map[string]interface{}
	if err := json.Unmarshal([]byte(string(body)), &dat); err == nil {
	}
	p := Person{UserName: dat["username"].(string), PassWord: dat["password"].(string)}
	ra, err := p.AddPerson()
	if err != nil {
		log.Fatalln(err)
	}
	msg := fmt.Sprintf("insert successful %d", ra)
	c.JSON(http.StatusOK, gin.H{
		"msg": msg,
	})
}

func GetPersonsApi(c *gin.Context) {
	var p Person
	persons, err := p.GetPersons()
	if err != nil {
		log.Fatalln(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"persons": persons,
	})

}

func GetPersonApi(c *gin.Context) {
	cid := c.Param("id")
	id, err := strconv.Atoi(cid)
	if err != nil {
		log.Fatalln(err)
	}
	p := Person{Id: id}
	person, err := p.GetPerson()
	if err != nil {
		log.Fatalln(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"person": person,
	})

}

func ModPersonApi(c *gin.Context) {
	body, _ := ioutil.ReadAll(c.Request.Body)
	var dat map[string]interface{}
	if err := json.Unmarshal([]byte(string(body)), &dat); err == nil {
	}
	p := Person{Id: int(dat["id"].(float64)), UserName: dat["username"].(string), PassWord: dat["password"].(string)}
	ra, err := p.ModPerson()
	if err != nil {
		log.Fatalln(err)
	}
	msg := fmt.Sprintf("insert successful %d", ra)
	c.JSON(http.StatusOK, gin.H{
		"msg": msg,
	})
}

func DelPersonApi(c *gin.Context) {
	body, _ := ioutil.ReadAll(c.Request.Body)
	var dat map[string]interface{}
	if err := json.Unmarshal([]byte(string(body)), &dat); err == nil {
	}
	p := Person{Id: int(dat["id"].(float64))}
	ra, err := p.DelPerson()
	if err != nil {
		log.Fatalln(err)
	}
	msg := fmt.Sprintf("Delete person %d successful %d", int(dat["id"].(float64)), ra)
	c.JSON(http.StatusOK, gin.H{
		"msg": msg,
	})
}
