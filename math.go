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
	// find numbers ending in "d" and replace with expression to convert degrees to radians
	input = strings.Replace(input, "d", "*"+strconv.FormatFloat(math.Pi/180, 'f', -1, 64), -1)
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

func r2p(line string) string {
	// format: r2p (x, y)
	// remove first word (command)
	line = line[4:]
	// find x
	x := line[1:strings.Index(line, ",")]
	// evaluate x
	xValue, err := basicMath(x)
	if err != nil {
		return "Error: " + err.Error()
	}
	// find y
	y := line[strings.Index(line, ",")+1:]
	// if it exists, remove the last character (")")
	if y[len(y)-1] == ')' {
		y = y[:len(y)-1]
	}
	// evaluate y
	yValue, err := basicMath(y)
	if err != nil {
		return "Error: " + err.Error()
	}
	// calculate r and theta
	r := math.Sqrt(xValue*xValue + yValue*yValue)
	theta := math.Atan2(yValue, xValue)
	// convert theta to degrees
	theta = theta * 180 / math.Pi
	// return result
	return "(" + strconv.FormatFloat(r, 'f', -1, 64) + ", " + strconv.FormatFloat(theta, 'f', -1, 64) + "°)"
}

func polarVariations(line string) string {
	// find all possible ways to describe a polar coordinate
	// format: pVariations (r, theta)
	// remove first word (command)
	line = line[12:]
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
	// there are four possible ways to describe a polar coordinate
	// if r is positive, theta is either theta or 360 - theta
	// if r is negative, theta is either theta or -(360 - theta)
	// we can also make r opposite sign and then find the equivalent theta

	var potentialVariations []string

	// first, find theta in degrees
	thetaDegrees := thetaValue * 180 / math.Pi

	// do the first two cases
	if rValue > 0 {
		potentialVariations = append(potentialVariations, "("+strconv.FormatFloat(rValue, 'f', -1, 64)+", "+strconv.FormatFloat(thetaDegrees, 'f', -1, 64)+"°)")
		potentialVariations = append(potentialVariations, "("+strconv.FormatFloat(rValue, 'f', -1, 64)+", -"+strconv.FormatFloat(360-thetaDegrees, 'f', -1, 64)+"°)")
	} else {
		potentialVariations = append(potentialVariations, "("+strconv.FormatFloat(-rValue, 'f', -1, 64)+", -"+strconv.FormatFloat(thetaDegrees, 'f', -1, 64)+"°)")
		potentialVariations = append(potentialVariations, "("+strconv.FormatFloat(-rValue, 'f', -1, 64)+", "+strconv.FormatFloat(360-thetaDegrees, 'f', -1, 64)+"°)")
	}

	// find the equivalent theta for the second case
	thetaOpposite := thetaValue + math.Pi
	if thetaOpposite > math.Pi {
		thetaOpposite = thetaOpposite - math.Pi*2
	}

	// convert thetaOpposite to degrees
	thetaOppositeDegrees := thetaOpposite * 180 / math.Pi

	// do the other two cases
	if rValue > 0 {
		potentialVariations = append(potentialVariations, "(-"+strconv.FormatFloat(rValue, 'f', -1, 64)+", "+strconv.FormatFloat(thetaOppositeDegrees, 'f', -1, 64)+"°)")
		potentialVariations = append(potentialVariations, "(-"+strconv.FormatFloat(rValue, 'f', -1, 64)+", "+strconv.FormatFloat(360+thetaOppositeDegrees, 'f', -1, 64)+"°)")
	} else {
		potentialVariations = append(potentialVariations, "("+strconv.FormatFloat(rValue, 'f', -1, 64)+", -"+strconv.FormatFloat(thetaOppositeDegrees, 'f', -1, 64)+"°)")
		potentialVariations = append(potentialVariations, "("+strconv.FormatFloat(rValue, 'f', -1, 64)+", "+strconv.FormatFloat(360+thetaOppositeDegrees, 'f', -1, 64)+"°)")
	}

	// return result
	return strings.Join(potentialVariations, ", ")
}
