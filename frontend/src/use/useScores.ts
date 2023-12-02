import type { ScoreDetails, Scores } from "../types/scores";
import {
    GetBestScores,
    GetRecentScores,
    GetScoreDetails,
} from "../../wailsjs/go/app/App";
import { useWails } from "./useWails";

export async function getBest(
    uid: number,
    mode: number,
    page: number,
    judgement: string,
    status: string,
    ln: number,
) {
    return await useWails<Scores[]>(GetBestScores, ...arguments);
}

export async function getRecent(
    uid: number,
    mode: number,
    page: number,
    failed: boolean,
) {
    return await useWails<Scores[]>(GetRecentScores, ...arguments);
}

export async function getDetails(id: number) {
    return await useWails<ScoreDetails>(GetScoreDetails, id);
}
