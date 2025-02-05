package problem_208

//
//type TrieR struct {
//	TrieList []string
//}
//
//func Constructor() Trie {
//	treeList := []string{}
//	return Trie{TrieList: treeList}
//}
//
//func (this *Trie) Insert(word string) {
//	this.TrieList = append(this.TrieList, word)
//}
//
//func (this *Trie) Search(word string) bool {
//	for i := range this.TrieList {
//		if this.TrieList[i] == word {
//			return true
//		}
//	}
//	return false
//}
//
//func (this *Trie) StartsWith(prefix string) bool {
//	for i := range this.TrieList {
//		if strings.HasPrefix(this.TrieList[i], prefix) {
//			return true
//		}
//	}
//	return false
//}
//
///**
// * Your Trie object will be instantiated and called as such:
// * obj := Constructor();
// * obj.Insert(word);
// * param_2 := obj.Search(word);
// * param_3 := obj.StartsWith(prefix);
// */
