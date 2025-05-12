export namespace config {
	
	export class Config {
	    GamePath: string;
	    MainMode: boolean;
	
	    static createFrom(source: any = {}) {
	        return new Config(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.GamePath = source["GamePath"];
	        this.MainMode = source["MainMode"];
	    }
	}

}

export namespace database {
	
	export class ScoreBoard {
	    ScoreID: number;
	    MapID: number;
	    // Go type: struct { DateTime string; Grade string; Accuracy float64; RankedAccuracy float64; Mods string; PerformanceRating float64; JudgementWindowPreset string; PersonalBest bool }
	    Score: any;
	    // Go type: struct { Artist string; Title string; DifficultyName string; Creator string; RankedStatus string; LNPercent float64; Rating float64 }
	    Map: any;
	
	    static createFrom(source: any = {}) {
	        return new ScoreBoard(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.ScoreID = source["ScoreID"];
	        this.MapID = source["MapID"];
	        this.Score = this.convertValues(source["Score"], Object);
	        this.Map = this.convertValues(source["Map"], Object);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class CountedScoreBoard {
	    Total: number;
	    Scores: ScoreBoard[];
	
	    static createFrom(source: any = {}) {
	        return new CountedScoreBoard(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Total = source["Total"];
	        this.Scores = this.convertValues(source["Scores"], ScoreBoard);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class GradesCountStat {
	    AllGrades: Record<string, number>;
	    RankedGrades: Record<string, number>;
	
	    static createFrom(source: any = {}) {
	        return new GradesCountStat(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.AllGrades = source["AllGrades"];
	        this.RankedGrades = source["RankedGrades"];
	    }
	}
	export class JudgeCountStat {
	    CountMarv: number;
	    CountPerf: number;
	    CountGreat: number;
	    CountGood: number;
	    CountOkay: number;
	    CountMiss: number;
	
	    static createFrom(source: any = {}) {
	        return new JudgeCountStat(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.CountMarv = source["CountMarv"];
	        this.CountPerf = source["CountPerf"];
	        this.CountGreat = source["CountGreat"];
	        this.CountGood = source["CountGood"];
	        this.CountOkay = source["CountOkay"];
	        this.CountMiss = source["CountMiss"];
	    }
	}
	export class LocalUsers {
	    Id: string;
	    Username: string;
	
	    static createFrom(source: any = {}) {
	        return new LocalUsers(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Id = source["Id"];
	        this.Username = source["Username"];
	    }
	}
	export class LongestMap {
	    ScoreID: number;
	    MapID: number;
	    // Go type: struct { Title string; DifficultyName string }
	    Map: any;
	    Total: number;
	
	    static createFrom(source: any = {}) {
	        return new LongestMap(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.ScoreID = source["ScoreID"];
	        this.MapID = source["MapID"];
	        this.Map = this.convertValues(source["Map"], Object);
	        this.Total = source["Total"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class LongestStat {
	    Hits: LongestMap;
	    Combo: LongestMap;
	
	    static createFrom(source: any = {}) {
	        return new LongestStat(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Hits = this.convertValues(source["Hits"], LongestMap);
	        this.Combo = this.convertValues(source["Combo"], LongestMap);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class MostPlayed {
	    MapID: number;
	    Title: string;
	    DifficultyName: string;
	    Creator: string;
	    PlayCount: number;
	
	    static createFrom(source: any = {}) {
	        return new MostPlayed(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.MapID = source["MapID"];
	        this.Title = source["Title"];
	        this.DifficultyName = source["DifficultyName"];
	        this.Creator = source["Creator"];
	        this.PlayCount = source["PlayCount"];
	    }
	}
	export class OverallStats {
	    Accuracy: number;
	    RankedAccuracy: number;
	    Performance: number;
	
	    static createFrom(source: any = {}) {
	        return new OverallStats(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Accuracy = source["Accuracy"];
	        this.RankedAccuracy = source["RankedAccuracy"];
	        this.Performance = source["Performance"];
	    }
	}
	export class PlayStats {
	    Playcount: number;
	    Passed: number;
	    Failed: number;
	
	    static createFrom(source: any = {}) {
	        return new PlayStats(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Playcount = source["Playcount"];
	        this.Passed = source["Passed"];
	        this.Failed = source["Failed"];
	    }
	}
	
	export class ScoreDetails {
	    ScoreID: number;
	    MapID: number;
	    // Go type: struct { LocalProfileId int; Name string; DateTime string; TotalScore int; MaxCombo int; Mods string; ScrollSpeed int; PauseCount int; PerformanceRating float64; JudgementWindowPreset string; PersonalBest bool; JudgedHits struct { Accuracy float64; Grade string; RankedAccuracy float64; RankedGrade string; CountMarv int; CountPerf int; CountGreat int; CountGood int; CountOkay int; CountMiss int }; JudgementConfig struct { CountMarv int; CountPerf int; CountGreat int; CountGood int; CountOkay int; CountMiss int }; RatingVersion string }
	    Score: any;
	    // Go type: struct { Artist string; Title string; DifficultyName string; Creator string; RankedStatus string; DifficultyInfo struct { SongLength int; BPM float64; LNPercent float64; Rating float64 }; ModeInfo struct { Mode int; DifficultyProcessorVersion string; HasScratchKey bool } }
	    Map: any;
	
	    static createFrom(source: any = {}) {
	        return new ScoreDetails(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.ScoreID = source["ScoreID"];
	        this.MapID = source["MapID"];
	        this.Score = this.convertValues(source["Score"], Object);
	        this.Map = this.convertValues(source["Map"], Object);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class TotalStats {
	    Hits: number;
	    Score: number;
	
	    static createFrom(source: any = {}) {
	        return new TotalStats(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Hits = source["Hits"];
	        this.Score = source["Score"];
	    }
	}
	export class Users {
	    Local: LocalUsers[];
	    Unknown: LocalUsers[];
	
	    static createFrom(source: any = {}) {
	        return new Users(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Local = this.convertValues(source["Local"], LocalUsers);
	        this.Unknown = this.convertValues(source["Unknown"], LocalUsers);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}

}

export namespace images {
	
	export class FileLoader {
	    Handler: any;
	
	    static createFrom(source: any = {}) {
	        return new FileLoader(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Handler = source["Handler"];
	    }
	}

}

