import type { RankedStatus } from "../types/maps";
import { GetBestScores, GetRecentScores } from "../../wailsjs/go/app/App";
import { useWails } from "./useWails";

interface MapInfo {
    Artist: string;
    Title: string;
    DifficultyName: string;
    Creator: string;
    RankedStatus: RankedStatus;
    LNPercent: number;
}

interface ScoreInfo {
    Name: string;
    DateTime: string;
    TotalScore: number;
    Grade: string;
    Accuracy: number;
    RankedAccuracy: number;
    Mods: string;
    PerformanceRating: number;
    JudgementWindowPreset: string;
}

export interface Scores {
    ScoreID: number;
    MapID: number;
    Score: ScoreInfo;
    Map: MapInfo;
}

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
