package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func read_input(file_name string) string {
	result := []string{}
	file, err := os.Open(file_name)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		result = append(result, line)
	}
	return result[0]
}

type Node struct {
	id     int
	index  int
	length int
}

func swapElm(array []Node, a int, b int) []Node {
	if a > b {
		a, b = b, a
	}
	aNode := array[a]
	bNode := array[b]
	if aNode.length != bNode.length {
		log.Fatal("Length not match")
	}
	aNode.index, bNode.index = bNode.index, aNode.index
	// From lastNode -> end
	temp_last := append([]Node{aNode}, array[b+1:]...)
	// From firstNode -> lastNode
	next_last := append([]Node{bNode}, append(array[a+1:b], temp_last...)...)
	// From start -> firstNode
	array = append(array[:a], next_last...)
	return array
}

func cleanUp(array []Node) []Node {
	lastIndex := len(array) - 1
	lastNode := array[lastIndex]
	total_length := 0
	for {
		total_length += lastNode.length
		newLastIndex := lastIndex - 1
		newLastNode := array[newLastIndex]
		if newLastNode.id != -1 {
			if total_length == 0 {
				return array
			} else {
				return append(array[:lastIndex], Node{-1, lastNode.index, total_length})
			}
		}
		lastIndex = newLastIndex
		lastNode = newLastNode
	}
}

func checkSum(array []Node) int {
	result := 0
	for _, node := range array {
		if node.id == -1 {
			continue
		}
		for i := node.index; i < (node.index + node.length); i++ {
			result += node.id * i
		}
	}
	return result
}

func main01() {
	// file_name := "input.txt"
	file_name := "real_input.txt"
	input := read_input(file_name)
	// fmt.Println(input)

	frag_map := []Node{}
	current_index := 0
	for i, r := range input {
		// fmt.Println(i, string(r))
		thisSize, _ := strconv.Atoi(string(r))
		var thisNode Node
		if i%2 == 0 {
			thisNode = Node{id: i / 2, index: current_index, length: thisSize}
		} else {
			thisNode = Node{id: -1, index: current_index, length: thisSize}
		}
		current_index += thisSize
		frag_map = append(frag_map, thisNode)
	}

	// for _, node := range frag_map {
	// 	fmt.Println(node)
	// }

	isDone := false
	for { // Get last non -1 node
		var lastNode Node
		lastIndex := len(frag_map) - 1
		for {
			lastNode = frag_map[lastIndex]
			if lastNode.id != -1 {
				break
			}
			lastIndex--
		}
		// Get first -1 node
		var firstNode Node
		firstIndex := 0
		for {
			firstNode = frag_map[firstIndex]
			if firstNode.id == -1 {
				break
			}
			firstIndex++
		}
		// fmt.Println("Last non -1 is ", lastNode, "\nFirst -1 is ", firstNode)

		if lastNode.index < firstNode.index {
			// All -1 nodes are behind non -1 nodes
			isDone = true
			break
		}

		if firstNode.length < lastNode.length {
			// Need to split lastNode into 2 nodes
			firstPart := Node{lastNode.id, lastNode.index, lastNode.length - firstNode.length}
			secondPart := Node{lastNode.id, lastNode.index + firstPart.length, firstNode.length}
			// Update frag_map
			frag_map = append(frag_map[:lastIndex], append([]Node{firstPart, secondPart}, frag_map[lastIndex+1:]...)...)
			// Swap
			lastIndex = lastIndex + 1
		} else if firstNode.length > lastNode.length {
			// Need to split firstNode into 2 nodes
			firstPart := Node{firstNode.id, firstNode.index, lastNode.length}
			secondPart := Node{firstNode.id, firstNode.index + firstPart.length, firstNode.length - lastNode.length}
			// Update frag_map
			frag_map = append(frag_map[:firstIndex], append([]Node{firstPart, secondPart}, frag_map[firstIndex+1:]...)...)
			// Swap
			lastIndex = lastIndex + 1
		}
		frag_map = swapElm(frag_map, firstIndex, lastIndex)
		// fmt.Println("After move", frag_map)

		frag_map = cleanUp(frag_map)
		// fmt.Println("After cleanup", frag_map)

		if isDone {
			break
		}
	}

	// fmt.Println("The result", frag_map)
	fmt.Println(checkSum(frag_map))
}

func main02() {
	// file_name := "input.txt"
	file_name := "real_input.txt"
	input := read_input(file_name)
	// fmt.Println(input)

	frag_map := []Node{}
	current_index := 0
	for i, r := range input {
		// fmt.Println(i, string(r))
		thisSize, _ := strconv.Atoi(string(r))
		var thisNode Node
		if i%2 == 0 {
			thisNode = Node{id: i / 2, index: current_index, length: thisSize}
		} else {
			thisNode = Node{id: -1, index: current_index, length: thisSize}
		}
		current_index += thisSize
		frag_map = append(frag_map, thisNode)
	}

	nodeIndex := len(frag_map) - 1
	for {
		if nodeIndex < 0 {
			break
		}
		node := frag_map[nodeIndex]
		if node.id == -1 {
			nodeIndex--
			continue
		}

		// find the first -1 that can hold this node
		found := false
		var firstNode Node
		var firstIndex int
		for firstIndex = 0; firstIndex < nodeIndex; firstIndex++ {
			firstNode = frag_map[firstIndex]
			if firstNode.id != -1 {
				continue
			}
			if firstNode.length >= node.length {
				found = true
				break
			}
		}

		if !found {
			nodeIndex--
			continue
		}

		// fmt.Println("Swapping", firstNode, node)
		if firstNode.length > node.length {
			// Need to split firstNode into 2 nodes
			firstPart := Node{firstNode.id, firstNode.index, node.length}
			secondPart := Node{firstNode.id, firstNode.index + firstPart.length, firstNode.length - node.length}
			// Update frag_map
			frag_map = append(frag_map[:firstIndex], append([]Node{firstPart, secondPart}, frag_map[firstIndex+1:]...)...)
			// Swap
			nodeIndex++
		}
		frag_map = swapElm(frag_map, firstIndex, nodeIndex)
		// fmt.Println("After move", frag_map)
		frag_map = cleanUp(frag_map)
		// fmt.Println("After cleanup", frag_map)

		nodeIndex--
	}
	// fmt.Println("The result", frag_map)
	fmt.Println(checkSum(frag_map))
}

func main() {
	main01()
	main02()
}
