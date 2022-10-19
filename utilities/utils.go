package utils

import (
	"fmt"
	"strings"
	"gopkg.in/toast.v1"
)

// Returns a value to a type of string
//
// Usage: utilities.JoinArray[Type](Value of type Type, One Line Stringify)
//
// Example:
//
//	numb := 25
//	str := utilities.Stringify[int](numb, false)
//	fmt.Print(str) // "25"
func Stringify[T any](value T, one_line bool) string {
	var str string

	if !one_line {
		str = fmt.Sprint(value)
	} else {
		str = fmt.Sprintln(value)
	}

	return str
}

// Returns a string which is joined by item & seperator to seperate each item from each other by a string.
//
// Usage: utilities.JoinArray[Type](Array of type Type, Seperator)
//
// Example:
//
//	arr := []int{ 1, 3, 7, 4, 6, 8, 1 }
//	str := utilities.JoinArray[int](arr, "\n")
//	fmt.Print(str) /* "1
//	3
//	7
//	4
//	6
//	8
//	1" */
func JoinArray[Type any](array []Type, seperator string) string {
	var result string
	for i := 0; i < len(array); i++ {
		length := len(array) - 1
		if i == length {
			result = result + fmt.Sprint(array[i])
		}

		result = result + fmt.Sprint(array[i]) + seperator
	}

	return result
}

// Appends an item to the specified array using the specified value.
//
// Usage: utilities.AppendArray[Type](Array of type Type, Value of type Type)
//
// Example:
//
//	arr1 := []int{ 1, 3, 7, 4, 6, 8, 1 }
//	arr2 := utilities.AppendArray[int](arr1, 3)
//	fmt.Print(arr2) // 1, 3, 7, 4, 6, 8, 1, 3
func AppendArray[T any](array []T, value T) []T {
	var arr []T = append(array, value)

	return arr
}

// Formats a partition if it doesn't include ":", else, return the given partition.
//
// Usage: utilities.FormatPartition(Partition)
//
// Example:
// 	partition := "c"
// 	formatted := utilities.FormatPartition(partition)
//	fmt.Print(formatted) // c:
func FormatPartition(partition string) string {
	var part string 
	
	if strings.Contains(partition, ":") {
		part = partition
	} else {
		part = partition + ":"
	}

	return part
}

func SendFinishedNotification(command string) {
	notification := toast.Notification{
		AppID: "Commander",
		Title: "Commander | Commando Uitgevoerd",
		Message: command + " is klaar, wil je nog een commando uitvoeren",
		Icon: "icon.png",
		Actions: []toast.Action{
			{
				Type: "protocol", 
				Label: "Ja, graag", 
				Arguments: ""
			},
			{Type: "protocol", Label: "Nee, ik ben klaar", Arguments: ""},
		},
	}

	RunButtonAction(notification.Actions[0])
	RunButtonAction(notification.Actions[1])
}

func RunButtonAction(button toast.Action) {
	
}