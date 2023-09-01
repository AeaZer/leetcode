package dictionaries

type Trie struct {
	children [26]*Trie
	isEnd bool
}


func Constructor() Trie {
	return Trie{}
}

func (this *Trie) Insert(word string)  {
	dummyNode := this
    for _, ch := range word {
        ch -= 'a'
        if dummyNode.children[ch] == nil {
            dummyNode.children[ch] = &Trie{}
        }
        dummyNode = dummyNode.children[ch]
    }
    dummyNode.isEnd = true
}

func (this *Trie) SearchNode(word string) *Trie {
	dummyNode := this
	wordLength := len(word)
	for i := 0; i < wordLength; i++ {
		ch := word[i] - 'a'
		if dummyNode.children[ch] == nil {
			return nil
		}
		dummyNode = dummyNode.children[ch]
	}
	return dummyNode
}


func (this *Trie) Search(word string) bool {
	node := this.SearchNode(word)
	return node != nil && node.isEnd
}


func (this *Trie) StartsWith(prefix string) bool {
	return this.SearchNode(prefix) != nil
}
