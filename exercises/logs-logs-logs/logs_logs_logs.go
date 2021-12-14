package logs
import "fmt"
// Application identifies the application emitting the given log.
func Application(log string) string {
	for _,character:= range log{

		myRune:=rune(character)
		if myRune == 0x2757{
			return "recommendation"
		}else if myRune == 0x1F50D{
			return "search"
		}else if myRune == 0x2600{
			return "weather"
		}
	}

	return "default"

	panic("Please implement the Application() function")
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	var msg []rune = []rune(log)

	for index,character:=range msg{
		//myRune:= rune(character)
		fmt.Println("index: ",index)
		if character == oldRune{
			fmt.Println("it got in at index: ",index)
			//msg = []rune(log[:index] + log[index+1:])


			msg[index]=newRune
			//log = log[:index] + + log[index+1:]
		}
	}
	newLog:=string(msg)
	return newLog
	panic("Please implement the Replace() function")
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	myLog:=[]rune(log)
	return len(myLog)<=limit
	panic("Please implement the WithinLimit() function")
}
