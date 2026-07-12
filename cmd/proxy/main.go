package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"strings"
)

var (
	openAIKeyRegex = regexp.MustCompile(`sk-[a-zA-Z0-9]{48}`)
	hfTokenRegex   = regexp.MustCompile(`hf_[a-zA-Z0-9]{34}`)
	injectionTerms = []string{"ignore previous instructions", "system prompt", "override settings", "you must follow"}
)

type LLMRequest struct {
	Prompt string `json:"prompt"`
}

type LLMResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Prompt  string `json:"processed_prompt,omitempty"`
}

func sanitizePrompt(prompt string) (string, bool) {
	lowered := strings.ToLower(prompt)
	for _, term := range injectionTerms {
		if strings.Contains(lowered, term) {
			return "", true
		}
	}

	sanitized := openAIKeyRegex.ReplaceAllString(prompt, "[REDACTED_OPENAI_KEY]")
	sanitized = hfTokenRegex.ReplaceAllString(sanitized, "[REDACTED_HF_TOKEN]")

	return sanitized, false
}

func handleProxy(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	var req LLMRequest
	if err := json.Unmarshal(body, &req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(LLMResponse{Status: "error", Message: "Invalid JSON structure"})
		return
	}

	sanitizedPrompt, isInjected := sanitizePrompt(req.Prompt)
	w.Header().Set("Content-Type", "application/json")

	if isInjected {
		w.WriteHeader(http.StatusForbidden)
		json.NewEncoder(w).Encode(LLMResponse{Status: "blocked", Message: "Prompt injection attempt detected"})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(LLMResponse{
		Status:  "success",
		Message: "Prompt passed validation",
		Prompt:  sanitizedPrompt,
	})
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "7860"
	}

	http.HandleFunc("/v1/sanitize", handleProxy)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "eBPF Contextual LLM Firewall Gateway Engine Running. Send POST to /v1/sanitize")
	})

	if err := http.ListenAndServe(":"+port, nil); err != nil {
		os.Exit(1)
	}
}
