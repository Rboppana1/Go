/*Write a program which reads information from a file and represents it in a slice of structs. Assume that there is a text file which contains a series of names. Each line of the text file has a first name and a last name, in that order, separated by a single space on the line.

Your program will define a name struct which has two fields, fname for the first name, and lname for the last name. Each field will be a string of size 20 (characters).

Your program should prompt the user for the name of the text file. Your program will successively read each line of the text file and create a
struct which contains the first and last names found in the file. Each struct created will be added to a slice, and after all lines have been read from the file,
your program will have a slice containing one struct for each line in the file. After reading all lines from the file, your program should iterate through your slice
of structs and print the first and last names found in each struct.*/

package main

import (
   
    "fmt"
    "io/ioutil"
    "log"
  
)

type Name struct {

fname string
lname string
}

func main() {
    fmt.Println("Enter filename: ")
    var filename string
    fmt.Scanln(&filename)
 
 
    ReadFile(filename)
}
func ReadFile(filename string) {
 
    fmt.Printf("\n\nReading a file in Go lang\n")
     
   
    data, err := ioutil.ReadFile(filename)
     
  
    if err != nil {
        log.Panicf("failed reading data from file: %s", err)
    }
    fmt.Printf("\nFile Name: %s", filename)
    fmt.Printf("\nData: %s", data)
	slice := make([]Name, 0, 3)
	var namee Name
	slice = append(slice, namee)
fmt.Println("\n********************\n")

for _, v := range slice {
	fmt.Println(v.fname, v.lname)
}
 
}