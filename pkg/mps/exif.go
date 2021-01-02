package mps

import (
	"fmt"
	"io"
	"time"

	"github.com/dsoprea/go-exif/v3"
	jpegis "github.com/dsoprea/go-jpeg-image-structure/v2"
	"github.com/h2non/filetype/matchers"
	"github.com/h2non/filetype/types"
)

const exifTimeFormat = "2006:01:02 15:04:05"

// getExifData gets the EXIF tags from a supported file
func getExifData(f io.ReadSeeker, fileSize int, fileType types.Type) (*exif.Ifd, error) {
	switch fileType {
	case matchers.TypeJpeg:
		{
			ec, err := jpegis.NewJpegMediaParser().Parse(f, fileSize)
			if err != nil {
				return nil, fmt.Errorf("could not parse jpeg file: %w", err)
			}

			rootIfd, _, err := ec.Exif()
			if err != nil {
				return nil, fmt.Errorf("could not get exif data from jpeg file: %w", err)
			}

			return rootIfd, nil
		}
	}

	return nil, fmt.Errorf("unimplemented file type: %v", fileType.Extension)
}

// getDate gets the photo date from the EXIF data
func getDate(ifd *exif.Ifd) (*time.Time, error) {

	// TODO! Implement support for TimeZoneOffset EXIF tag

	rootTags := ifd.DumpTags()
	foundTime := time.Time{}

	for _, t := range rootTags {
		if t.TagName() == "DateTime" {
			tagString, err := t.Format()
			if err != nil {
				return nil, fmt.Errorf("could not format exif datetime string: %w", err)
			}

			foundTime, err = time.Parse(exifTimeFormat, tagString)
			if err != nil {
				return nil, fmt.Errorf("could not parse formatted datetime string: %w", err)
			}
			break
		}
	}

	if (foundTime == time.Time{}) {
		return nil, fmt.Errorf("no datetime found")
	}

	return &foundTime, nil
}
