package main

import (
	"errors"
	"flag"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func validateArguments(args []string) error {
	var err error
	if len(args) < 1 {
		flag.Usage()
		err = errors.New("Invalid input")
	}
	return err
}

func calculateOperation(args []string) (float64, string, error) {
	sum := 0.0
	tempNum := 0.0
	var err error
	operand := ""
	operation := ""
	expectNumber := true
	for k, v := range args {
		if err != nil {
			//If we found an error in the switch statement we need to break the for loop as well
			break
		}
		switch v {
		case "*":
			if expectNumber {
				err = errors.New("Invalid input")
				break
			}
			expectNumber = true
			operand = "*"
			operation += "* "
		case "-":
			if expectNumber {
				err = errors.New("Invalid input")
				break
			}
			expectNumber = true
			operand = "-"
			operation += "- "
		case "+":
			if expectNumber {
				err = errors.New("Invalid input")
				break
			}
			expectNumber = true
			operand = "+"
			operation += "+ "
		case "/":
			if expectNumber {
				err = errors.New("Invalid input")
				break
			}
			expectNumber = true
			operand = "/"
			operation += "/ "
		default:
			if !expectNumber {
				err = errors.New("Invalid input")
				break
			}
			expectNumber = false
			tempNum, err = convertToFloat64(v)
			if err != nil {
				break
			}
			operation += v + " "
			if k == 0 {
				sum = tempNum
			} else if tempNum > 0 {
				switch operand {
				case "*":
					sum *= tempNum
				case "-":
					sum -= tempNum
				case "+":
					sum += tempNum
				case "/":
					sum /= tempNum
				}
			}
		}
	}

	if expectNumber {
		err = errors.New("Invalid input")
	}

	return sum, operation, err
}

func convertToFraction(sum float64) string {
	//Round the sum to two decimals
	sum = math.Round(sum*100) / 100

	negativeNumber := false
	//If the number is negegative, convert the number to positive first
	if sum < 0 {
		negativeNumber = true
		sum *= -1
	}

	//Convert to the 100 to remove decimal places
	sumAsInt := int(sum * 100)
	//Get fraction by returning the greatest common denominator
	gcf := getGreatestCommonDenominator(sumAsInt, 100)
	fraction := fmt.Sprintf("%d/%d", (sumAsInt / gcf), (100 / gcf))

	//Convert back to negative (in case that was the case)
	if negativeNumber {
		fraction = "-" + fraction
	}
	return fraction
}

func main() {
	runProgram(os.Args[1:])
}

func runProgram(args []string) {
	flag.Usage = func() {
		fmt.Printf("Usage: %s [operations] \n", os.Args[0])
		fmt.Printf("Example: 2_3/4 + 1/2 - 7/2 * 4 / 1/8")
	}
	err := validateArguments(args)
	if err != nil {
		fmt.Printf("Invalid Operation \n")
		os.Exit(-1)
	} else {
		sum, operation, err := calculateOperation(args)
		if err != nil {
			fmt.Printf("Invalid Operation \n")
			os.Exit(-1)
		}
		fmt.Printf("Operation: %s = %0.2f \n", operation, sum)
		fmt.Println(fmt.Sprintf("Fraction of %0.2f: %s", sum, convertToFraction(sum)))
	}

}

func getGreatestCommonDenominator(number1 int, number2 int) int {
	if number1 == 0 {
		return number2
	}
	return getGreatestCommonDenominator(number2%number1, number1)

}

func convertToFloat64(value string) (float64, error) {

	if value == " " || len(value) == 0 {
		return 0.0, nil
	}

	//Possible Scenarios
	// Whole number: 1  15   101
	// Fraction:  1/4  2/7   1/8
	// Mixer Number:   3_1/4    10_3/4   2_1/2
	// Improper Fractoin: 9/8   11/4

	//First determine if its an full number, a fraction, an improper fraction or a mixed number
	charArray := strings.Split(value, "")
	whole := ""
	numerator := ""
	denominator := ""
	tempNum := ""
	expectNumber := false
	for i := 0; i < len(charArray); i++ {
		if charArray[i] == "_" { //Syntax should match a whole number
			if expectNumber {
				return 0.0, errors.New("Invalid Syntax")
			}
			expectNumber = true
			whole = tempNum
			tempNum = ""
		} else if charArray[i] == "/" { //Syntax should match a fraction
			if expectNumber {
				return 0.0, errors.New("Invalid Syntax")
			}
			expectNumber = true
			numerator = tempNum
			tempNum = ""
		} else { //Syntax should match a whole number
			tempNum = tempNum + charArray[i]
			expectNumber = false
		}
	}
	if numerator != "" {
		denominator = tempNum
	} else {
		whole = tempNum
	}

	fraction := 0.0
	//First process as if it was an improper fraction or a fraction
	if denominator == "" && numerator == "" { //It is a whole number
		return strconv.ParseFloat(whole, 2)
	} else if (denominator != "") && numerator != "" { //This is a mixed number, we need to convert mixed to decimal.
		denominatorNumber, err := strconv.Atoi(denominator)
		if err != nil {
			return 0.0, errors.New("Invalid Syntax")
		}
		numeratorNumber, err := strconv.Atoi(numerator)
		if err != nil {
			return 0.0, errors.New("Invalid Syntax")
		}
		fraction = float64(numeratorNumber) / float64(denominatorNumber)
	}

	wholeNumber := 0
	if whole != "" {
		var err error
		wholeNumber, err = strconv.Atoi(whole)
		if err != nil || wholeNumber <= 0 {
			return 0.0, errors.New("Invalid Syntax")
		}
	}

	return float64(wholeNumber) + fraction, nil

}
