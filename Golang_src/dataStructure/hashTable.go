package dataStructure

const (
	defaultTableSize = 10
)

type Record struct {
	next *Record
	key  int
	val  int
}

type Hash struct {
	records []*Record
}

type HashTable struct {
	table        *Hash
	numOfRecords int
}

func NewHashTable() *HashTable {
	num := 0
	hash := &Hash{make([]*Record, defaultTableSize)}
	return &HashTable{table: hash, numOfRecords: num}
}

func hashFunction(key, size int) int {
	return key % size
}

func (h *HashTable) Put(key, value int) bool {
	index := hashFunction(key, len(h.table.records))
	cur := h.table.records[index]
	node := &Record{key: key, val: value}
	if cur == nil {
		h.table.records[index] = node
	} else {
		pre := &Record{}
		for cur != nil {
			if cur.key == key {
				cur.val = value
				return false
			}
			pre = cur
			cur = cur.next
		}
		pre.next = cur
	}
	h.numOfRecords++
	return true
}

func (h *HashTable) Get(key int) (bool, int) {
	index := hashFunction(key, len(h.table.records))
	cur := h.table.records[index]
	for cur != nil {
		if cur.key == key {
			return true, cur.val
		}
		cur = cur.next
	}
	return false, 0
}

func (h *HashTable) Remove(key int) bool {
	index := hashFunction(key, len(h.table.records))
	cur := h.table.records[index]
	if cur == nil {
		return false
	}
	if cur.key == key {
		h.table.records[index] = cur.next
		h.numOfRecords--
		return true
	} else {
		pre := cur
		cur = cur.next
		for cur != nil {
			if cur.key == key {
				pre.next = cur.next
				h.numOfRecords--
				return true
			}
			pre = cur
			cur = cur.next
		}
	}
	return false
}

func (h *HashTable) Size() int {
	return h.numOfRecords
}
