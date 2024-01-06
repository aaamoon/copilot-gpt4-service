package api

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"

	"github.com/gin-gonic/gin"

	"copilot-gpt4-service/config"
	"copilot-gpt4-service/utils"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	route := SetupRouter()
	route.ServeHTTP(w, r)
}

func SetupRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.Use(CORSMiddleware())
	r.POST("/v1/chat/completions", chatCompletions)
	r.GET("/v1/models", createMockModelsResponse)
	r.NoRoute(func(c *gin.Context) {
		c.String(http.StatusMethodNotAllowed, "Method Not Allowed")
	})
	return r
}

// Handle the Cross-Origin Resource Sharing (CORS) for requests.
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "*")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "*")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
			return
		}

		c.Next()
	}
}

// Represent the JSON data structure for the request body.
type JsonData struct {
	Messages    interface{} `json:"messages"`
	Model       string      `json:"model"`
	Temperature float64     `json:"temperature"`
	TopP        float64     `json:"top_p"`
	N           int64       `json:"n"`
	Stream      bool        `json:"stream"`
}

type Delta struct {
	Content string `json:"content"`
}

type Choice struct {
	Delta Delta `json:"delta"`
	Index int   `json:"index"`
}

type Data struct {
	Choices []Choice `json:"choices,omitempty"`
	Created int      `json:"created,omitempty"`
	ID      string   `json:"id,omitempty"`
}

// Generate a random hexadecimal string of the specified length.
func genHexStr(length int) string {
	bytes := make([]byte, length/2)
	if _, err := rand.Read(bytes); err != nil {
		panic(err)
	}
	return hex.EncodeToString(bytes)
}

// Create request headers to mock Github Copilot Chat requests.
func createHeaders(authorization string) map[string]string {
	headers := make(map[string]string, 0)
	headers["Authorization"] = "Bearer " + authorization
	headers["X-Request-Id"] = genHexStr(8) + "-" + genHexStr(4) + "-" + genHexStr(4) + "-" + genHexStr(4) + "-" + genHexStr(12)
	headers["Vscode-Sessionid"] = genHexStr(8) + "-" + genHexStr(4) + "-" + genHexStr(4) + "-" + genHexStr(4) + "-" + genHexStr(25)
	headers["Vscode-Machineid"] = genHexStr(64)
	headers["Editor-Version"] = "vscode/1.83.1"
	headers["Editor-Plugin-Version"] = "copilot-chat/0.8.0"
	headers["Openai-Organization"] = "github-copilot"
	headers["Openai-Intent"] = "conversation-panel"
	headers["Content-Type"] = "text/event-stream; charset=utf-8"
	headers["User-Agent"] = "GitHubCopilotChat/0.8.0"
	headers["Accept"] = "*/*"
	headers["Accept-Encoding"] = "gzip,deflate,br"
	headers["Connection"] = "close"

	return headers
}

func mainRequest(c *gin.Context) {
	content := c.Query("content")
	url := "https://api.githubcopilot.com/chat/completions"
	authorization := config.Authorization
	headers := createHeaders(authorization)
	jsonBody := &JsonData{
		Messages: []map[string]string{
			{"role": "system",
				"content": "\nYou are ChatGPT, a large language model trained by OpenAI.\nKnowledge cutoff: 2021-09\nCurrent model: gpt-4\n"},
			{"role": "user",
				"content": content},
		},
		Model:       "gpt-4",
		Temperature: 0.5,
		TopP:        1,
		N:           1,
		Stream:      false,
	}
	_ = c.BindJSON(&jsonBody)

	jsonData, err := json.Marshal(jsonBody)
	if err != nil {
		return
	}

	req, _ := http.NewRequest("POST", url, bytes.NewReader(jsonData))
	for k, v := range headers {
		req.Header.Set(k, v)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Encountering an error when sending the request.")
	} else {
		defer resp.Body.Close()
		if resp.StatusCode != http.StatusOK {
			return
		} else {
			// Set the headers for the response
			c.Writer.Header().Set("Transfer-Encoding", "chunked")
			c.Writer.Header().Set("X-Accel-Buffering", "no")
			c.Header("Content-Type", "text/event-stream; charset=utf-8")
			c.Header("Cache-Control", "no-cache")
			c.Header("Connection", "keep-alive")
			// Read the response body in chunks and write it to the response writer
			body := make([]byte, 64)
			for {
				n, err := resp.Body.Read(body)
				if err != nil && err != io.EOF {
					c.AbortWithError(http.StatusBadGateway, err)
					return
				}
				if n > 0 {
					rlt_body := body[:n]

					// Add the missing fields to the response body
					if bytes.Contains(rlt_body, []byte(`"choices":`)) && !bytes.Contains(rlt_body, []byte(`"object":`)) {
						rlt_body = bytes.ReplaceAll(rlt_body, []byte(`"choices":`), []byte(`"object": "chat.completion.chunk", "choices":`))
					}
					if bytes.Contains(rlt_body, []byte(`"choices":`)) && !bytes.Contains(rlt_body, []byte(`"model":`)) {
						rlt_body = bytes.ReplaceAll(rlt_body, []byte(`"choices":`), []byte(`"model": "gpt-4", "choices":`))
					}

					c.Writer.Write(rlt_body)
				}
				if err == io.EOF {
					break
				}
			}
		}
	}
}

func chatCompletions(c *gin.Context) {
	if !utils.GetAuthorization(c) {
		return
	}
	mainRequest(c)
}

func createMockModel(modelId string) gin.H {
	return gin.H{
		"id":       modelId,
		"object":   "model",
		"created":  1677610602,
		"owned_by": "openai",
		"permission": []gin.H{
			{
				"id":                   "modelperm-" + genHexStr(12),
				"object":               "model_permission",
				"created":              1677610602,
				"allow_create_engine":  false,
				"allow_sampling":       true,
				"allow_logprobs":       true,
				"allow_search_indices": false,
				"allow_view":           true,
				"allow_fine_tuning":    false,
				"organization":         "*",
				"group":                nil,
				"is_blocking":          false,
			},
		},
		"root":   modelId,
		"parent": nil,
	}
}

func createMockModelsResponse(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"object": "list",
		"data": []gin.H{
			createMockModel("gpt-3.5-turbo"),
			createMockModel("gpt-4"),
		},
	})
}
