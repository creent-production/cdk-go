package filestatic

import "net/http"

// FileSystem custom file system handler
type FileSystem struct {
	Static http.FileSystem
}

// Open opens file
func (fs FileSystem) Open(path string) (http.File, error) {
	f, err := fs.Static.Open(path)
	if err != nil {
		return nil, err
	}

	s, err := f.Stat()
	if err != nil {
		return nil, err
	}

	if s.IsDir() {
		return fs.Static.Open("upps!")
	}

	return f, nil
}
