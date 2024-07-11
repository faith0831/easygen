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
	export class GetGenerateColumnRequest {
	    table: string;
	
	    static createFrom(source: any = {}) {
	        return new GetGenerateColumnRequest(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.table = source["table"];
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
	    filteredTablePrefixes: string;
	    filteredCreatedColumns: string;
	    filteredUpdatedColumns: string;
	
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
	        this.filteredTablePrefixes = source["filteredTablePrefixes"];
	        this.filteredCreatedColumns = source["filteredCreatedColumns"];
	        this.filteredUpdatedColumns = source["filteredUpdatedColumns"];
	    }
	}

}

