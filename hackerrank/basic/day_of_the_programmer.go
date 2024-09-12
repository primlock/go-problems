/*
https://www.hackerrank.com/challenges/day-of-the-programmer/problem

- Given a year, find the date of the 256th day of that year (Day of the Programmer) according to the official
Russian calendar during that year.
- The range of dates to find the Day of the Programmer is between 1700 and 2700.
- From 1700 to 1917 Russia's official calendar was the Julian calendar.
- Since 1919 Russia's official calendar is the Gregorian calendar.
- In the year 1918 the next day after January 31st was February 14th, making February 14th the 32nd day of the
year.

- February has a variable amount of days for both calendars, 29 days for a leap year and 28 days otherwise.
	- Julian leap year is divisible by 4.
	- Gregorian leap year is divisible by 400 OR divisible by 4 but not divisible by 100.

- Return a string of the 256th day of that year in the format dd.mm.yyyy
---------------------------------------------------------
- Determine if this year uses the Julian calendar, Gregorian calendar or is the transition year 1918.
	- If year is <= 1917 then its Julian
	- If year is >= 1919 then its Gregorian
	- Else its the transition year

- Determine if this was a leap year or a non-leap year.
	- This depends on the calendar system.
	- Julian:
		- If year % 4 == 0 than its a leap year.
	- Gregorian:
		- If year % 4 == 0 & year % 100 != 0 then its a leap year.

- With the information above, determine what day the 256th day fell on.
	- If this year was a leap year, February had 29 days. It has 28 days otherwise.
		- All other months have the same amount of days.
			- January 31
			- March 31
			- April 30
			- May 31
			- June 30
			- July 31
			- August 31
		- We only need the first 8 months (excluding Feb or we will go over 256 days).

	- Adding up all the days between January and August (8 months) will land you on September 1st of that year.
		- Taking the difference between the sum and 256 will tell you what day that 256th day was.
*/

package basic

import (
	"fmt"
	"math"
)

const (
	January = 31
	March   = 31
	April   = 30
	May     = 31
	June    = 30
	July    = 31
	August  = 31
)

func DayOfProgrammer(year int32) string {
	target_date := ""
	february := 28

	if year <= 1917 {
		// Julian calendar
		if year%4 == 0 {
			february = 29
		}
	} else if year >= 1919 {
		// Gregorian calendar
		leap := false
		if year%400 == 0 {
			leap = true
		} else if year%4 == 0 && year%100 != 0 {
			leap = true
		}

		if leap {
			february = 29
		}
	} else {
		// Transition
		february = 15
	}

	sum := January + february + March + April + May + June + July + August
	day := int(math.Abs(float64(sum) - 256))

	target_date = fmt.Sprintf("%02d.09.%d\n", day, year)

	return target_date
}
