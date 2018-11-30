package main                                                                                                             
                                                                                                                         
import (                                                                                                                 
        "fmt"         // *Println, Fprintln                                                                              
        "net/http"    // *HandleFunc, ListenAndServe                                                                     
        "crypto/rand" // Generating secure numbers                                                                       
        "math/big"    // This is used with >crypto/rand                                                                  
        "strings"     // Splitting UserAgent                                                                             
        "log"         // *Fatal                                                                                          
)                                                                                                                        
                                                                                                                         
var requestCount int                                                                                                     
                                                                                                                         
func handler(w http.ResponseWriter, r *http.Request) {                                                                   
        // Generate a crptographically-secure LUID                                                                       
        // Defines the maximum value                                                                                     
                                                                                                                         
        max := big.NewInt(65536)                                                                                         
                                                                                                                         
        // Start off with a blank string to append the text to                                                           
        result := ""                                                                                                     
                                                                                                                         
        // Add to <result> 256 times with a crypto/rand Int                                                              
        for i := 0; i < 256; i++ {                                                                                       
                // Create a random integer                                                                               
                // This returns big.Int for some reason                                                                  
                n, _ := rand.Int(rand.Reader, max)                                                                       
                                                                                                                         
                // Convert it to an int that string() can parse                                                          
                conv := n.Int64()                                                                                        
                                                                                                                         
                // Convert the integer to a rune and store it in                                                         
                // the result                                                                                            
                result += string(conv)                                                                                   
        }                                                                                                                
                                                                                                                         
        // Increase the request counter by one                                                                           
        requestCount += 1                                                                                                
                                                                                                                         
        // This information is used to determine whether or not a                                                        
        // user is using curl to grab information. If the user is,                                                       
        // Then return the raw text, but if not, then surround it                                                        
        // with HTML <p></p> tags so it looks better.                                                                    
        ua := r.UserAgent() // Grab their user agent                                                                     
        split := strings.Split(ua, "/") // Array of UA info                                                              
        if split[0] == "curl" {                                                                                          
                // Just hand them the raw info                                                                           
                fmt.Fprintln(w, result)                                                                                  
        } else {                                                                                                         
                fmt.Fprintln(w, "<p>", result, "</p>")                                                                   
        }                                                                                                                
                                                                                                                         
        // Log the request IP, Port, User Agent, and request counter                                                     
        fmt.Printf("LUID Call: %-7d %-25s %s\n", requestCount, r.RemoteAddr, r.UserAgent())                              
        // fmt.Println("Request from", r.RemoteAddr, "         ", r.UserAgent(), requestCount)                           
}                                                                                                                        
                                                                                                                         
func handleepic(w http.ResponseWriter, r *http.Request) {                                                                
        fmt.Fprintf(w, "welcome to epic gamer games")                                                                    
}                                                                                                                        
                                                                                                                         
func main() {                                                                                                            
        fmt.Println("Starting...")                                                                                       
        http.HandleFunc("/", handler)                                                                                    
        log.Fatal(http.ListenAndServe(":80", nil))                                                                       
} 
