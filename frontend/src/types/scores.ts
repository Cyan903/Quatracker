import type { RankedStatus } from "@/types/maps";

// Scores
export interface Scores {
    ScoreID: number;
    MapID: number;

    Score: {
        DateTime: string;
        Grade: string;
        Accuracy: number;
        RankedAccuracy: number;
        Mods: string;
        PerformanceRating: number;
        JudgementWindowPreset: string;
        PersonalBest: boolean;
    };

    Map: {
        Artist: string;
        Title: string;
        DifficultyName: string;
        Creator: string;
        RankedStatus: RankedStatus;
        LNPercent: number;
        Rating: number;
    };
}

// Details
interface Judge {
    Accuracy?: number;
    RankedAccuracy?: number;
    CountMarv: number;
    CountPerf: number;
    CountGreat: number;
    CountGood: number;
    CountOkay: number;
    CountMiss: number;
}

export interface ScoreDetails {
    ScoreID: number;
    MapID: number;

    Score: {
        LocalProfileId: number;
        Name: string;
        DateTime: string;
        TotalScore: number;
        Grade: string;
        MaxCombo: number;
        Mods: string;
        Mode: number;
        ScrollSpeed: number;
        PauseCount: number;
        PerformanceRating: number;
        JudgementWindowPreset: string;
        JudgedHits: Judge;
        JudgementConfig: Judge;
        PersonalBest: boolean;

        Versions: {
            RatingVersion: string;
            DifficultyVersion: string;
        };
    };

    Map: {
        Artist: string;
        Title: string;
        DifficultyName: string;
        Creator: string;
        RankedStatus: RankedStatus;

        DifficultyInfo: {
            SongLength: number;
            BPM: number;
            LNPercent: number;
            Rating: number;
        };

        ModeInfo: {
            DifficultyProcessorVersion: string;
            HasScratchKey: boolean;
        };
    };
}
