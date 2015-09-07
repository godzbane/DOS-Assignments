package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Initialize a single bus for communication between notes

//var bus = make(chan Node)

var noOfNodes = 10

var fullyFilledNodes = 0

var stepsTaken = 0

type Node struct {

	id int
	knowns[10]int
	knownPointer int
	
	nodeChannel chan Node
	
	channelList[10] chan Node

	nodesThatKnowMe int
}

func main() {

	start := time.Now()

	i := 0

	fmt.Println("Initialized Node with ID's from 1 to", noOfNodes)
	fmt.Println("------------------")

	var r [10] Node
	
	for i < (noOfNodes) {
		
		r[i] = (Node{id: (i+1), nodeChannel: make(chan Node)})
		i++
	}
	
	// Give each node its own channel and point it to the previous channel to make a circle
	
	r[0].knowns[0] = 1
	r[0].knowns[1] = noOfNodes
	r[0].channelList[0] = r[0].nodeChannel
	r[0].channelList[1] = r[noOfNodes-1].nodeChannel
	
	i = 1
	
	for i < (noOfNodes) {
		
		r[i].knowns[0] = i
		r[i].knowns[1] = r[i-1].id-1
		r[i].channelList[0] = r[i].nodeChannel
		r[i].channelList[1] = r[i-1].nodeChannel
	
		go NodeProcess(r[i])
		i++
	}

	for {

		if fullyFilledNodes >= noOfNodes {
			fmt.Println("------------------\nAll nodes now know each other. \n\nSteps taken: ", stepsTaken)
			elapsed := time.Since(start)
			fmt.Println("Time Taken: ", elapsed)
			break
		}

	}

}

func NodeProcess(n Node) {
	
		
	n.knownPointer = 2
	
	
	// Infinite loop emulating nodes of a peer to peer network

	for {
	
		select {

		case gotNode := <- n.channelList[random (0, n.knownPointer)]:

			i := 0

			exists := 0

			for i < n.knownPointer {

				if n.knowns[i] == gotNode.id {
					exists = 1
				}
				i++
			}

			if exists == 0 {

				stepsTaken++
				
				n.knowns[n.knownPointer] = gotNode.id
				n.channelList[n.knownPointer] = gotNode.nodeChannel
				n.knownPointer++
				// Uncomment the following lines to view the progress of the graph filling
				
				fmt.Print ("\nUpdated node ", n.id)
				fmt.Println (" with list of known nodes:  ", n.knowns)

				gotNode.nodesThatKnowMe++

				if n.knownPointer >= noOfNodes {
					fullyFilledNodes++

					fmt.Printf("Node %d address-book filled. Total filled nodes: %d\n", n.id, fullyFilledNodes)
				}

			} 

			break

		case <-time.After(time.Duration(random(0,50)) * time.Microsecond):

			// Pick a channel to send self data
		
			// Stop broadcasting if everyone knows about this node

			if n.nodesThatKnowMe < (noOfNodes - 1) {
						
				n.channelList[random (1, n.knownPointer)] <- n
				
			}

			break
		
		}

	}

}

func random(min, max int) int {
rand.Seed( time.Now().UTC().UnixNano())
return rand.Intn(max - min) + min
}


