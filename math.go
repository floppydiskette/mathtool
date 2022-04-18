package main

import (
	"github.com/Knetic/govaluate"
	"math"
	"strconv"
	"strings"
)

// if we should show solution steps
var steps bool = false

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

// function to simplify a fraction
func simpFrac(top, bot int) (int, int) {
	if steps {
		print("Simplifying fraction: " + strconv.Itoa(top) + "/" + strconv.Itoa(bot) + "\n")
	}
	for i := 2; i <= top; i++ {
		if steps {
			print("Checking if " + strconv.Itoa(i) + " is a factor of " + strconv.Itoa(bot) + "\n")
		}
		if top%i == 0 && bot%i == 0 {
			if steps {
				print("Found factor: " + strconv.Itoa(i) + "\n")
			}
			top /= i
			bot /= i
			i = 1
			if steps {
				print("Simplified to: " + strconv.Itoa(top) + "/" + strconv.Itoa(bot) + "\n")
			}
		}
	}
	return top, bot
}

func simplifyFraction(line string) string {
	// input will be simp x/y
	// remove command
	line = line[5:]
	// evaluate each side of the fraction
	top, err := basicMath(line[:strings.Index(line, "/")])
	if err != nil {
		return "Error: " + err.Error()
	}
	bot, err := basicMath(line[strings.Index(line, "/")+1:])
	if err != nil {
		return "Error: " + err.Error()
	}
	// simplify the fraction
	topInt, botInt := simpFrac(int(top), int(bot))
	// return the simplified fraction
	return strconv.Itoa(topInt) + "/" + strconv.Itoa(botInt)
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
	if steps {
		print("pythagorean theorem: " + strconv.FormatFloat(xValue, 'f', -1, 64) + "^2 + " + strconv.FormatFloat(yValue, 'f', -1, 64) + "^2\n")
	}
	r := math.Sqrt(xValue*xValue + yValue*yValue)
	if steps {
		print("r: " + strconv.FormatFloat(r, 'f', -1, 64) + "\n")
	}
	if steps {
		print("getting theta from arctan: " + strconv.FormatFloat(yValue, 'f', -1, 64) + " / " + strconv.FormatFloat(xValue, 'f', -1, 64) + "\n")
	}
	theta := math.Atan2(yValue, xValue)
	// convert theta to degrees
	if steps {
		print("theta: " + strconv.FormatFloat(theta, 'f', -1, 64) + "\n")
		print("converting theta to degrees: " + strconv.FormatFloat(theta*180/math.Pi, 'f', -1, 64) + "\n")
	}
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

func degreesToPolarPi(line string) string {
	// convert theta degrees to a theta like (5pi/4)
	// format: d2PP (theta)
	// remove first word (command)
	line = line[5:]
	// find theta
	theta := line
	// evaluate theta
	thetaValue, err := basicMath(theta)
	if err != nil {
		return "Error: " + err.Error()
	}
	// radian value will be thetaValue * pi / 180
	radianNumerator := thetaValue
	radianDenominator := 180
	// simplify the fraction
	top, bot := simpFrac(int(radianNumerator), int(radianDenominator))
	// return formatted result
	return "(" + strconv.Itoa(top) + "pi/" + strconv.Itoa(bot) + ")"
}

func radsToDegrees(line string) string {
	// convert radians to degrees
	// format: rad2d (radians)
	// remove first word (command)
	line = line[6:]
	// find radians
	radians := line
	// evaluate radians
	radiansValue, err := basicMath(radians)
	if err != nil {
		return "Error: " + err.Error()
	}
	// degrees value will be radiansValue * 180 / pi
	degrees := radiansValue * 180 / math.Pi
	// return formatted result
	return strconv.FormatFloat(degrees, 'f', -1, 64)
}
