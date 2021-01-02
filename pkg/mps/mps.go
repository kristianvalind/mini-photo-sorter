package mps

import (
	"fmt"
	"os"
	"time"

	"github.com/h2non/filetype"
	"github.com/h2non/filetype/matchers"
)

// Types we can extract EXIF or similar data from
var exifTypes = matchers.Map{
	matchers.TypeJpeg: matchers.Jpeg,
	matchers.TypeHeif: matchers.Heif,
	matchers.TypeGif:  matchers.Gif,
	matchers.TypePng:  matchers.Png,
	matchers.TypeTiff: matchers.Tiff,
}

var otherTypes = matchers.Map{
	matchers.TypeMp4: matchers.Mp4,
	matchers.TypeMov: matchers.Mov,
	matchers.TypeMkv: matchers.Mkv,
}

// GetDate deducts the date a photo or movie (stored in a file) was taken
func GetDate(path string) (*time.Time, error) {
	// Get file info
	fInfo, err := os.Stat(path)
	if err != nil {
		return nil, fmt.Errorf("could not get file info: %w", err)
	}

	if fInfo.Size() < 261 {
		return nil, fmt.Errorf("file too small")
	}

	// Read file header to deduce file type
	headf, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("could not open file for header read: %w", err)
	}
	defer headf.Close()

	head := make([]byte, 261)
	_, err = headf.Read(head)
	if err != nil {
		return nil, fmt.Errorf("could not read file header: %w", err)
	}
	headf.Close()

	detectedType, err := filetype.Get(head)
	if err != nil {
		return nil, fmt.Errorf("could not get file type: %w", err)
	}

	if filetype.MatchesMap(head, exifTypes) {
		f, err := os.Open(path)
		if err != nil {
			return nil, fmt.Errorf("could not open file for exif extraction: %w", err)
		}
		defer f.Close()

		exifData, err := getExifData(f, int(fInfo.Size()), detectedType)
		if err != nil {
			return nil, fmt.Errorf("could not extract exif data: %w", err)
		}
		f.Close()

		date, err := getDate(exifData)
		if err != nil {
			return nil, fmt.Errorf("could not get date from exif data: %w", err)
		}

		return date, nil
	} else if filetype.MatchesMap(head, otherTypes) {

		return nil, fmt.Errorf("non exif types not yet implemented set: %v", detectedType)
	}

	return nil, fmt.Errorf("unsupported file type: %v", detectedType.Extension)

}
