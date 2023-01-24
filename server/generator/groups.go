package generator

// Assigns days to groups.
func splitDaysIntoGroups(days []*WorkingDay, groups []*Group) []*WorkingDay {

	groupIndex := 0
	groupHours := groups[groupIndex].Hours

	for dayIndex := range days {

		if days[dayIndex].Hours == 0 {
			continue
		}

		days[dayIndex].GroupId = groups[groupIndex].GroupId

		if (groupHours - days[dayIndex].Hours) < 0 {

			remainingHours := days[dayIndex].Hours - groupHours
			days[dayIndex].Hours = groupHours

			days[dayIndex+1].Hours += remainingHours

			groupIndex += 1
			groupHours = groups[groupIndex].Hours

		} else {
			groupHours -= days[dayIndex].Hours
			if groupHours == 0 && len(groups) > groupIndex+1 {
				groupIndex += 1
				groupHours = groups[groupIndex].Hours
			}
		}
	}

	return days
}
