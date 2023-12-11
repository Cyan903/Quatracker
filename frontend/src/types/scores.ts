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

export interface CountedScores {
    Total: number;
    Scores: Scores[];
}

// Details
export interface Judges {
    CountMarv: number;
    CountPerf: number;
    CountGreat: number;
    CountGood: number;
    CountOkay: number;
    CountMiss: number;
}

export interface JudgesAcc extends Judges {
    Accuracy: number;
    Grade: string;
    RankedAccuracy: number;
    RankedGrade: string;
}

export interface ScoreDetails {
    ScoreID: number;
    MapID: number;

    Score: {
        LocalProfileId: number;
        Name: string;
        DateTime: string;
        TotalScore: number;
        MaxCombo: number;
        Mods: string;
        ScrollSpeed: number;
        PauseCount: number;
        PerformanceRating: number;
        JudgementWindowPreset: string;
        PersonalBest: boolean;
        JudgedHits: JudgesAcc;
        JudgementConfig: Judges;
        RatingVersion: string;
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
            Mode: number;
            DifficultyProcessorVersion: string;
            HasScratchKey: boolean;
        };
    };
}
