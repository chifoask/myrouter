package myrouter

import (
	"net/http"
)

// ノードを表す構造体を定義する
type Node struct {
	isRoot    bool
	character byte
	children  []*Node
	handlers  map[string]http.Handler
}

func newNode(character byte) *Node {
	return &Node{
		character: character,
		children:  []*Node{},
		handlers:  make(map[string]http.Handler),
	}
}

// Routerの構造体を定義する
type Router struct {
	tree *Node
}

func NewRouter() *Router {
	return &Router{
		tree: &Node{
			isRoot:   true,
			handlers: nil,
		},
	}
}

// Routerにinsertメソッドを実装する
func (r *Router) insert(method, endpoint string, handler http.Handler) {
	currentNode := r.tree

	for i := 0; i < len(endpoint); i++ {
		target := endpoint[i]

		nextNode := currentNode.nextChild(target)
		if nextNode == nil {
			node := newNode(target)
			currentNode.children = append(currentNode.children, node)
			currentNode = node
			continue
		}

		currentNode = nextNode
	}

	currentNode.handlers[method] = handler
}

// NextChildメソッドをNode構造体に実装する
func (n *Node) nextChild(character byte) *Node {
	for _, child := range n.children {
		if child.character == character {
			return child
		}
	}

	return nil
}

// RouterにGETメソッドを実装する
func (r *Router) GET(endpoint string, handler http.Handler) {
	r.insert(http.MethodGet, endpoint, handler)
}

// Routerにsearchメソッドを実装する
func (r *Router) Search(method, endpoint string) http.Handler {
	currentNode := r.tree
	IcpIndex := 0

	for {
		nextNode := currentNode.nextChild(endpoint[IcpIndex])
		if nextNode == nil {
			return nil
		}

		// 各ノードの文字数は1文字と限定されているため、
		// IcpIndexをインクリメントするだけで良い
		IcpIndex++
		currentNode = nextNode
		if IcpIndex == len(endpoint) {
			break
		}
	}
	return currentNode.handlers[method]
}

// RouterにServeHTTPメソッドを実装する
func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	handler := r.Search(req.Method, req.URL.Path)
	if handler != nil {
		handler.ServeHTTP(w, req)
		return
	}

	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("404 Not Found"))
	// return
}
