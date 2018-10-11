// yml2props project main.go
package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"yml2props/properties"
	"yml2props/yml"
)

func currentPath(arg string) string {
	file, _ := exec.LookPath(arg)
	path, _ := filepath.Abs(file)
	return filepath.Dir(path)
}

func isExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

func isYml(path string) bool {
	var ext string
	ext = filepath.Ext(path)
	if ext == ".yml" || ext == ".yaml" {
		return true
	}
	return false
}

func getYamls(dirpath string) ([]string, error) {
	var dir_list []string
	dir_err := filepath.Walk(dirpath,
		func(path string, f os.FileInfo, err error) error {
			if f == nil {
				return err
			}
			if isYml(path) {
				dir_list = append(dir_list, path)
				return nil
			}
			return nil
		})

	return dir_list, dir_err
}

func makePropFilename(ymlPath string) string {
	ext := filepath.Ext(ymlPath)
	filename := strings.TrimSuffix(ymlPath, ext)
	return filename + ".properties"
}

func main() {
	path := currentPath(os.Args[0])
	cfgPath := flag.String("d", path+string(os.PathSeparator), "Enter a dir path")

	flag.Parse()

	if !isExist(*cfgPath) {
		fmt.Println("Directory is not existed:", *cfgPath)
		return
	}

	ymls, err := getYamls(*cfgPath)

	if len(ymls) == 0 {
		fmt.Println("No yml file found, exit")
	}

	if err != nil {
		log.Fatal(err)
		return
	}
	for _, yaml := range ymls {
		fmt.Printf("Read from %s\n", yaml)
		dict := yml.Read(yaml)
		propFile := makePropFilename(yaml)
		properties.Write(propFile, *dict)
	}
}
