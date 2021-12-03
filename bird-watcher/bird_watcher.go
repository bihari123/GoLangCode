package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
    sum:=0
    for _,num:=range birdsPerDay{
        sum+=num
    }
    return sum
	panic("Please implement the TotalBirdCount() function")
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
    var weekData []int
    weekData=append(birdsPerDay[7*(week-1):(7*(week-1))+7])
    return TotalBirdCount(weekData)
	panic("Please implement the BirdsInWeek() function")
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
    for index,num:= range birdsPerDay{
        if index%2==0{
           birdsPerDay[index]=num+1
        }
    } 
    return birdsPerDay
	panic("Please implement the FixBirdCountLog() function")
}
