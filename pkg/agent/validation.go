package agent

import (
	"path/filepath"

	"github.com/nSakkriou/utils/pkg/logn"
	"github.com/nSakkriou/utils/pkg/util"
)

// Checker le fichier de config
// Tous les champs rempli
// Au moins 1 script dans la liste
// Pas de doublons dans les endpoints
// checker si les fichiers existes
func ValidConfig(config *AgentFile) bool {
	if config.Port == 0 || config.Port == -1 {
		logn.Error("invalid port number")
		return false
	}

	if util.IsEmpty(config.ScriptsFolderPath) {
		logn.Error("scripts_folder_path field is empty")
		return false
	}

	// Check si le dossier existe
	if util.FileExist(config.ScriptsFolderPath) {
		logn.Error("folder %s doesn't exist", config.ScriptsFolderPath)
		return false
	}

	endpointNames := []string{}
	for _, endCommand := range config.EndCommands {
		// Checker les doublons
		if util.ArrayContains(endpointNames, endCommand.EndpointName) {
			logn.Error("endpoint %s cannot be defined twice", endCommand.EndpointName)
			return false
		}

		endpointNames = append(endpointNames, endCommand.EndpointName)

		// Checker si les fichiers existe
		for _, scriptName := range endCommand.ScriptsFilesNames {
			path := filepath.Join(config.ScriptsFolderPath, scriptName)
			if util.FileExist(path) {
				logn.Error("script %s doesn't exist", path)
				return false
			}
		}
	}

	return true
}
