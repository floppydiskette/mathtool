package main

import (
	"github.com/Knetic/govaluate"
	"math"
	"strconv"
	"strings"
)

func basicMath(input string) (float64, error) {
	// replace all occurances of "pi" with a string representation of pi
	input = strings.Replace(input, "pi", strconv.FormatFloat(math.Pi, 'f', -1, 64), -1)
	expression, err := govaluate.NewEvaluableExpression(input)
	if err != nil {
		return 0, err
	}
	parameters := make(map[string]interface{})
	result, err := expression.Evaluate(parameters)
	if err != nil {
		return 0, err
	}
	return result.(float64), nil
}

func p2r(line string) string {
	// format: p2r (r, theta)
	// remove first word (command)
	line = line[4:]
	// find r
	r := line[1:strings.Index(line, ",")]
	// evaluate r
	rValue, err := basicMath(r)
	if err != nil {
		return "Error: " + err.Error()
	}
	// find theta
	theta := line[strings.Index(line, ",")+1:]
	// if it exists, remove the last character (")")
	if theta[len(theta)-1] == ')' {
		theta = theta[:len(theta)-1]
	}
	// evaluate theta
	thetaValue, err := basicMath(theta)
	if err != nil {
		return "Error: " + err.Error()
	}
	// calculate x and y
	x := rValue * math.Cos(thetaValue)
	y := rValue * math.Sin(thetaValue)
	// return result
	return "(" + strconv.FormatFloat(x, 'f', -1, 64) + ", " + strconv.FormatFloat(y, 'f', -1, 64) + ")"
}
