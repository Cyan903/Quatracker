import type { ScoreDetails, Scores } from "@/types/scores";
import {
    GetBestScores,
    GetRecentScores,
    GetScoreDetails,
} from "@wails/go/app/App";
import { useWails } from "@/use/useWails";

export async function getBest(
    uid: number,
    mode: number,
    page: number,
    judgement: string,
    status: string,
    ln: number,
) {
    return await useWails<Scores[]>(
        GetBestScores,
        uid,
        mode,
        page,
        judgement,
        status,
        ln,
    );
}

export async function getRecent(
    uid: number,
    mode: number,
    page: number,
    failed: boolean,
) {
    return await useWails<Scores[]>(GetRecentScores, uid, mode, page, failed);
}

export async function getDetails(id: number) {
    return await useWails<ScoreDetails>(GetScoreDetails, id);
}
