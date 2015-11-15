package env

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

func GetHome() string {
	nv_home := os.Getenv("NV_HOME")
	if nv_home != "" {
		nv_home = filepath.Join(os.Getenv("HOME"), ".envirius")
	}
	return nv_home
}

// Get root directory for all environments
func GetEnvHome() string {
	return filepath.Join(GetHome(), "envs")
}

// Get list of existing environments
func GetEnvList() ([]Environment, error) {
	var envs []Environment

	files, err := ioutil.ReadDir(GetEnvHome())
	if err != nil {
		return nil, err
	}

	for _, f := range files {
		envs = append(envs, *NewEnv(f.Name()))
	}
	return envs, nil
}
