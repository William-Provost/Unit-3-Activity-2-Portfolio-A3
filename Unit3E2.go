// Author: William Provost
// Version: 1.0.0
// Date: 2025-11-08
// Fileoverview: This program will ask the user to enter a length
// in centimetres. It will then convert the length to inches using
// a constant conversion factor (1 inch = 2.54 cm) and display the result.

package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

func main() {
    // constants
    const CM_PER_INCH = 2.54

    // variables
    var lengthCmAsString string
    var lengthCmAsNumber float64
    var lengthInInches float64

    reader := bufio.NewReader(os.Stdin)

    // input
    fmt.Print("Enter the length in centimetres: ")
    lengthCmAsString, _ = reader.ReadString('\n')
    lengthCmAsString = strings.TrimSpace(lengthCmAsString)

    // process
    lengthCmAsNumber, _ = strconv.ParseFloat(lengthCmAsString, 64)
    lengthInInches = lengthCmAsNumber / CM_PER_INCH

    // output
    fmt.Println()
    fmt.Printf("%.2f cm is equal to %.2f inches.\n", lengthCmAsNumber, lengthInInches)
    fmt.Println("Done.")
}
