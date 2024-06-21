package agent

type AgentFile struct {
	Port              int          `json:"port"`
	ScriptsFolderPath string       `json:"scripts_folder_path"`
	EndCommands       []EndCommand `json:"endcommands"`
}

type EndCommand struct {
	EndpointName      string   `json:"endpoint_name"`
	Method            string   `json:"method"`
	ScriptsFilesNames []string `json:"scripts_files_names"`
}
