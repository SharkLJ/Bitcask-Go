// Package index @Author LJ 2024-01-04 22:08:00
package index

import (
	"bitcask/data"
	"github.com/google/btree"
	"sync"
)

// BTree 索引数据结构，封装了google的btree库
type BTree struct {
	tree *btree.BTree
	lock *sync.RWMutex
}

// NewBTree 初始化一个BTree对象
func NewBTree() *BTree {
	return &BTree{
		//这里的degree指的是BTree的度，也就是每个节点最多有多少个子节点
		tree: btree.New(32),
		//lock: &sync.RWMutex{},
		lock: new(sync.RWMutex),
	}
}

func (bt *BTree) Put(key []byte, pos *data.LogRecordPos) bool {
	it := &Item{key: key, pos: pos}
	bt.lock.Lock()
	bt.tree.ReplaceOrInsert(it)
	defer bt.lock.Unlock()
	return true
}

func (bt *BTree) Get(key []byte) *data.LogRecordPos {
	it := &Item{key: key}
	bt.lock.RLock()
	value := bt.tree.Get(it)
	bt.lock.RUnlock()
	if value == nil {
		return nil
	}
	return value.(*Item).pos
}

func (bt *BTree) Delete(key []byte) bool {
	it := &Item{key: key}
	bt.lock.Lock()
	//会返回被删除的值，如果btree中没有这个值，则返回nil
	oldItem := bt.tree.Delete(it)
	bt.lock.Unlock()
	if oldItem == nil {
		return false
	}
	return true
}
