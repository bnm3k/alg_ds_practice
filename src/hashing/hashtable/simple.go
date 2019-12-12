package hashtable

import (
	"fmt"
)

type entry struct {
	key   string
	value string
}

//HashTable ...
type HashTable struct {
	size    int
	count   int
	entries []entry
}

/*
func (ht *hashTable) insert(key, value string) {
	e := entry{key, value}
	attempt := 0
	index := ht.getHash(key, attempt)
	attempt = 1
	curr := ht.entries[index]
	for curr != nil {

	}

}*/

func (ht *HashTable) getHash(key string, attempt int) int {
	prime1 := 3
	prime2 := 7
	hashA := hash(key, prime1, ht.size)
	hashB := hash(key, prime2, ht.size)
	return (hashA + (attempt * (hashB + 1))) % ht.size
}

//NewHashTable Returns new hash Table
func NewHashTable() *HashTable {
	size := 53
	return &HashTable{size, 0, make([]entry, size)}
}

func pow(a, b int) int {
	var p int = 1
	for b > 0 {
		if b&1 != 0 {
			p *= a
		}
		b >>= 1
		a *= a
	}
	return p
}

func hash(str string, a, b int) int {
	var hash int64 = 0
	strLen := len(str)
	for i, char := range str {
		hash += int64(pow(a, strLen-i-1) * int(char))
		hash = hash % int64(b)
	}
	return int(hash)
}

func main() {
	//dict := NewHashTable()
	//fmt.Println(dict)
	val := hash("cat", 163, 53)
	fmt.Println(val)

}

func print(s interface{}) {
	fmt.Println(s)
}
