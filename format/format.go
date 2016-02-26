package format

import (
	"io"
	"strings"

	"github.com/songtools/songtools"
)

// Reader represents the ability to read a SongSet.
type Reader interface {
	Read(io.Reader) (*songtools.Song, error)
}

// Writer represents the ability to write a SongSet.
type Writer interface {
	Write(io.Writer, *songtools.Song) error
}

// Format represents a named ability to read and write a SongSet.
type Format struct {
	Name   string
	Reader Reader
	Writer Writer
}

var registeredFormats = []*Format{}

// Formats returns all the registered formats.
func Formats() []*Format {
	return registeredFormats
}

// Names returns the names of the registered formats.
func Names() []string {
	names := []string{}
	for _, f := range registeredFormats {
		names = append(names, f.Name)
	}

	return names
}

// ByName returns a registered format by name. It compares names case-insensitively.
func ByName(name string) (*Format, bool) {

	name = strings.ToLower(name)
	for _, f := range registeredFormats {
		if strings.ToLower(f.Name) == name {
			return f, true
		}
	}

	return nil, false
}

// Register registers a format.
func Register(f *Format) {
	registeredFormats = append(registeredFormats, f)
}