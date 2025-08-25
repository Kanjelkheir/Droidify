package agent

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/kanjelkheir/droidify/internal/database"
)

type Data struct {
	Model               string `json:"model"`
	Twrp                string `json:"twrp"`
	Stock_firmware      string `json:"stock_firmware"`
	Odin                string `json:"odin"`
	Orange_fox_recovery string `json:"orange_fox_recovery"`
	Custom_firmware     string `json:"custom_firmware"`
	Shrp_recovery       string `json:"shrp_recovery"`
}

type RequestData struct {
	Model string `json:"model"`
}

// Configurator defines the interface for configuration objects that can get responses.
type Configurator interface {
	GetResponse() (string, error)
}

// newConfigFunc is a package-level variable to create Configurator instances, allowing for mocking in tests.
var newConfigFunc func(message, apiKey, endpoint string) Configurator = func(message, apiKey, endpoint string) Configurator {
	return &Config{
		message:  message,
		api_key:  apiKey,
		endpoint: endpoint,
	}
}

func GetData(r *http.Request, db *database.Queries) (Data, error) {
	api_key := os.Getenv("GEMINI_API_KEY")
	endpoint := "https://generativelanguage.googleapis.com/v1beta/models/gemini-2.5-flash:generateContent?key="
	var reqData RequestData
	var resultData Data

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&reqData); err != nil {
		return Data{}, errors.New("Invalid request body")
	}

	// check if device already exist in the data base (add caching)

	deviceData, err := db.GetDevice(context.Background(), reqData.Model)
	if err == nil {
		resultData.Model = deviceData.Model
		resultData.Custom_firmware = deviceData.CustomFirmeware.String
		resultData.Odin = deviceData.Odin.String
		resultData.Twrp = deviceData.Twrp.String
		resultData.Stock_firmware = deviceData.StockFirmware.String
		resultData.Orange_fox_recovery = deviceData.OrangeFoxRecovery.String

		return resultData, nil
	}

	prompt := fmt.Sprintf(`I want you to generate similar results for the following device-model: %v.
		Provide the output in JSON format following this structure:
		{
			"Model": "device model",
			"Twrp": "twrp link",
			"Stock_firmware": "stock firmware link",
			"Odin": "odin link",
			"Orange_fox_recovery": "orange fox recovery link",
			"Custom_firmware": "custom firmware link",
			"Shrp_recovery": "shrp recovery link"
		}
 		The device model is: %s`, reqData.Model, reqData.Model)

	configurator := newConfigFunc(prompt, api_key, endpoint)

	generated_response, err := configurator.GetResponse()
	if err != nil {
		return Data{}, err
	}

	// Check if the response is wrapped in markdown code block and extract the JSON
	if strings.HasPrefix(generated_response, "```json\n") && strings.HasSuffix(generated_response, "\n```") {
		generated_response = strings.TrimPrefix(generated_response, "```json\n")
		generated_response = strings.TrimSuffix(generated_response, "\n```")
	}

	bytes := []byte(generated_response)
	err = json.Unmarshal(bytes, &resultData)
	if err != nil {
		return Data{}, err
	}

	return resultData, nil

}

type Config struct {
	message  string
	api_key  string
	endpoint string
}

func (c *Config) GetResponse() (string, error) {
	if c.message == "" {
		return "", errors.New("Invalid Message")
	}

	url := c.endpoint + c.api_key

	requestBody := map[string]any{
		"contents": []map[string]any{
			{
				"parts": []map[string]string{
					{"text": c.message},
				},
			},
		},
		"generationConfig": map[string]any{
			"temperature": 0.7,
			"topK":        40,
			"topP":        0.95,
		},
	}

	// make json string from request body
	jsonData, err := json.Marshal(requestBody)

	if err != nil {
		return "", err
	}

	// get buffer from json
	buffer := bytes.NewBuffer(jsonData)
	req, err := http.NewRequest("post", url, buffer)
	if err != nil {
		return "", err
	}

	// set the data sent to be json data (change the headers)
	req.Header.Add("Content-Type", "application/json")

	// initialize a new client and send the request
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return "", err
	}

	defer res.Body.Close()

	var data GeminiResponse
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&data); err != nil {
		return "", err
	}

	result := data.Candidates[0].Content.Parts[0].Text
	return result, nil
}

type GeminiResponse struct {
	Candidates []Candidate `json:"candidates"`
}

type Candidate struct {
	Content       Content        `json:"content"`
	FinishReason  string         `json:"finishReason"`
	SafetyRatings []SafetyRating `json:"safetyRatings"`
}

type Content struct {
	Parts []Part `json:"parts"`
}

type Part struct {
	Text string `json:"text"`
}

type SafetyRating struct {
	Category    string `json:"category"`
	Probability string `json:"probability"`
}
