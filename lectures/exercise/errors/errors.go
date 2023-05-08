//--Summary:
//  Create a function that can parse time strings into component values.
//
//--Requirements:
//* The function must parse a string into a struct containing:
//  - Hour, minute, and second integer components
//* If parsing fails, then a descriptive error must be returned
//* Write some unit tests to check your work
//  - Run tests with `go test ./exercise/errors`
//
//--Notes:
//* Example time string: 14:07:33
//* Use the `strings` package from stdlib to get time components
//* Use the `strconv` package from stdlib to convert strings to ints
//* Use the `errors` package to generate errors

package main

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"
)

type RealTime struct {
	Hour   int
	Minute int
	Second int
}

//* The function must parse a string into a struct containing:
//  - Hour, minute, and second integer components

func TimeParse(s string) (RealTime, error) {
	newStruct := RealTime{}
	timeString := strings.Split(s, ":")

	if len(timeString) != 3 {
		return RealTime{}, errors.New("hours, minutes and seconds must be provided in full")
	}

	//HOUR
	hr, err := strconv.Atoi(timeString[0])
	//* If parsing fails, then a descriptive error must be returned
	if err != nil {
		fmt.Println(errors.New("error parsing the Hour"))
		return RealTime{}, err
	} else if hr < 0 || hr >= 24 {
		return RealTime{}, errors.New("hour is not in limits of 0 to 23")
	} else {
		newStruct.Hour = hr
	}
	//MINUTE
	min, err := strconv.Atoi(timeString[1])
	//* If parsing fails, then a descriptive error must be returned
	if err != nil {
		fmt.Println(errors.New("error parsing the Minute"))
		return RealTime{}, err
	} else if min < 0 || min >= 59 {
		return RealTime{}, errors.New("minute is not in limits of 0 to 59")
	} else {
		newStruct.Minute = min
	}
	//SECOND
	sec, err := strconv.Atoi(timeString[2])
	//* If parsing fails, then a descriptive error must be returned
	if err != nil {
		fmt.Println(errors.New("error parsing the Second"))
		return RealTime{}, err
	} else if sec < 0 || sec >= 59 {
		return RealTime{}, errors.New("second is not in limits of 0 to 59")
	} else {
		newStruct.Second = sec
	}
	//RETURN THE STRUCT
	return newStruct, nil
}

func TimeGiver() (RealTime, error) {
	t := time.TimeOnly
	tString := fmt.Sprint(t)
	if reflect.TypeOf(tString) != reflect.TypeOf("abc") {
		//checking if the time string is available
		return RealTime{}, errors.New("time is not obtained in a string format")
	} else {
		return TimeParse(tString)
	}
}
func main() {

	timeStruct, err := TimeGiver()
	if err != nil {
		fmt.Println(errors.New("time didn't parse correctly"))
	} else {
		fmt.Println(timeStruct)
	}

}
