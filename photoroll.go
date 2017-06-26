package main

import (
	"html/template"
	"image"
	_ "image/jpeg"
	"log"
	"os"
	"path/filepath"
)

type picture struct {
	Src     string
	Caption string
	Ratio   float64
}

func processTemplate(images []picture) {
	f, err := os.Create("index.html")
	t, err := template.ParseFiles("template.html")
	if err != nil {
		log.Print(err)
		return
	}

	if err != nil {
		log.Println("create file: ", err)
		return
	}

	context := map[string]interface{}{
		"Images": images,
		"Title":  "Photojournal",
	}
	t.Execute(f, context)
	f.Close()
}

func getImageDimension(imagePath string) (int, int) {
	file, err := os.Open(imagePath)
	if err != nil {
		print(os.Stderr, "%v\n", err)
	}

	image, _, err := image.DecodeConfig(file)
	if err != nil {
		print(os.Stderr, "%s: %v\n", imagePath, err)
	}
	return image.Width, image.Height
}

func main() {
	images := []picture{}
	files, err := filepath.Glob("images/*.jpg")
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		width, height := getImageDimension(file)
		ratio := float64(height) / float64(width) * 100
		images = append(images, picture{Src: file, Ratio: ratio})
		print("\n" + file)
	}

	processTemplate(images)
}
