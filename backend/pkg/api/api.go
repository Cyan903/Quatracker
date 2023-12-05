package api

// https://github.com/Quaver/Quaver.API/blob/8950e207ead59e6cc2fa22fa1827db5644337302/Quaver.API/Helpers/GradeHelper.cs#L20C65-L20C65
func ConvertGrade(accuracy float64, grade int) string {
	if grade == 5 {
		return "F"
	}

	if accuracy == 100 {
		return "X"
	} else if accuracy >= 99.0 {
		return "SS"
	} else if accuracy >= 95.0 {
		return "S"
	} else if accuracy >= 90.0 {
		return "A"
	} else if accuracy >= 80.0 {
		return "B"
	} else if accuracy >= 70.0 {
		return "C"
	}

	return "D"
}

// https://github.com/Quaver/Quaver.API/blob/8950e207ead59e6cc2fa22fa1827db5644337302/Quaver.API/Enums/RankedStatus.cs#L8
func ConvertRankedStatus(status int) string {
	rank := [4]string{"Not Submitted", "Unranked", "Ranked", "Dan Course"}

	if status > 3 || status < 0 {
		return "Unknown"
	}

	return rank[status]
}
