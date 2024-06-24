package agent

import (
	"encoding/json"
	"os"

	"github.com/nSakkriou/utils/pkg/logn"
)

// Charger la config
func Load(path string) (*AgentFile, error) {
	var agentFile *AgentFile

	byteValue, err := os.ReadFile(path)
	if err != nil {
		logn.Fatal("Failed to read file: %s", err)
		return nil, err
	}

	if err := json.Unmarshal(byteValue, &agentFile); err != nil {
		logn.Fatal("Failed to parse JSON: %s", err)
		return nil, err
	}

	return agentFile, nil
}
