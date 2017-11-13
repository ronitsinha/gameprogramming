package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"

	sf "github.com/zyedidia/sfml/v2.3/sfml"
)

type Resources struct {
	images map[string]*sf.Texture
}

func NewResources() *Resources {
	r := new(Resources)

	r.images = make(map[string]*sf.Texture)

	r.LoadImages("./assets/images")

	return r
}

func (r *Resources) LoadImages(dir string) {
	files, err := ioutil.ReadDir(dir)

	if err != nil {
		fmt.Println(err)
		return
	}

	for _, f := range files {
		if f.IsDir() {
			r.LoadImages(dir + "/" + f.Name())
		} else if filepath.Ext(f.Name()) == ".png" {
			texture := sf.NewTexture(dir + "/" + f.Name())
			r.images[f.Name()] = texture
		}
	}
}
