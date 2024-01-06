// Package fio @Author LJ 2024-01-06 21:23:00
package fio

const FilePermCode = 0644

// IOManager 文件IO管理器
type IOManager interface {
	// Read 从文件指定位置读取数据
	Read(b []byte, offset int64) (int, error)
	// Write 写入数据到文件指定位置
	Write(b []byte) (int, error)
	// Sync 将存放在内存缓冲区的数据持久化到磁盘中
	Sync() error
	// Close 关闭文件
	Close() error
}
