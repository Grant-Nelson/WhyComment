package frequency

import (
	"sync"
)

type Frequency struct {
	size  int
	list  *deltaList
	table *hashTable
	lock  *sync.RWMutex
}

func New(size, hashTableSize int, hash HashFunction) *Frequency {
	return &Frequency{
		size:  size,
		list:  newDeltaList(),
		table: newHashTable(hashTableSize, hash),
		lock:  &sync.RWMutex{},
	}
}

func (f *Frequency) Add(data string) {
	f.lock.Lock()
	defer f.lock.Unlock()

	if entry := f.table.FindEntry(data); entry != nil {
		f.list.Increment(entry)
		return
	}

	if f.list.Count() < f.size {
		entry := f.table.InsertEntry(data)
		f.list.Insert(entry)
		return
	}

	entries := f.list.DecrementAll()
	for _, entry := range entries {
		f.table.RemoveEntry(entry)
	}
}

func (f *Frequency) Results() []string {
	f.lock.RLock()
	defer f.lock.RUnlock()
	return f.list.AllEntries().getData()
}
