export interface OverallStats {
    Accuracy: number;
    RankedAccuracy: number;
    Performance: number;
}

export interface TotalStats {
    Hits: number;
    Score: number;
}

export interface PlayStats {
    Playcount: number;
    Passed: number;
    Failed: number;
}

export interface LongestStats {
    Hits: LongestStatsCombo;
    Combo: LongestStatsCombo;
}

export interface LongestStatsCombo {
    ScoreID: number;
    MapID: number;
    Map: {
        Title: string;
        DifficultyName: string;
    };

    Total: number;
}

export interface MostPlayed {
    MapID: number;
    Title: string;
    DifficultyName: string;
    Creator: string;
    PlayCount: number;
}
