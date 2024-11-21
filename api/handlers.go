package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Handler for Time API
func (app *Config) GetTime(w http.ResponseWriter, r *http.Request) {
	// Fetch time zone from query parameters
	timeZone := r.URL.Query().Get("timeZone")
	if timeZone == "" {
		app.errorJSON(w, fmt.Errorf("missing 'timeZone' parameter"))
		return
	}

	// Construct API URL
	apiURL := fmt.Sprintf("https://timeapi.io/api/time/current/zone?timeZone=%s", timeZone)

	// Make the API request
	resp, err := http.Get(apiURL)
	if err != nil || resp.StatusCode != http.StatusOK {
		app.errorJSON(w, fmt.Errorf("failed to fetch time: %v", err))
		return
	}
	defer resp.Body.Close()

	// Decode response
	var timeData map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&timeData)
	if err != nil {
		app.errorJSON(w, fmt.Errorf("error decoding response: %v", err))
		return
	}

	// Return the data as JSON
	app.writeJSON(w, http.StatusOK, timeData)
}

// Placeholder handler for a second webhook
func (app *Config) PlaceholderWebhook(w http.ResponseWriter, r *http.Request) {
	var requestPayload struct {
		Data string `json:"data"`
	}

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	response := map[string]string{
		"message": "This is a placeholder webhook",
		"data":    requestPayload.Data,
	}

	app.writeJSON(w, http.StatusOK, response)
}
