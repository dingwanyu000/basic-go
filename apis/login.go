package apis

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	. "github.com/testGo/models"
	"io/ioutil"
	"net/http"
)

func LoginApi(c *gin.Context) {
	body, _ := ioutil.ReadAll(c.Request.Body)
	var dat map[string]interface{}
	if err := json.Unmarshal([]byte(string(body)), &dat); err == nil {
	}
	l := Login{PassWord: dat["password"].(string), UserName: dat["username"].(string)}
	state := l.Login(l)
	c.JSON(http.StatusOK, gin.H{
		"state": state,
	})

}
