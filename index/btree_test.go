// @Author LJ 2024-01-04 22:53:00
package index

import (
	"bitcask/data"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBTree_Put(t *testing.T) {
	bTree := NewBTree()

	rep1 := bTree.Put(nil, &data.LogRecordPos{Fid: 1, Offset: 10})
	assert.True(t, rep1)

	rep2 := bTree.Put([]byte("a"), &data.LogRecordPos{Fid: 1, Offset: 20})
	assert.True(t, rep2)
}

func TestBTree_Get(t *testing.T) {
	bTree := NewBTree()

	rep1 := bTree.Put(nil, &data.LogRecordPos{Fid: 1, Offset: 10})
	assert.True(t, rep1)

	pos1 := bTree.Get(nil)
	assert.Equal(t, uint32(1), pos1.Fid)
	assert.Equal(t, int64(10), pos1.Offset)

	rep2 := bTree.Put([]byte("a"), &data.LogRecordPos{Fid: 1, Offset: 20})
	assert.True(t, rep2)
	rep3 := bTree.Put([]byte("a"), &data.LogRecordPos{Fid: 1, Offset: 21})
	assert.True(t, rep3)

	pos2 := bTree.Get([]byte("a"))
	assert.Equal(t, uint32(1), pos2.Fid)
	assert.Equal(t, int64(21), pos2.Offset)
}

func TestBTree_Delete(t *testing.T) {
	bTree := NewBTree()
	//1
	rep1 := bTree.Put(nil, &data.LogRecordPos{Fid: 1, Offset: 10})
	assert.True(t, rep1)

	pos1 := bTree.Get(nil)
	assert.Equal(t, uint32(1), pos1.Fid)
	assert.Equal(t, int64(10), pos1.Offset)

	del1 := bTree.Delete(nil)
	assert.True(t, del1)

	pos2 := bTree.Get(nil)
	assert.Nil(t,pos2)

	//2
	rep2 := bTree.Put([]byte("aa"), &data.LogRecordPos{Fid: 2, Offset: 20})
	assert.True(t, rep2)

	pos3 := bTree.Get([]byte("aa"))
	assert.Equal(t, uint32(2), pos3.Fid)

	del2 := bTree.Delete([]byte("aa"))
	assert.True(t, del2)

	pos4 := bTree.Get([]byte("aa"))
	assert.Nil(t, pos4)
}