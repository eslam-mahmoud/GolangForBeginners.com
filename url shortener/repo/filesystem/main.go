package filesystem

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strconv"
	"sync/atomic"
)

// Storage represent filesystem storage
// use NewStorage as constructor
type Storage struct {
	lastID uint64
	dir    string
}

// NewStorage constructor for Storage
func NewStorage(dir string) (*Storage, error) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	return &Storage{
		lastID: uint64(len(files)),
		dir:    dir,
	}, nil
}

// Shorten take file content save it and retun id or error
func (s *Storage) Shorten(content string) (string, error) {
	fmt.Println(s.lastID)
	fileID := strconv.FormatUint(atomic.AddUint64(&s.lastID, 1), 10)
	f, err := os.Create(path.Join(s.dir, fileID))
	if err != nil {
		return "", err
	}
	defer f.Close()
	_, err = f.WriteString(content)
	if err != nil {
		return "", err
	}
	return string(fileID), nil
}

// Expand return file content by id
func (s *Storage) Expand(id string) (string, error) {
	content, err := ioutil.ReadFile(path.Join(s.dir, id))
	if err != nil {
		return "", err
	}
	return string(content), nil
}
