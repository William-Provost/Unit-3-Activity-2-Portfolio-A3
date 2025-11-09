/**
 * @author William Provost
 * @version 1.0.0
 * @date 2025-11-08
 * @fileoverview This program will ask the user to enter a length
 * in centimetres. It will then convert the length to inches using
 * a constant conversion factor (1 inch = 2.54 cm) and display the result.
 */

// constant
const CM_TO_INCH: number = 2.54;

// variables
let lengthCmAsString: string;
let lengthCmAsNumber: number;
let lengthInches: number;

// input
lengthCmAsString = prompt("Enter the length in centimetres:") || "0";

// process
lengthCmAsNumber = parseInt(lengthCmAsString);
lengthInches = lengthCmAsNumber / CM_TO_INCH;

// output
console.log("\n");
console.log(lengthCmAsNumber + " cm is equal to " + lengthInches + " inches.");
console.log("\nDone.");

