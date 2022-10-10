/*
QUESTION:
Write a program which prompts the user to enter a floating point number and prints the integer
which is a truncated version of the floating point number that was entered. 
Truncation is the process of removing the digits to the right of the decimal place.
This code is my solution for a coursera problem.
*/
package main

import ( // imports libraries 
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

)

type Result struct { //creating a struct for use with a channel 
	Month		string	`json:"month"`
	Num		int	`json:"num"`
	Link		string	`json:"link"`
	Year		string	`json:"year"`
	News		string	`json:"news"`
	SafeTitle	string `json:"safe_title"`
	Transcript	string	`json:"transcript"`
	Alt		string	`json:"alt"`
	Img		string	`json:"img"`
	Title		string	`json:"title"`
	Day		string	`json:"day"`
}

const Url = "https://xkcd.com" //storing url in global variable 

func fetch(n int) (*Result, error) { 

	client := &http.Client{
		Timeout: 5 * time.Minute, //setting timeout for the client
	}

	// concatenate strings to get url; ex: https://xkcd.com/571/info.0.json
	url := strings.Join([]string{Url, fmt.Sprintf("%d", n), "info.0.json"}, "/")

	req, err := http.NewRequest("GET", url, nil) //get request and accompanying error message 

	if err != nil { //if error not equal to nil 
		return nil, fmt.Errorf("http request: %v", err)
	}

	resp, err := client.Do(req) //client executes request 
	if err != nil { //if error nopt equal to nil 
		return nil, fmt.Errorf("http err: %v", err)
	}
	
	var data Result //setting the result as a local variable 
	
	// error from web service, empty struct to avoid disruption of process
	if resp.StatusCode != http.StatusOK {
		data = Result{
		}
		} else {
			if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
				return nil, fmt.Errorf("json err: %v", err) //returning error in format 
			}
		}
		
	resp.Body.Close() //closing the body 

	return &data, nil //returning the result
}

func main() {
	n := 200 //assigning a value to "n" 
	result, err := fetch(n) //calling the funtion to execute providing n as parameter 
	if err != nil {
		log.Fatal(err) //logging the error if it is fatal 
	}
	fmt.Printf("%v\n", result.Title) //printing the result 
}