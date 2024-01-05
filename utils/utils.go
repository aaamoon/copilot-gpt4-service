package utils

import (
	"copilot-gpt4-service/config"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

var authorizations = make(map[string]Authorization)

type Authorization struct {
	Token     string `json:"token"`
	ExpiresAt int64  `json:"expires_at"`
}

// Set the Authorization in the cache.
func setAuthorizationToCache(copilotToken string, authorization Authorization) {
	authorizations[copilotToken] = authorization
}

// Obtain the Authorization from the cache.
func getAuthorizationFromCache(copilotToken string) *Authorization {
	extraTime := rand.Intn(600) + 300
	if authorization, ok := authorizations[copilotToken]; ok {
		if authorization.ExpiresAt > time.Now().Unix()+int64(extraTime) {
			return &authorization
		}
	}
	return &Authorization{}
}

// When obtaining the Authorization, first attempt to retrieve it from the cache. If it is not available in the cache, retrieve it through an HTTP request and then set it in the cache.
func getAuthorizationFromToken(c *gin.Context, copilotToken string) {
	authorization := getAuthorizationFromCache(copilotToken)
	if authorization.Token == "" {
		getAuthorizationUrl := "https://api.github.com/copilot_internal/v2/token"
		client := &http.Client{}
		req, _ := http.NewRequest("GET", getAuthorizationUrl, nil)
		req.Header.Set("Authorization", "token "+copilotToken)
		response, err := client.Do(req)
		if err != nil || response.StatusCode != 200 {
			c.JSON(response.StatusCode, gin.H{"error": err.Error()})
		}
		defer response.Body.Close()

		body, _ := ioutil.ReadAll(response.Body)

		newAuthorization := &Authorization{}
		if err = json.Unmarshal(body, &newAuthorization); err != nil {
			fmt.Println("err", err)
		}
		authorization.Token = newAuthorization.Token
		setAuthorizationToCache(copilotToken, *newAuthorization)
	}
	config.Authorization = authorization.Token
}

// Retrieve the GitHub Copilot Plugin Token from the request header.
func GetAuthorization(c *gin.Context) {
	copilotToken := strings.TrimPrefix(c.GetHeader("Authorization"), "Bearer ")
	if copilotToken == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
		return
	}
	// Obtain the Authorization from the GitHub Copilot Plugin Token.
	getAuthorizationFromToken(c, copilotToken)
}
