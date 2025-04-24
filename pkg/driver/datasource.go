package driver

import (
	"errors"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

func MakeBaseDir(curDir string) (string, error) {
	baseDir := curDir
	dataSourceName := os.Getenv("TBLS_DSN")
	if len(dataSourceName) > 0 {
		pos := strings.Index(dataSourceName, ":")
		if pos == -1 {
			return "", errors.New("Bad data source name: " + dataSourceName)
		}
		dataSourceName = dataSourceName[pos+1:]
		if strings.HasPrefix(dataSourceName, "///") {
			if runtime.GOOS == "windows" {
				if dataSourceName[4] == ':' || strings.HasPrefix(dataSourceName, "/////") {
					// example: "///C:/Users", "/////Server/Users"
					baseDir = dataSourceName[3:]
				} else {
					// example: "///Users"
					baseDir = dataSourceName[2:]
				}
			} else {
				// example: "///usr"
				baseDir = dataSourceName[2:]
			}
			baseDir = filepath.FromSlash(baseDir)
		} else if strings.HasPrefix(dataSourceName, "./") || dataSourceName == "." {
			baseDir = filepath.FromSlash(dataSourceName)
		} else {
			return "", errors.New("Bad path: " + dataSourceName)
		}
	}

	return baseDir, nil
}
