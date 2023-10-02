package main

/**
 * crf2html
 *
 * This program generates an HTML page displaying image textures from a given directory or CRF/ZIP file.
 * It resizes and encodes the images as base64 and creates an organized HTML page.
 *
 * Usage: go build -o crf2html main.go && ./crf2html source_path output_path [-title "Page Title"]
 * Example: go build -o crf2html main.go && ./crf2html ./fam.crf ./textures.html -title "My Custom Title"
 *
 * Arguments:
 *  - source_path: Path to the directory containing image files or a CRF/ZIP file.
 *  - output_path: Path to the HTML file to be generated.
 *
 * Options:
 *  -title: (Optional) Custom title for the HTML page. If not provided, the default title is "Textures".
 *  -size: (Optional) Custom thumbnail size for the HTML page. If not provided, the default size is "128".
 */

import (
	"archive/zip"
	"bytes"
	"encoding/base64"
	"fmt"
	"html"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	"github.com/ftrvxmtrx/tga"
	"github.com/nfnt/resize"
	"github.com/samuel/go-pcx/pcx"
)

type ProgramSettings struct {
	SourcePath      string
	OutputPath      string
	PageTitle       string
	ThumbnailSize   int
	BackgroundColor color.RGBA
}

func FileListing(directoryPath string) ([]string, error) {
	var files []string

	err := filepath.Walk(directoryPath, func(filePath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			files = append(files, filePath)
		}

		return nil
	})

	return files, err
}

func GetImageFromZip(zipReader *zip.ReadCloser, filePath string) (image.Image, error) {
	for _, file := range zipReader.File {
		if file.Name == filePath {
			reader, err := file.Open()

			if err != nil {
				return nil, err
			}

			defer reader.Close()

			img, _, err := image.Decode(reader)

			if err != nil {
				return nil, err
			}

			return img, nil
		}
	}

	return nil, fmt.Errorf("file not found: %s", filePath)
}

type Texture struct {
	Caption string
	HTML    string
}

