package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Initialize a single bus for communication between notes

var bus = make(chan Node)

var noOfNodes = 5

var fullyFilledNodes = 0

var stepsTaken = 0

type Node struct {

	id int
	knowns[5]int
	knownPointer int

	nodesThatKnowMe int
}

func main() {

	start := time.Now()

	i := 1

	fmt.Println("Initialized Node with ID's from 1 to", noOfNodes)
	fmt.Println("------------------")

	for i < (noOfNodes + 1) {

		go NodeProcess(Node{id: i})
		i++
	}

	for {

		if fullyFilledNodes >= noOfNodes {
			fmt.Println("------------------\nAll nodes now know each other. \n\nMessages read: ", stepsTaken)
			elapsed := time.Since(start)
			fmt.Println("Time Taken: ", elapsed)
			break
		}

	}

}

func NodeProcess(n Node) {

	n.knowns[0] = n.id

	// Infinite loop emulating nodes of a peer to peer network

	for {

		select {

		case gotNode := <-bus:
		
		stepsTaken++

			i := 0

			exists := 0

			for i <= n.knownPointer {

				if n.knowns[i] == gotNode.id {
					exists = 1
				}
				i++
			}

			if exists == 0 {

				

				n.knownPointer++
				n.knowns[n.knownPointer] = gotNode.id

				// Uncomment the following lines to view the progress of the graph filling
				
				fmt.Print ("Updated node ", n.id)
				fmt.Println (" with list of known nodes:  ", n.knowns)

				gotNode.nodesThatKnowMe++

				if n.knownPointer >= (noOfNodes - 1) {
					fullyFilledNodes++

					fmt.Printf("Node %d address-book filled. Total filled nodes: %d\n", n.id, fullyFilledNodes)
				}

			}

			break

		case <-time.After(time.Duration(rand.Int31n(500)) * time.Microsecond):

			// Stop broadcasting if everyone knows about this node

			if n.nodesThatKnowMe < (noOfNodes - 1) {
				bus <- n
			}

			break

		}

	}

}
