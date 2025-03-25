package fileutil

import (
	"fmt"
	"os"
	"image"
	"image/png"
)


func SaveImage(img image.Image, filename string) error{
	file, err := os.Create(filename)
	if err!=nil {
		fmt.Println(err)
		return err
	}
	defer file.Close()

	err = png.Encode(file, img)
	if err!=nil {
		fmt.Println("Error encoding image", err)
		return err
	}

	fmt.Println("Image generated successfully")
	return nil
}

