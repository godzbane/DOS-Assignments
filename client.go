package main 
import "net" 
import "fmt"
import "bufio" 
import "os"


func main() {  
	
	conn, _ := net.Dial("tcp", "127.0.0.1:8081")   

	fmt.Print("Please enter a BODMAS type query. Example: (6^2)/((12-4)+1)\n Functions allowed: ^ / * + - ()\nNote: Performs integer division with integer output (eg 2/3 = 0).\n")
	
	for {

		reader := bufio.NewReader(os.Stdin)     
		fmt.Print("Input query: ")     
		
		text, _ := reader.ReadString('\n')    
		fmt.Fprintf(conn, text + "\n")     
		
		message, _ := bufio.NewReader(conn).ReadString('\n')     
		fmt.Print("Response from calculation server: "+message)
	} 
}

