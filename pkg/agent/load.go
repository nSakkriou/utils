package agent

import (
	"encoding/json"
	"log"
	"os"
)

var agentFile *AgentFile
var isLoad = false

// Charger la config
func Load() error {
	byteValue, err := os.ReadFile(AgentFileName)
	if err != nil {
		log.Fatalf("Failed to read file: %s", err)
		return err
	}

	if err := json.Unmarshal(byteValue, &agentFile); err != nil {
		log.Fatalf("Failed to parse JSON: %s", err)
	}

	isLoad = true
	return nil
}

func IsLoad() bool {
	return isLoad
}
