package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	var i, total = 0, 0
	for i = 0; i < len(birdsPerDay); i++ {
		total += birdsPerDay[i]
	}
	return total
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	var i, total, weekStart, weekEnd = 0, 0, (week - 1) * 7, week * 7
	if weekEnd > len(birdsPerDay) {
		weekEnd = len(birdsPerDay)
	}
	for i = weekStart; i < weekEnd; i++ {
		total += birdsPerDay[i]
	}
	return total
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	var i = 0
	for i = 0; i < len(birdsPerDay); i++ {
		if i%2 == 0 {
			birdsPerDay[i]++
		}
	}
	return birdsPerDay
}
