package frequency

import (
	"fmt"
	"strings"
)

type (
	// deltaListNode is a node in the linked list of the delta list.
	deltaListNode struct {
		next    *deltaListNode
		entries hashTableEntries
		delta   uint32
	}

	// deltaList is a collection of counts based on growing differences.
	deltaList struct {
		count int
		head  *deltaListNode
	}
)

// newDeltaList creates a new delta list.
func newDeltaList() *deltaList {
	return &deltaList{
		count: 0,
		head:  nil,
	}
}

// Count is the number of entries in the delta list.
func (dl *deltaList) Count() int {
	return dl.count
}

// AllEntries gets all the entries from the delta list.
// Values are listed with highest count first.
func (dl *deltaList) AllEntries() hashTableEntries {
	entries := make(hashTableEntries, dl.count)
	i := dl.count - 1
	for node := dl.head; node != nil; node = node.next {
		for _, entry := range node.entries {
			entries[i] = entry
			i--
		}
	}
	return entries
}

// insertAt inserts an entry into the list at or before the given node.
// The node maybe null. The new node or given node is returned.
func (dl *deltaList) insertAt(entry *hashTableEntry, node *deltaListNode) *deltaListNode {
	if node != nil {
		if node.delta == 1 {
			node.entries = append(node.entries, entry)
			entry.node = node
			return node
		}
		node.delta--
	}

	newNode := &deltaListNode{
		next:    node,
		entries: hashTableEntries{entry},
		delta:   1,
	}
	entry.node = newNode
	return newNode
}

// Insert will insert a new entry into the delta list.
func (dl *deltaList) Insert(entry *hashTableEntry) {
	dl.head = dl.insertAt(entry, dl.head)
	dl.count++
}

// DecrementAll will decrement the value of all items in the list,
// by decrementing the first node and removing it if the new count is zero.
// Returns any entry which no longer exists in the delta list.
func (dl *deltaList) DecrementAll() hashTableEntries {
	if dl.head != nil {
		dl.head.delta--
		if dl.head.delta <= 0 {
			node := dl.head
			dl.head = node.next
			dl.count -= len(node.entries)
			return node.entries
		}
	}
	return nil
}

// Increment will increase the count of the node for this entry
// but not for the other entries in the same node.
// The given entry must already exist in the delta list.
func (dl *deltaList) Increment(entry *hashTableEntry) {
	node := entry.node
	if len(node.entries) == 1 {
		node.delta++
		if next := node.next; next != nil {
			if next.delta > 1 {
				next.delta--
				return
			}

			node.entries = append(node.entries, next.entries...)
			for _, entry := range next.entries {
				entry.node = node
			}
			node.next = next.next
		}
		return
	}

	node.entries = node.entries.removeEntry(entry)
	node.next = dl.insertAt(entry, node.next)
}

// String gets a human readable string for debugging.
func (dl *deltaList) String() string {
	parts := []string{}
	sum := uint32(0)
	for node := dl.head; node != nil; node = node.next {
		entries := strings.Join(node.entries.getData(), `, `)
		sum += node.delta
		parts = append(parts, fmt.Sprintf(`%d: [%s]`, sum, entries))
	}
	return strings.Join(parts, `, `)
}
