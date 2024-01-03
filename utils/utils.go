package utils

import (
  "encoding/json"
  "fmt"
  "io/ioutil"
  "net/http"
  "copilot-gpt4-service/config"
  "math/rand"
  "strings"
  "time"

  "github.com/gin-gonic/gin"
)

var tokens = make(map[string]Token)

type Token struct {
  Token     string `json:"token"`
  ExpiresAt int64  `json:"expires_at"`
}

// Setting the Copilot token in the cache.
func setTokenToCache(githubToken string, copilotToken Token) {
  tokens[githubToken] = copilotToken
}

// Obtaining the Copilot token from the cache.
func getTokenFromCache(githubToken string) *Token {
  extraTime := rand.Intn(600) + 300
  if token, ok := tokens[githubToken]; ok {
      if token.ExpiresAt > time.Now().Unix()+int64(extraTime) {
          return &token
      }
  }
  return &Token{}
}

// When obtaining the Copilot token, first attempt to retrieve it from the cache. If it is not available in the cache, retrieve it through an HTTP request and then set it in the cache.
func getCopilotToken(c *gin.Context, githubToken string) {
  token := getTokenFromCache(githubToken)
  if token.Token == "" {
    getTokenUrl := "https://api.github.com/copilot_internal/v2/token"
    client := &http.Client{}
    req, _ := http.NewRequest("GET", getTokenUrl, nil)
    req.Header.Set("Authorization", "token "+githubToken)
    response, err := client.Do(req)
    if err != nil || response.StatusCode != 200 {
      c.JSON(response.StatusCode, gin.H{"error": err.Error()})
    }
    defer response.Body.Close()

    body, _ := ioutil.ReadAll(response.Body)

    copilotToken := &Token{}
    if err = json.Unmarshal(body, &copilotToken); err != nil {
      fmt.Println("err", err)
    }
    token.Token = copilotToken.Token
    setTokenToCache(githubToken, *copilotToken)
  }
  config.CoToken = token.Token
}

// Retrieve the GitHub token from the request header.
func GetGithubTokens(c *gin.Context) {
  githubToken := strings.TrimPrefix(c.GetHeader("Authorization"), "Bearer ")
  if githubToken == "" {
    c.JSON(http.StatusUnauthorized, gin.H{
      "message": "Unauthorized",
    })
    return
  }
  // Obtain the Copilot token from the GitHub token.
  getCopilotToken(c, githubToken)
}
