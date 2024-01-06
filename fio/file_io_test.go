// @Author LJ 2024-01-06 22:14:00
package fio

import (
	"github.com/stretchr/testify/assert"
	"os"
	"path/filepath"
	"testing"
)

func destroyFile(path string) {
	if err := os.RemoveAll(path); err != nil{
		panic(err)
	}
}

func TestNewFileIOManager(t *testing.T) {
	path := filepath.Join("/tmp", "test.data")
	io, err := NewFileIO(path)
	cerr := io.Close()
	assert.Nil(t, cerr)
	defer destroyFile(path)
	assert.Nil(t, err)
	assert.NotNil(t, io)
}

func TestFileIO_Writer(t *testing.T) {
	path := filepath.Join("/tmp", "test.data")
	io, err := NewFileIO(path)

	defer destroyFile(path)

	assert.Nil(t, err)
	assert.NotNil(t, io)

	n, err := io.Write([]byte("hello"))
	assert.Nil(t, err)
	assert.Equal(t, 5, n)

	n, err = io.Write([]byte("world"))
	assert.Nil(t, err)
	assert.Equal(t, 5, n)

	cerr := io.Close()
	assert.Nil(t, cerr)
}

func TestFileIO_Read(t *testing.T) {
	path := filepath.Join("/tmp", "test.data")
	io, err := NewFileIO(path)

	defer destroyFile(path)

	assert.Nil(t, err)
	assert.NotNil(t, io)

	n, err := io.Write([]byte("hello"))
	assert.Nil(t, err)
	assert.Equal(t, 5, n)

	n, err = io.Write([]byte("world"))
	assert.Nil(t, err)
	assert.Equal(t, 5, n)

	b := make([]byte, 5)
	n, err = io.Read(b, 0)
	assert.Nil(t, err)
	assert.Equal(t, 5, n)
	assert.Equal(t, b, []byte("hello"))

	b = make([]byte, 5)
	n, err = io.Read(b, 5)
	assert.Nil(t, err)
	assert.Equal(t, 5, n)
	assert.Equal(t, b, []byte("world"))

	cerr := io.Close()
	assert.Nil(t, cerr)
}

func TestFileIO_Sync(t *testing.T) {
	path := filepath.Join("/tmp", "test.data")
	io, err := NewFileIO(path)

	defer destroyFile(path)

	assert.Nil(t, err)
	assert.NotNil(t, io)

	err = io.Sync()
	assert.Nil(t, err)

	cerr := io.Close()
	assert.Nil(t, cerr)
}

func TestFileIO_Close(t *testing.T) {
	path := filepath.Join("/tmp", "test.data")
	io, err := NewFileIO(path)
	defer destroyFile(path)

	assert.Nil(t, err)
	assert.NotNil(t, io)

	err = io.Close()
	assert.Nil(t, err)
}