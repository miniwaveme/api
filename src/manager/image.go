package manager

import (
	"fmt"
	"github.com/miniwaveme/api/src/conf"
	"github.com/miniwaveme/api/src/document"
	"github.com/nfnt/resize"
	"github.com/satori/go.uuid"
	"image"
	"image/jpeg"
	"os"
	"strings"
)

const SmallWidth = 64
const MediumWidth = 128
const LargeWidth = 256

func getDefaultArtistImage() *document.Image {
	image := &document.Image{}
	return image
}

func getDefaultAlbumImage() *document.Image {
	image := &document.Image{}
	return image
}

func CreateImage(image image.Image) (*document.Image, error) {
	dirs := strings.Split(uuid.NewV4().String(), "-")
	dir := fmt.Sprintf("%s/%s/%s/%s/%s", dirs[0], dirs[1], dirs[2], dirs[3], dirs[4])

	imageDocument := &document.Image{}
	imageOriginal, err := createOrginalImage(dir, image)
	if err != nil {
		return imageDocument, err
	}
	imageDocument.Original = imageOriginal

	imageSmall, err := createSmallImage(dir, image)
	if err != nil {
		return imageDocument, err
	}
	imageDocument.Small = imageSmall

	imageMedium, err := createMediumImage(dir, image)
	if err != nil {
		return imageDocument, err
	}
	imageDocument.Medium = imageMedium

	imageLarge, err := createLargeImage(dir, image)
	if err != nil {
		return imageDocument, err
	}
	imageDocument.Large = imageLarge

	return imageDocument, nil
}

func createOrginalImage(dir string, image image.Image) (string, error) {
	return saveImage(image, dir, "original")
}

func createSmallImage(dir string, image image.Image) (string, error) {

	resizeImage := resize.Resize(SmallWidth, 0, image, resize.Lanczos3)
	return saveImage(resizeImage, dir, "small")
}

func createMediumImage(dir string, image image.Image) (string, error) {
	resizeImage := resize.Resize(MediumWidth, 0, image, resize.Lanczos3)
	return saveImage(resizeImage, dir, "medium")
}

func createLargeImage(dir string, image image.Image) (string, error) {
	resizeImage := resize.Resize(LargeWidth, 0, image, resize.Lanczos3)
	return saveImage(resizeImage, dir, "large")
}

func saveImage(image image.Image, dir string, typeString string) (string, error) {
	uploadDir := conf.C().GetString("upload.directory")
	filename := fmt.Sprintf("%s.jpg", typeString)
	saveDirPath := fmt.Sprintf("%s/%s", uploadDir, dir)
	saveFilePath := fmt.Sprintf("%s/%s", saveDirPath, filename)

	err := os.MkdirAll(saveDirPath, 777)
	if err != nil {
		return "", err
	}

	out, err := os.Create(saveFilePath)
	if err != nil {
		return "", err
	}

	defer out.Close()

	jpeg.Encode(out, image, nil)
	return fmt.Sprintf("%s/%s/%s", makeUploadHost(), dir, filename), nil
}

func makeUploadHost() string {
	return fmt.Sprintf("%s://%s:%s", conf.C().GetString("upload.scheme"), conf.C().GetString("upload.host"), conf.C().GetString("upload.port"))
}
