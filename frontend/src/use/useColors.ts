import type { RankedStatus } from "@/types/maps";

// https://github.com/Quaver/Quaver.Website/blob/edc5c05fb344f3a34c553d33bd86c32457e4d075/src/utils/ColorHelper.ts#L1
export function useDifficulty(diff: number) {
    if (diff < 1) return "difficulty-beginner";
    else if (diff < 2.5) return "difficulty-easy";
    else if (diff < 10) return "difficulty-normal";
    else if (diff < 20) return "difficulty-hard";
    else if (diff < 30) return "difficulty-insane";
    else if (diff < 40) return "difficulty-expert";
    else if (diff < 50) return "difficulty-extreme";

    return "difficulty-rating-master";
}

export function useRank(rank: RankedStatus) {
    switch (rank) {
        case "Ranked":
            return "map-ranked";
        case "Unranked":
            return "map-unranked";
        case "Dan Course":
            return "map-danCourse";
        default:
            return "map-notSubmitted";
    }
}
