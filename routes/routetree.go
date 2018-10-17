package routes

// import (
// 	"fmt"
// 	"regexp"
// 	"strings"
// 	. "github.com/flame/controllers"
// )

// type RouteTree struct {
// 	Path string
// 	Regex regexp.Regexp
// 	Parent *RouteTree
// 	Children map[string]*RouteTree
// 	Method BaseController
// 	Params map[string]string
// }

// func NewTree() *RouteTree {
// 	t := new(RouteTree)
// 	t.Children = make(map[string]*RouteTree)
// 	return t
// }

// func (r *RouteTree) AppendChildAndWalk(child *RouteTree) *RouteTree{
// 	child.Parent = r

// 	alreadyContainsChild := false
// 	returnNode := child

// 	if val, ok := r.Children[child.Path]; ok {
// 		alreadyContainsChild = true
// 		returnNode = val
// 	}

// 	if alreadyContainsChild == false {
// 		r.Children[child.Path] = child
// 	}

// 	return returnNode
// }

// func (r RouteTree) IsRoot() bool {
// 	return r.Parent == nil
// }

// func (r RouteTree) IsLeaf() bool {
// 	return len(r.Children) == 0
// }

// func (root *RouteTree) FindRoute(url string) *RouteTree {
// 	split := strings.Split(url, "/")
// 	currentNode := root
// 	var matchedNode *RouteTree
// 	params := make(map[string]string)


// 	if len(split) == 2 && split[0] == "" && split[1] == "" {
// 		return currentNode
// 	} else {
// 		for i := 0; i < len(split); i++ {
// 			path := split[i]
// 			if currentNode.Path == path || strings.Contains(currentNode.Path, ":") {
// 				matchedNode = currentNode
// 				fmt.Println(currentNode.Path)
// 				fmt.Println(strings.Contains(currentNode.Path, ":"))
// 				if strings.Contains(currentNode.Path, ":"){
// 					params[currentNode.Path] = path
// 				}

// 				if (i + 1) < len(split) {
// 					path = split[i + 1]
// 					if val, ok := currentNode.Children[path], ok := currentNode.Children[path];; ok {
// 						currentNode = val
// 					} 
// 				}
// 			} else {
// 				matchedNode = nil
// 				break;
// 			}
// 		}
// 	}	
		
// 	if matchedNode != nil {
// 		matchedNode.Params = params
// 		return matchedNode
// 	}

// 	return nil
// }
