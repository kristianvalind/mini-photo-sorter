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

type exifWrapper struct {
	rootIfd *exif.Ifd
	data    []byte
}

// getExifData gets the EXIF tags from a supported file
func getExifData(f io.ReadSeeker, fileSize int, fileType types.Type) (*exifWrapper, error) {
	switch fileType {
	case matchers.TypeJpeg:
		{
			ec, err := jpegis.NewJpegMediaParser().Parse(f, fileSize)
			if err != nil {
				return nil, fmt.Errorf("could not parse jpeg file: %w")
			}

			rootIfd, data, err := ec.Exif()
			if err != nil {
				return nil, fmt.Errorf("could not get exif data from jpeg file: %w")
			}

			return &exifWrapper{
				rootIfd: rootIfd,
				data:    data,
			}, nil
		}
	}

	return nil, fmt.Errorf("unimplemented file type: %w", fileType.Extension)
}

// getDate gets the photo date from the wrapped EXIF data
func (w *exifWrapper) getDate() (*time.Time, error) {

	captureIfd, err := exif.FindIfdFromRootIfd()

}
