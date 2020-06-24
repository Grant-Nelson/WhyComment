package frequency

import (
	"strings"
)

type (
	// HashFunction is a hash function definition.
	HashFunction func(data string) uint64

	// hashTableEntry is an entry into a hash table.
	// This entry also points to a node in the deltaList.
	hashTableEntry struct {
		data  string
		index int
		node  *deltaListNode
	}

	// hashTableEntries is a collection of hash table entries.
	hashTableEntries []*hashTableEntry
)

// newHashTableEntry creates a new hash table entry.
func newHashTableEntry(data string) *hashTableEntry {
	return &hashTableEntry{
		data:  data,
		index: -1,
		node:  nil,
	}
}

// removeEntry removes the given entry from this slice of entries.
// It returns the new slice without the entry or nil if slice is empty.
func (entries hashTableEntries) removeEntry(entry *hashTableEntry) hashTableEntries {
	for i, other := range entries {
		if entry == other {
			maxIndex := len(entries) - 1
			entries[i] = entries[maxIndex]
			return entries[:maxIndex]
		}
	}
	return entries
}

// getData gets all the data from these entries.
func (entries hashTableEntries) getData() []string {
	result := make([]string, len(entries))
	for i, entry := range entries {
		result[i] = entry.data
	}
	return result
}

// hashTable is a hash table which uses a slice to contain
// data with equal hash values. This is designed to work with
// the deltaList as part of a Frequency Algorithm.
type hashTable struct {
	hash  HashFunction
	table []hashTableEntries
}

// newHashTable creates a new hash table.
func newHashTable(size int, hash HashFunction) *hashTable {
	return &hashTable{
		hash:  hash,
		table: make([]hashTableEntries, size),
	}
}

// FindEntry will look up the entry for the given data.
// If no entry is found then nil is returned.
func (ht *hashTable) FindEntry(data string) *hashTableEntry {
	index := int(ht.hash(data) % uint64(len(ht.table)))
	for _, entry := range ht.table[index] {
		if entry.data == data {
			return entry
		}
	}
	return nil
}

// InsertEntry will insert a new entry for the given data.
// Warning: This does not check if the data is already in the table,
// so use FindEntry first to prevent repeat values.
func (ht *hashTable) InsertEntry(data string) *hashTableEntry {
	index := int(ht.hash(data) % uint64(len(ht.table)))
	entry := newHashTableEntry(data)
	entry.index = index
	ht.table[index] = append(ht.table[index], entry)
	return entry
}

// RemoveEntry will remove the given entry from the table.
func (ht *hashTable) RemoveEntry(entry *hashTableEntry) {
	if entry != nil {
		index := entry.index
		ht.table[index] = ht.table[index].removeEntry(entry)
	}
}

// String gets the human readable string for debugging.
func (ht *hashTable) String() string {
	parts := make([]string, len(ht.table))
	for i, entry := range ht.table {
		if len(entry) <= 0 {
			parts[i] = `-`
		} else {
			parts[i] = `[` + strings.Join(entry.getData(), `, `) + `]`
		}
	}
	return strings.Join(parts, `, `)
}
