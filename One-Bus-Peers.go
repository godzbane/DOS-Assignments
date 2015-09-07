package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Initialize a single common bus for communication between notes

var bus = make(chan Node, 5)



var fullyFilledNodes = 0

// stepsTaken counts the total number of times a node reads from the bus

var commonBusUsed = 0

var personalBusUsed = 0

var speed = 200 // Inversely proportional to simulation speed. Used in wait random timeout, so keep higher than number of nodes

var noOfNodes = 5 // Change number of nodes here as well as once in Node struct

type Node struct {

	id int
	knowns[5]int // Change number of nodes here
	knownPointer int
	
	nodeChannel chan Node

	nodesThatKnowMe int
}

func main() {

	start := time.Now()

	i := 1

	fmt.Println("Initialized Node with ID's from 1 to", noOfNodes)
	fmt.Println("------------------")

	// Initialize all the nodes:
	
	for i < (noOfNodes + 1) {

		go NodeProcess(Node{id: i, nodeChannel: make (chan Node)})
		i++
	}
	
	// Just keep running until all the nodes know each other:

	for {

		if fullyFilledNodes >= noOfNodes {
			fmt.Printf("\n------------------\nAll nodes now know each other. \n\nCommon Bus Requests: %d\nPersonal Bus Requests: %d\n", commonBusUsed, personalBusUsed)
			elapsed := time.Since(start)
			fmt.Println("Time Taken: ", elapsed)
			break
		}

	}

}

	/* Each instance of NodeProcess called functions as a virtual node. Each node has a number of known peers (0 others in the beginning).
		 For simplicity here, knowing a peer is just marking down it's ID instead of its reference	*/

func NodeProcess(n Node) {

	n.knowns[0] = n.id

	for {

		select {
		
		// Receive incoming request from Node personal channel containing the ID of the other node
		
		case gotNode:= <- n.nodeChannel:
			
			personalBusUsed++
			
			// Comment out the next line for ease of reading output
			fmt.Printf ("Node %d: Received reply via personal bus from Node %d.\n", n.id, gotNode.id)
			
			i := 0

			exists := 0

			for i <= n.knownPointer {

				if n.knowns[i] == gotNode.id {
					exists = 1
				}
				i++
			}
			
			// If node does not know the other node, add it to own address

			if exists == 0 {

				n.knownPointer++
				n.knowns[n.knownPointer] = gotNode.id
				
				fmt.Printf ("Updated node %d with list of known nodes: %v\n", n.id, n.knowns)

				gotNode.nodesThatKnowMe++

				if n.knownPointer >= (noOfNodes - 1) {
					fullyFilledNodes++

					fmt.Printf("Node %d address-book filled.\n", n.id)
				}

			}
			
			break
		
		// Receive incoming request from bus containing the ID of the other node

		case gotNode := <-bus:
		
		if (n.id == gotNode.id) {
		
		bus <- n
		
		time.Sleep(time.Duration(random(speed/2, 2*speed)) * time.Millisecond)
		
		} else {
		
		// Comment out the next line for ease of reading output
		fmt.Printf ("Node %d: Received request via common bus from Node %d.\n", n.id, gotNode.id)
		
		commonBusUsed++

			i := 0

			exists := 0

			for i <= n.knownPointer {

				if n.knowns[i] == gotNode.id {
					exists = 1
				}
				i++
			}
			
			// If node does not know the other node, add it to own address

			if exists == 0 {

				n.knownPointer++
				n.knowns[n.knownPointer] = gotNode.id
				
				fmt.Print ("Updated node ", n.id)
				fmt.Println (" with list of known nodes:  ", n.knowns)

				gotNode.nodesThatKnowMe++

				if n.knownPointer >= (noOfNodes - 1) {
					fullyFilledNodes++

					fmt.Printf("Node %d address-book filled.\n", n.id)
				}

			}
			
			
			// Send self Node data to other node if other node does not have me in its address book
			
			j := 0

			otherExists := 0

			for j <= gotNode.knownPointer {
				if gotNode.knowns[j] == n.id {
					otherExists = 1
				}
				j++
			}
			
			// If node does not know the other node, add it to own address

			if otherExists == 0 {

				// Comment out the next line to simulate without replies on personal channels
				gotNode.nodeChannel <- n

			}
			
			}
			
			break
			
			

		case <-time.After(time.Duration(random(speed/2, 2*speed)) * time.Millisecond):

			// Stop broadcasting if everyone knows about this node

			if n.nodesThatKnowMe < (noOfNodes - 1) {
			fmt.Printf ("Node %d: Placing request in bus.\n", n.id)
				bus <- n
			} 

			break

		}

	}

}


func random(min, max int) int {
rand.Seed( time.Now().UTC().UnixNano())
return rand.Intn(max - min) + min
}
