package main

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"strings"
)

type Info struct {
	Files int
	Info  map[string]int
}

func (i Info) String() string {

	var result string

	for key, val := range i.Info {
		result += fmt.Sprintf("Extension : %s occurence %d\n", key, val)
	}

	return fmt.Sprintf("Number files : %d\n%s", i.Files, result)
}

type OrderFolder struct {
	Path string
	Info Info
	Dir  []string
}

func NewOrder(path string) *OrderFolder {

	order := OrderFolder{Path: path}

	order.infoFolder()

	return &order
}

func (n *OrderFolder) infoFolder() (*Info, error) {

	files, err := ioutil.ReadDir(n.Path)

	if err != nil {
		return nil, err
	}

	infos := make(map[string]int)

	for _, file := range files {

		if file.IsDir() {
			n.Dir = append(n.Dir, file.Name())
		} else {
			extension := getExtensionFile(file.Name())

			if _, ok := infos[extension]; ok {
				infos[extension] += 1
			} else {
				infos[extension] = 1
			}
		}

	}

	info := Info{Files: len(files), Info: infos}

	n.Info = info

	return &info, nil
}

func (n OrderFolder) listFiles() ([]fs.FileInfo, error) {
	files, err := ioutil.ReadDir(n.Path)

	if err != nil {
		return nil, err
	}

	return files, nil
}

func (n OrderFolder) moveFile(src, dest string) error {

	dest = fmt.Sprintf("%s/%s/%s", n.Path, dest, src)

	src = fmt.Sprintf("%s/%s", n.Path, src)

	err := os.Rename(src, dest)
	if err != nil {
		return err
	}

	return nil
}

func (n OrderFolder) tidyFolder() error {

	files, err := n.listFiles()

	if err != nil {
		return err
	}

	//Chaque fichier

	for _, file := range files {

		if !file.IsDir() {

			extension := strings.ToUpper(getExtensionFile(file.Name()))

			if !contains(n.Dir, extension) {

				path := fmt.Sprintf("%s/%s", n.Path, extension)

				err := os.MkdirAll(path, 0755)
				if err != nil {
					return err
				}
			}

			err = n.moveFile(file.Name(), extension)
			if err != nil {
				return err
			}

		}

	}

	return nil

}

func contains(array []string, value string) bool {
	for _, val := range array {
		if val == value {
			return true
		}
	}
	return false
}

func getExtensionFile(fileName string) string {
	splitName := strings.Split(fileName, ".")
	return splitName[len(splitName)-1]
}
