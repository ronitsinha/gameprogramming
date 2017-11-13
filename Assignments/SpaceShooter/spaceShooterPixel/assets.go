package main

import (
	"image"
	"os"
	"io/ioutil"
	"path/filepath"

	_ "image/png"

	"fmt"

	pixel "github.com/faiface/pixel"
)

type Resources struct {
	images map[string]pixel.Picture
}

func NewResources() *Resources {
	r := new(Resources)

	r.images = make(map[string]pixel.Picture)

	r.LoadAllPictures ("./assets/images")

	return r
}

func LoadPicture (path string) (pixel.Picture, error) {
	file, err := os.Open (path)
	if err != nil {
		return nil, err
	}
	defer file.Close ()
	img, _, err := image.Decode (file)
	if err != nil {
		return nil, err
	} 

	return pixel.PictureDataFromImage (img), nil
}

func (r *Resources) LoadAllPictures(dir string) {
	files, err := ioutil.ReadDir(dir)

	if err != nil {
		fmt.Println(err)
		return
	}

	for _, f := range files {
		if f.IsDir() {
			r.LoadAllPictures(dir + "/" + f.Name())
		} else if filepath.Ext(f.Name()) == ".png" {
			pic, er := LoadPicture (dir + "/" + f.Name ())
			if er != nil {
				fmt.Println (er)
				return
			}

			r.images[f.Name()] = pic
		}
	}
}