// Package trie implements the trie data structure
// See https://en.wikipedia.org/wiki/Trie
package trie

import (
	"github.com/markelog/trie/node"
)

// Trie essential structure
type Trie struct {
	Root *node.Node
}

// Walker for the traverse function
// If such function returns false it will stop traverse
type Walker func(item *node.Node) bool

// New returns new Trie
func New() *Trie {
	return &Trie{
		Root: node.New("", ""),
	}
}

// Add stuff to the trie
func (trie Trie) Add(key string, value interface{}) Trie {
	var (
		current = trie.Root
		bKey    = []byte(key)
		length  = len(bKey)
	)

	for i := 0; i < length; i++ {
		newNode := node.New(string(bKey[i]), nil)
		newNode.Parent = current

		// Check to see if character node exists in children.
		if current.Keys[bKey[i]] == nil {
			newNode.Leaf = false

			current.Keys[bKey[i]] = newNode
		}

		// Check if its last letter
		if i == length-1 {
			println(current == trie.Root)
			// Then its a leaf
			newNode.Leaf = true
			newNode.Value = value
			current.Children = append(current.Children, newNode)

			// litter.Dump(current.Children)
		}

		current = current.Keys[bKey[i]]
	}

	return trie
}

// Contains check presence of the key in the trie
func (trie Trie) Contains(key string) bool {
	var (
		current = trie.Root
		bKey    = []byte(key)
		length  = len(bKey)
	)

	// for every character in the word
	for i := 0; i < length; i++ {

		// Check if we have such, since if its not then we can abort
		if current.Keys[bKey[i]] == nil {
			return false
		}

		// Key exist - proceed to the next depth of the trie
		current = current.Keys[bKey[i]]
	}

	// We finished going through all the words, but is it a whole word?
	return current.Leaf
}

// Returns every word with the given prefix
func (trie Trie) Find(prefix string) (result []*node.Node) {
	var (
		current = trie.Root
		bPrefix = []byte(prefix)
		length  = len(bPrefix)
	)

	for i := 0; i < length; i++ {

		// If we don't have anything anymore - return what we got
		if current.Keys[bPrefix[i]] == nil {
			return
		}

		// Proceed forward then
		current = current.Keys[bPrefix[i]]
	}

	return findAll(current, result)
}

// findAll recursively find the words
func findAll(current *node.Node, tmp []*node.Node) (result []*node.Node) {
	if current.Leaf {
		result = append(tmp, current)
	}

	for key, _ := range current.Keys {
		result = findAll(current.Keys[key], result)
	}

	return
}

// Traverse the leaves (not the branches)
func (trie Trie) Traverse(fn Walker) {
	trie.Visit(trie.Root, fn)
}

// Visit specific part of the tree
func (trie Trie) Visit(current *node.Node, fn Walker) {
	var (
		children = current.Children
		length   = len(children)
	)

	for i := 0; i < length; i++ {
		result := fn(children[i])
		if result == false {
			return
		}

		trie.Visit(children[i], fn)
	}

	return
}
