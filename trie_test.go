package trie_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/markelog/trie"
	"github.com/markelog/trie/node"
)

var _ = Describe("trie", func() {
	Describe("Contains and Add", func() {
		It("should add node to the root", func() {
			trie := New()
			trie.Add("test", 1)

			Expect(trie.Root.Children).To(HaveLen(1))
			Expect(trie.Root.Children[0].Value.(int)).To(Equal(1))
			Expect(trie.Root.Children[0].Parent).To(Equal(trie.Root))
		})

		It("should add node to the trie", func() {
			trie := New()
			trie.Add("test", 1)

			Expect(trie.Contains("test")).To(Equal(true))
		})

		It("should not show false positive for non-existent branch/leaf", func() {
			trie := New()

			Expect(trie.Contains("test")).To(Equal(false))
		})

		It("should not show false positive for non-existent leaf", func() {
			trie := New()

			trie.Add("test", 1)

			Expect(trie.Contains("tes")).To(Equal(false))
		})
	})

	Describe("Find", func() {
		It("should find all the words", func() {
			trie := New()
			trie.Add("t", 1)
			trie.Add("te", "foo")
			trie.Add("tes", 3)
			trie.Add("test", 4)

			result := trie.Find("t")

			Expect(result).To(HaveLen(4))

			Expect(result[0].Value.(int)).To(Equal(1))
			Expect(result[1].Value.(string)).To(Equal("foo"))
		})

		It("should not find anything", func() {
			trie := New()
			trie.Add("t", 1)
			trie.Add("te", 1)
			trie.Add("tes", 1)
			trie.Add("test", 1)

			Expect(trie.Find("q")).To(HaveLen(0))
		})
	})

	Describe("Traverse", func() {
		FIt("should traverse all the nodes", func() {
			trie := New()
			trie.Add("t", 1)
			trie.Add("yo!", "foo")
			trie.Add("sup?", 3)

			println(len(trie.Root.Children))

			trie.Traverse(func(item *node.Node) bool {
				return true
			})
		})
	})
})
