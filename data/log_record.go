// Package data @Author LJ 2024-01-03 23:19:00
package data

// LogRecordPos 数据内存索引，记录数据在磁盘中的位置
type LogRecordPos struct {
	Fid    uint32 // 文件id，表示数据存放在哪个文件中
	Offset int64  // 表示将数据存放在了文件中的哪个位置
}