func main() {
	args := os.Args

	if len(args) < 3 {
		fmt.Println("Usage: program source_path output_path [-title \"Page Title\"]")
		return
	}

	settings := ProgramSettings{
		SourcePath:      args[1],
		OutputPath:      args[2],
		PageTitle:       "Textures",
		ThumbnailSize:   128,
		BackgroundColor: color.RGBA{255, 255, 255, 255},
	}

	for i := 3; i < len(args); i += 2 {
		if i+1 < len(args) && args[i] == "-title" {
			settings.PageTitle = args[i+1]
		}
	}

	for i := 3; i < len(args); i += 2 {
		if i+1 < len(args) && args[i] == "-size" {
			if size, err := strconv.Atoi(args[i+1]); err == nil {
				settings.ThumbnailSize = size
			} else {
				fmt.Fprintf(os.Stderr, "Invalid value for -size: %s\n", args[i+1])

				return
			}
		}
	}

	var fileList []string
	var zipReader *zip.ReadCloser
	var err error

	if fileInfo, err := os.Stat(settings.SourcePath); err == nil && fileInfo.IsDir() {
		fileList, err = FileListing(settings.SourcePath)

		if err != nil {
			fmt.Println(err)

			return
		}
	} else {
		zipReader, err = zip.OpenReader(settings.SourcePath)

		if err != nil {
			fmt.Println(err)

			return
		}

		defer zipReader.Close()

		for _, file := range zipReader.File {
			fileList = append(fileList, file.Name)
		}
	}

	families := make(map[string][]Texture)

	var imageObj image.Image

	for _, filePath := range fileList {
		parts := strings.Split(strings.ToLower(filePath), string(filepath.Separator))

		family, filename := parts[len(parts)-2], parts[len(parts)-1]

		extension := filepath.Ext(filename)
		allowedExtensions := map[string]bool{".pcx": true, ".gif": true, ".png": true, ".jpg": true, ".tga": true}

		if !allowedExtensions[extension] || filename == "full.pcx" {
			fmt.Fprintf(os.Stderr, "skipping %s\n", filePath)

			continue
		}

		if fileInfo, _ := os.Stat(settings.SourcePath); fileInfo.IsDir() {
			imageFile, err := os.Open(filePath)

			if err != nil {
				fmt.Println(err)

				return
			}

			defer imageFile.Close()

			if extension == ".pcx" {
				imageObj, err = pcx.Decode(imageFile)
			} else if extension == ".tga" {
				imageObj, err = tga.Decode(imageFile)
			} else {
				imageObj, _, err = image.Decode(imageFile)
			}

			if err != nil {
				fmt.Println(err)

				return
			}
		} else {
			imageObj, err = GetImageFromZip(zipReader, filePath)

			if err != nil {
				fmt.Println(err)

				return
			}
		}

		newBounds := imageObj.Bounds().Max

		if newBounds.X > newBounds.Y {
			newBounds.Y = int(float64(settings.ThumbnailSize) * float64(newBounds.Y) / float64(newBounds.X))
			newBounds.X = settings.ThumbnailSize
		} else {
			newBounds.X = int(float64(settings.ThumbnailSize) * float64(newBounds.X) / float64(newBounds.Y))
			newBounds.Y = settings.ThumbnailSize
		}

		imageObj = resize.Resize(uint(newBounds.X), uint(newBounds.Y), imageObj, resize.Bilinear)

		if imageObj.ColorModel() == color.RGBAModel || imageObj.ColorModel() == color.NRGBAModel {
			backgroundImage := image.NewRGBA(imageObj.Bounds())
			draw.Draw(backgroundImage, backgroundImage.Bounds(), &image.Uniform{settings.BackgroundColor}, image.Point{}, draw.Over)
			draw.Draw(backgroundImage, backgroundImage.Bounds(), imageObj, imageObj.Bounds().Min, draw.Over)
			imageObj = backgroundImage
		}

		buffer := new(bytes.Buffer)
		err := jpeg.Encode(buffer, imageObj, &jpeg.Options{Quality: 100})

		if err != nil {
			fmt.Println(err)

			return
		}

		contentType := "image/jpg"
		encodedImage := base64.StdEncoding.EncodeToString(buffer.Bytes())
		uri := fmt.Sprintf("data:%s;base64,%s", contentType, encodedImage)

		filenameWithoutExtension := strings.TrimSuffix(filepath.Base(filePath), filepath.Ext(filePath))
		imageDimensions := fmt.Sprintf("%dx%d", imageObj.Bounds().Dx(), imageObj.Bounds().Dy())
		imageFormat := strings.TrimPrefix(filepath.Ext(filePath), ".")

		filenameSpan := fmt.Sprintf("<span class='filename'>%s</span>", strings.ToLower(filenameWithoutExtension))
		infoSpan := fmt.Sprintf("<span class='info'>%s (%s)</span>", strings.ToLower(imageDimensions), strings.ToLower(imageFormat))
		caption := fmt.Sprintf("%s %s", filenameSpan, infoSpan)

		texture := Texture{
			Caption: caption,
			HTML:    fmt.Sprintf("<div class='texture'><div class='image'><img src='%s'></div><div class='caption'>%s</div></div>", uri, caption),
		}

		families[family] = append(families[family], texture)
	}

	var familyKeys []string
	for family := range families {
		familyKeys = append(familyKeys, family)
	}

	sort.Strings(familyKeys)

	var sections []string

	for _, family := range familyKeys {
		textures := families[family]

		sort.Slice(textures, func(i, j int) bool {
			return textures[i].Caption < textures[j].Caption
		})

		var texturesHTML []string

		for _, texture := range textures {
			texturesHTML = append(texturesHTML, texture.HTML)
		}

		sections = append(sections, fmt.Sprintf("<section><h2>%s</h2><div class='family'>%s</div></section>", html.EscapeString(family), strings.Join(texturesHTML, "")))
	}

	page := fmt.Sprintf(
		`<!DOCTYPE html>
		<html>
		<head>
		<title>%s</title>
		<style>
		body,h1,h2{color:#fff;font-family:Arial,sans-serif;line-height:1}
		body{background:#333}
		h1{font-size:18px;text-transform:uppercase}
		h2{border-bottom:1px solid #899;font-size:16px;padding:0 0 8px;text-transform:capitalize}
		section{padding:24px 0}
		.family{display:flex;flex-wrap:wrap;gap:16px}
		.texture,.image{width:%dpx}
		.texture{flex:0 0 auto}
		.image{height:%dpx}
		img{width:100%%;height:100%%;object-fit:contain}
		.caption{color:#899;font-size:12px;text-align:center;padding:16px 0;display:flex;flex-direction:column;gap:8px}
		.filename{font-size:14px;font-weight:bold}
		</style>		
		</head>
		<body>
		<h1>%s</h1>
		%s
		</body>
		</html>`,
		html.EscapeString(settings.PageTitle),
		settings.ThumbnailSize,
		settings.ThumbnailSize,
		html.EscapeString(settings.PageTitle),
		strings.Join(sections, ""),
	)

	err = os.WriteFile(settings.OutputPath, []byte(page), 0644)

	if err != nil {
		fmt.Println(err)
	}
}
