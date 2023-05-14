// The following codebase will show you how to create a binary tree and how to traverse it in order to print all of its elements 
// as proof that Go can be used for creating a tree data structure. Therefore, it will not implement the full functionality of a binary tree, 
// which also includes deleting a tree node and balancing a tree.


package main

import (
	"fmt"
	"math/rand"
	"time"
)


// Defining the structure of a binary tree node . 
// The `Tree` struct contains 3 fields 
	type Tree struct { 
	Left *Tree // represents the left child of the node . 
		Value int  // stores the integer value of the node . 
	Right *Tree // represents the right child of the node . 
}



// The second part of the program contains functions allowing you to traverse a tree in order to 
// print all of its elements, create a tree with randomly generated numbers , 
// and insert a node into it:

//the `traverse` function is a recursive function  , meaning it is a function that calls itself during execution. 
// The function perfoms an inorder traversal of the binary tree and pritn the values of each node.
// An inorder traversal is a way of visiting nodes in a binary tree. 
// For each node, the `traverse` function first vists the left sub-tree, tehn nprints the value of the current node, 
// and finally visits the right sub-tree. 

// By following this process recursively for each node, the `traverse` function ensures that all nodes 
// are visited in the correct order . 

func traverse(t *Tree) { 
	// If the current node  `t` is `nil ` , , it means that we have reached the end of a branch, so the fucntion returns.   
			if t == nil { 
			return 

		}
		// The function calls itself recursively to traverse the left sub-tree. 
		traverse(t.Left) 
		// Printing the value of the current node followed by a space .  
		fmt.Print(t.Value, " ") 
		// Calling itself recursively to traverse the right sub-tree. 
		traverse(t.Right) 


		}
// The `create` function generates a binary tree with random values, and takes two parameters: 
// `n ` which is an integer representing the range of random number, 
// and `seed`,  which  is also an integer used as the seed for random numbers generation)
func create(n int, seed int64 ) *Tree { 
	// The function initializes a var `t` as a pointer to a `Tree` structure. 
	var t *Tree


	//creates a new random generator r using the provided seed value. 
	//This ensures that the random numbers generated will have a specific result sequence based on the given seed .  
	r := rand.New(rand.NewSource(seed))

	// The function iterates from 0 to `2*n` , generating a random value `temp` . 
	for  i:=0; i <2*n ; i++ { 
		temp := r.Intn(n)  
		// The function calls the `insert ` function (below) to insert the value `temp ` into the binary tree `t`. 
		t = insert(t, temp) 
	} 
  // The function finally, returns the generated binary tree . 
	return t 
}




// The insert function inserts a new node with value v into the binary tree t.
func insert(t *Tree, v int) *Tree {  

// If the tree is empty (t == nil), it creates a new node with v as the value and returns a pointer to it.
	if t ==nil { 
		return&Tree{nil, v, nil }
	}

// If the value v is already present in the tree (v == t.Value), it returns the current node t without making any changes.
	if v == t.Value { 
		return t 
	} 

// If v is less than the value of the current node (v < t.Value), it recursively inserts v into the left subtree.
	if v <t.Value { 
		t.Left = insert(t.Left, v ) 
		return t 
	}


//Otherwise, it recursively inserts v into the right subtree
	t.Right = insert(t.Right, v ) 
	return t 
}







func main() {
	//assigning the current Unix timestamp value to the seed variable. This generates a unique seed value based on the current time.
	seed := time.Now().Unix()  

	//passing 30 as the range of random numbers and seed as the seed value for random number generation.
	tree := create(30, seed)
	traverse(tree) 
	fmt.Println() 
		fmt.Println("The value of the root of the tree is", tree.Value)
}