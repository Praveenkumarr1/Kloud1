package main

import (
    "fmt"
    "json/encoding"
)


func main() {
    tree := new(avltree.AVLTree)

    keys := []int{3,2,4,1,5}
    for _, key := range keys {
        tree.Add(key, key*key)
    }   

    tree.Remove(2)
    tree.Update(5, 6, 6*6)
    tree.DisplayInOrder()

    val := tree.Search(3).Value
}