package main

import (
	"github.com/BurntSushi/toml"
	"github.com/anriclee/leego/http"
	"github.com/pkg/errors"
	"io/ioutil"
	"os"
	"path/filepath"
)

func main() {
	http.Serve()
}

func createNewSite(basePath string) error {
	// 检验路径是否存在
	fileInfo, err := os.Stat(basePath)
	if err != nil && !os.IsNotExist(err) {
		return err
	}
	if !fileInfo.IsDir() {
		return errors.New("target path should be a directory")
	}
	// 检验是否为空
	items, err := ioutil.ReadDir(basePath)
	if err != nil {
		return err
	}
	if len(items) != 0 {
		return errors.New("target directory should be empty")
	}
	archeTypePath := filepath.Join(basePath, "archetypes")
	dirs := []string{
		filepath.Join(basePath, "layouts"),
		filepath.Join(basePath, "content"),
		archeTypePath,
		filepath.Join(basePath, "static"),
		filepath.Join(basePath, "data"),
		filepath.Join(basePath, "themes"),
	}
	for _, dir := range dirs {
		if err := os.MkdirAll(dir, 0777); err != nil {
			return errors.Wrap(err, "Failed to create dir")
		}
	}
	configFile := filepath.Join(basePath, "config.toml")
	in := map[string]string{
		"baseURL":      "http://example.org/",
		"title":        "My New Hugo Site",
		"languageCode": "en-us",
	}
	file, err := os.Create(configFile)
	if err != nil {
		return err
	}
	err = toml.NewEncoder(file).Encode(in)
	if err != nil {
		return err
	}
	return nil
}
