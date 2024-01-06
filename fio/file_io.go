// Package fio @Author LJ 2024-01-06 21:23:00
package fio

import "os"

// FileIO 对标准文件系统的一些接口进行简单封装
type FileIO struct {
	//系统文件描述符
	fd *os.File
}

// NewFileIO 初始化标准文件IO
func NewFileIO(filename string) (*FileIO, error) {
	fd, err := os.OpenFile(
		filename,
		os.O_RDWR|os.O_CREATE|os.O_APPEND,
		FilePermCode)
	if err != nil {
		return nil, err
	}
	return &FileIO{fd: fd}, nil
}

// Read 从文件指定位置读取数据
func (fio *FileIO) Read(b []byte, offset int64) (int, error) {
	// ReadAt reads len(b) bytes from the File starting at byte offset off.
	// It returns the number of bytes read and the error, if any.
	// ReadAt always returns a non-nil error when n < len(b).
	// At end of file, that error is io.EOF.
	return fio.fd.ReadAt(b, offset)
}

// Write 写入数据到文件指定位置
func (fio *FileIO) Write(b []byte) (int, error) {
	// Write writes len(b) bytes to the File.
	// It returns the number of bytes written and an error, if any.
	// Write returns a non-nil error when n != len(b).
	return fio.fd.Write(b)
}

// Sync 将存放在内存缓冲区的数据持久化到磁盘中
func (fio *FileIO) Sync() error {
	// Sync commits the current contents of the file to stable storage.
	// Typically, this means flushing the file system's in-memory copy
	// of recently written data to disk.
	return fio.fd.Sync()
}

// Close 关闭文件
func (fio *FileIO) Close() error {
	return fio.fd.Close()
}
