package sort

import (
	"strings"
	"unicode/utf8"

	"github.com/untref-ayp2/data-structures/dictionary"
	"github.com/untref-ayp2/data-structures/list"
)

const abecedario = "*.0123456789abcdefghijklmnñopqrstuvwxyz"

func RadixSort(arr []string) {
	if len(arr) < 2 {
		return
	}
	buckets := dictionary.NewDictionary[string, list.LinkedList[string]]()
	initBuckets(buckets)
	length := maxLen(arr)
	for i := length - 1; i >= 0; i-- {
		fillBuckets(arr, buckets, i)
		emptyBuckets(arr, buckets)
	}
}

func initBuckets(b *dictionary.Dictionary[string, list.LinkedList[string]]) {
	for _, key := range abecedario {
		b.Put(string(key), *list.NewLinkedList[string]())
	}
}

func fillBuckets(arr []string, b *dictionary.Dictionary[string, list.LinkedList[string]], pos int) {
	for _, value := range arr {
		key := "*"
		if pos < utf8.RuneCountInString(value) {
			key = strings.ToLower(string([]rune(value)[pos]))
		}
		list, _ := b.Get(key)
		list.Append(value)
		b.Put(key, list)
	}
}

func emptyBuckets(arr []string, b *dictionary.Dictionary[string, list.LinkedList[string]]) {
	index := 0
	for _, key := range abecedario {
		list, _ := b.Get(string(key))
		for !list.IsEmpty() {
			arr[index] = list.Head().Data()
			list.RemoveFirst()
			index++
		}
		b.Put(string(key), list)
	}
}

func maxLen(arr []string) int {
	max := utf8.RuneCountInString(arr[0])
	for _, value := range arr {
		if utf8.RuneCountInString(value) > max {
			max = utf8.RuneCountInString(value)
		}
	}
	return max
}
