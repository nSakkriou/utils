package agent

func GetConfig() *AgentFile {
	if isLoad {
		return agentFile
	}

	panic("impossible de récupérer la config")
}
