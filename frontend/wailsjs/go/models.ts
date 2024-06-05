export namespace builder {
	
	export class GenerateRequest {
	    lang: string;
	    table: string;
	    template: string;
	    env: {[key: string]: any};
	
	    static createFrom(source: any = {}) {
	        return new GenerateRequest(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.lang = source["lang"];
	        this.table = source["table"];
	        this.template = source["template"];
	        this.env = source["env"];
	    }
	}

}

export namespace config {
	
	export class Config {
	    driver: string;
	    host: string;
	    username: string;
	    password: string;
	    database: string;
	    prefixes: string;
	
	    static createFrom(source: any = {}) {
	        return new Config(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.driver = source["driver"];
	        this.host = source["host"];
	        this.username = source["username"];
	        this.password = source["password"];
	        this.database = source["database"];
	        this.prefixes = source["prefixes"];
	    }
	}

}

