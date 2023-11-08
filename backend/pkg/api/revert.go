package api

func RevertRankedStatus(status string) int {
	switch status {
	case "Not Submitted":
		return 0
	case "Unranked":
		return 1
	case "Ranked":
		return 2
	case "Dan Course":
		return 3
	default:
		return -1
	}
}
