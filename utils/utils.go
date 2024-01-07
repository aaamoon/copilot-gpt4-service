package utils

import (
	"copilot-gpt4-service/config"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

var authorizations = make(map[string]Authorization)

type Authorization struct {
	Token     string `json:"token"`
	ExpiresAt int64  `json:"expires_at"`
}

var defaultCopilotToken = os.Getenv("DEFAULT_COPILOT_TOKEN")
var superToken = os.Getenv("SUPER_TOKEN")

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
func getAuthorizationFromToken(c *gin.Context, copilotToken string) bool {
	authorization := getAuthorizationFromCache(copilotToken)
	if authorization == nil || authorization.Token == "" {
		getAuthorizationUrl := "https://api.github.com/copilot_internal/v2/token"
		client := &http.Client{}
		req, _ := http.NewRequest("GET", getAuthorizationUrl, nil)
		req.Header.Set("Authorization", "token "+copilotToken)
		response, err := client.Do(req)
		if err != nil {
			if response == nil {
				// handle connection not available
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return false
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "code": response.StatusCode})
			return false
		}
		if response.StatusCode != 200 {
			body, _ := io.ReadAll(response.Body)
			c.JSON(response.StatusCode, gin.H{"error": string(body), "code": response.StatusCode})
			return false
		}
		defer response.Body.Close()

		body, _ := io.ReadAll(response.Body)

		newAuthorization := &Authorization{}
		if err = json.Unmarshal(body, &newAuthorization); err != nil {
			fmt.Println("err", err)
		}
		authorization.Token = newAuthorization.Token
		setAuthorizationToCache(copilotToken, *newAuthorization)
	}
	config.Authorization = authorization.Token
	return true
}

// Retrieve the GitHub Copilot Plugin Token from the request header.
func GetAuthorization(c *gin.Context) bool {
	copilotToken := strings.TrimPrefix(c.GetHeader("Authorization"), "Bearer ")
	if copilotToken == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
		return false
	}

	// Use the default token if the user's token is the SUPER_TOKEN.
	if superToken != "" && superToken != "your_super_token" && copilotToken == superToken {
		println("Found SUPER_TOKEN, using default token for this request.")
		copilotToken = defaultCopilotToken
	}

	// Obtain the Authorization from the GitHub Copilot Plugin Token.
	return getAuthorizationFromToken(c, copilotToken)
}
