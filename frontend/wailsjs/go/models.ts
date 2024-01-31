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

