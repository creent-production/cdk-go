package filestatic

import "net/http"

// FileSystem custom file system handler
type FileSystem struct {
	http.FileSystem
}

// Open opens file
func (fs FileSystem) Open(path string) (http.File, error) {
	f, err := fs.Open(path)
	if err != nil {
		return nil, err
	}

	s, err := f.Stat()
	if err != nil {
		return nil, err
	}

	if s.IsDir() {
		return fs.Open("upps!")
	}

	return f, nil
}
