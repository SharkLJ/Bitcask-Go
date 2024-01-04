package index

import (
	"bitcask/data"
	"bytes"
	"github.com/google/btree"
)

// Indexer 抽象索引接口，如果要实现不同的索引结构，只需要实现该接口即可
type Indexer interface {
	// Put 往索引中 添加key对应的索引位置信息
	Put(key []byte, pos *data.LogRecordPos) bool

	// Get 根据key获取索引位置信息
	Get(key []byte) *data.LogRecordPos

	// Delete 根据key删除索引位置信息
	Delete(key []byte) bool
}

type Item struct {
	key []byte
	pos *data.LogRecordPos
}

// Less 实现btree.Item接口的Less方法,因为BTree自带的ReplaceOrInsert方法需要比较key的大小
func (itemA *Item) Less(itemB btree.Item) bool {
	//.(*Item) 是一个类型断言，它尝试将 itemB 断言为 *Item 类型。
	//如果 itemB 确实是 *Item 类型，那么这个断言就会成功，我们就可以访问其 key 字段。
	//如果 itemB 不是 *Item 类型，那么这个断言就会失败，程序会抛出一个运行时错误
	return bytes.Compare(itemA.key, itemB.(*Item).key) == -1
}