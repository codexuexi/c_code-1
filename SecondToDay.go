package c_code

import "fmt"

func SecondsToDay(mysec int64) string {
	mins := int64(0)
	hours := int64(0)
	days := int64(0)

	if mysec >= 60 {
		mins = mysec / 60
		mysec = mysec % 60
	}
	if mins >= 60 {
		hours = mins / 60
		mins = mins % 60
	}
	if hours >= 24 {
		days = hours / 24
		hours = hours % 60
	}

	return fmt.Sprintf("%d 天 %d 时 %d 分", days, hours, mins)
}
