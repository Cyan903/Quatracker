import { useWails } from "@/use/useWails";
import type {
    MostPlayed,
    LongestStats,
    OverallStats,
    PlayStats,
    TotalStats,
    JudgesCount,
    GradesCount,
    PlayHistory,
} from "@/types/stats";

import {
    GetMostPlayed,
    GetLongestStats,
    GetOverallStats,
    GetPlayStats,
    GetTotalStats,
    GetJudgesCountStats,
    GetGradeCount,
    GetUserPlayGraph,
} from "@wails/go/app/App";

export async function getOverall(uid: number, mode: number) {
    return await useWails<OverallStats>(GetOverallStats, uid, mode);
}

export async function getTotal(uid: number, mode: number) {
    return await useWails<TotalStats>(GetTotalStats, uid, mode);
}

export async function getPlays(uid: number, mode: number) {
    return await useWails<PlayStats>(GetPlayStats, uid, mode);
}

export async function getLongest(uid: number, mode: number) {
    return await useWails<LongestStats>(GetLongestStats, uid, mode);
}

export async function getMost(uid: number, mode: number, page: number) {
    return await useWails<MostPlayed[]>(GetMostPlayed, uid, mode, page);
}

export async function getJudges(uid: number, mode: number) {
    return await useWails<JudgesCount>(GetJudgesCountStats, uid, mode);
}

export async function getGrades(
    uid: number,
    mode: number,
    personalBest: boolean,
) {
    return await useWails<GradesCount>(GetGradeCount, uid, mode, personalBest);
}

export async function getHistory(uid: number, mode: number) {
    return await useWails<PlayHistory>(GetUserPlayGraph, uid, mode);
}
