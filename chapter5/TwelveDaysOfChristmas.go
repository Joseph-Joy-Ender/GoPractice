package main

import "fmt"

func christmasSong() {
	for days := 1; days <= 12; days++ {
		switch days {
		case 1:
			fmt.Println("On the First day of christmas\n my true love sent to me")
		case 2:
			fmt.Println("On the Second day of christmas\n my true love sent to me")
		case 3:
			fmt.Println("On the Third day of christmas\n my true love sent to me")
		case 4:
			fmt.Println("On the Fourth day of christmas\n my true love sent to me")
		case 5:
			fmt.Println("On the Fifth day of christmas\n my true love sent to me")
		case 6:
			fmt.Println("On the Sixth day of christmas\n my true love sent to me")
		case 7:
			fmt.Println("On the Seventh day of christmas\n my true love sent to me")
		case 8:
			fmt.Println("On the Eighth day of christmas\n my true love sent to me")
		case 9:
			fmt.Println("On the Ninth day of christmas\n my true love sent to me")
		case 10:
			fmt.Println("On the Tenth day of christmas\n my true love sent to me")
		case 11:
			fmt.Println("On the Eleventh day of christmas\n my true love sent to me")
		case 12:
			fmt.Println("On the Twelfth day of christmas\n my true love sent to me")
		}

		switch days {
		case 12:
			fmt.Println("Twelve drummers drumming")
		case 11:
			fmt.Println("Eleven pipers piping")
		case 10:
			fmt.Println("Ten lords a- leaping")
		case 9:
			fmt.Println("Nine ladies dancing")
		case 8:
			fmt.Println("Eight maids a- milking")
		case 7:
			fmt.Println("Seven swarms a- swimming")
		case 6:
			fmt.Println("Six geese a- laying")
		case 5:
			fmt.Println("Five golden rings")
		case 4:
			fmt.Println("Four calling birds")
		case 3:
			fmt.Println("Three french hens")
		case 2:
			fmt.Println("Two turtle doves")
		case 1:
			fmt.Println("A partridge in a pear tree")
			fmt.Println()
		}

	}
}
