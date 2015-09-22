package main 

import "net" 
import "fmt" 
import "bufio" 
import "strings" 
import "strconv"

var errors = 0

func main() {   

	fmt.Println("Calculation server ready. Please run client.")   

	ln, _ := net.Listen("tcp", ":8081")   
	conn, _ := ln.Accept()   

	for {
		message, _ := bufio.NewReader(conn).ReadString('\n')     

		if (len (message) > 0) {
			fmt.Println("Query received:", string(message))
		}
		
		response := FindAns(message)     
		
	

		conn.Write([]byte(response + "\n"))
		
		fmt.Println("Query resolved and returned to client. Errors found: ", errors)

		errors = 0
	} 
}

func FindAns (str string) string {

for (strings.Contains(str, "/") || strings.Contains(str, "*") || strings.Contains(str, "+") || strings.Contains(str, "-") || strings.Contains(str, "^")){
	str = calc (str)
	}
	return str
	
	
}

func calc (str string) string {

	if (strings.Contains(str, "(")) {
		for strings.Contains(str, "(") {
		
		
			lastOpen := strings.LastIndex(str,"(")
			nextClose := strings.Index (str[lastOpen:], ")") + lastOpen
			
			if (nextClose == -1) {
				errors = errors + 1
				return ""
			}
				
			str = str[:lastOpen] + calc (str[lastOpen+1:nextClose]) + str[nextClose+1:]
			
		}
	} else {
	
		for strings.Contains(str, "^") {
			str = doOperation (str, "^")
			}
	
		for strings.Contains(str, "/") {
			str = doOperation (str, "/")
			}
		
		for strings.Contains(str, "*") {
			str = doOperation (str, "*")
			}
		
		for strings.Contains(str, "+") {
			str = doOperation (str, "+")
			}
		
		for strings.Contains(str, "-") {
			str = doOperation (str, "-")
			}
	}
return str
}


func doOperation (str, operator string) string {

			a := 0
			aLen := 1
		
			for i := 0; i < strings.Index(str,operator); i++ {
				if z, err := strconv.Atoi(str[i:strings.Index(str,operator)]); err == nil {

					a = z
					aLen = strings.Index(str,operator)-i
					break
				}
			}
			
			b := 1
			bLen := 1
		
			for i := (strings.Index(str,operator)+1); i <= len(str); i++ {

				if z, err := strconv.Atoi(str[(strings.Index(str,operator)+1):i]); err == nil {
					
					b = z
					bLen = i-(strings.Index(str,operator))
					break
				}
			}
			
			value := 1
			
			if (operator == "/") {
				if (b != 0){
					value = a/b
				} else {
					errors = errors + 1
				}
			}	else if (operator == "*") {
			value = a*b
			} else if (operator == "+") {
			value = a+b
			} else if (operator == "-") {
			value = a-b
			} else if (operator == "^") {
				for i := 1; i <= b; i++ {
				value = value*a
				}
			}
			
			str = str[:strings.Index(str,operator)-aLen] + strconv.Itoa(value) + str[strings.Index(str,operator)+bLen:]
			
			return str

}