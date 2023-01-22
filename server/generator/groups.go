package generator

func splitDaysIntoGroups(days []*WorkingDay, a_hours int) []*WorkingDay {

	groupName := "A"

	for dayIndex := range days {

		if days[dayIndex].Hours == 0 {
			continue
		}

		days[dayIndex].GroupId = groupName

		if a_hours != 0 {

			if (a_hours - days[dayIndex].Hours) < 0 {

				remainingHours := days[dayIndex].Hours - a_hours
				days[dayIndex].Hours = a_hours

				days[dayIndex+1].Hours += remainingHours
				groupName = "B"
				a_hours = 0
			} else {
				a_hours -= days[dayIndex].Hours
			}
		}
	}

	return days
}
