type TrieNode struct {
  val rune
  children []*TrieNode
  isEnd bool
}

func NewTrieNode(val rune) *TrieNode {
  t := &TrieNode {
    val: val,
    children: make([]*TrieNode, 26),
    isEnd: false,
  }
  return t
}

type Trie struct {
    root *TrieNode
}


/** Initialize your data structure here. */
func Constructor() Trie {
  t := Trie {
    root: NewTrieNode('/'),
  }
  return t
}


/** Inserts a word into the trie. */
func (this *Trie) Insert(word string)  {
  p := this.root
  for _, c := range word {
    idx := int(c - 'a')
    if p.children[idx] == nil {
      p.children[idx] = NewTrieNode(c)
    }
    p = p.children[idx]
  }
  p.isEnd = true
}


/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
  p := this.root
  for _, c := range word {
    idx := int(c - 'a')
    if p.children[idx] == nil {
      return false
    }
    p = p.children[idx]
  }
  return p.isEnd
}


/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
  p := this.root
  for _, c := range prefix {
    idx := int(c - 'a')
    if p.children[idx] == nil {
      return false
    }
    p = p.children[idx]
  }
  return true
}


/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */