package DataStructures

type Node struct {
	ch       rune
	isValid  bool
	count    int
	children [26]*Node
}

func (trie *Node) insert(name string) {
	trie.count = trie.count + 1
	tempNode := trie
	for index, value := range name {
		characterIndex := name[index] - 97
		currentNode := tempNode.children[characterIndex]
		if currentNode == nil {
			currentNode = &Node{ch: value}
		}
		isValid := len(name) == index-1
		currentNode.isValid = isValid
		currentNode.count = currentNode.count + 1
		tempNode.children[characterIndex] = currentNode
		tempNode = tempNode.children[characterIndex]
	}
}

func (trie *Node) initialize() {
	trie.count = 0
	trie.isValid = false
	trie.ch = '-'
}

func (trie *Node) numberOfPrefix(subString string) int {
	if subString == "-" {
		return trie.count
	}
	if len(subString) < 1 {
		return 0
	}
	tempNode := trie
	for index, _ := range subString {
		characterIndex := subString[index] - 97
		currentNode := tempNode.children[characterIndex]
		if currentNode == nil {
			return 0
		}
		tempNode = currentNode
	}
	return tempNode.count
}

func init() {
	trie := Node{}
	trie.initialize()

	/*trie.insert("balaji")
	fmt.Println("Trie : %v", trie)
	fmt.Println("Count: %v", trie.numberOfPrefix("-"))
	fmt.Println("Coutn of bal : %v", trie.numberOfPrefix("bal"))*/

}
