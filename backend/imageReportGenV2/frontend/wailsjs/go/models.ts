export namespace app {
	
	export class GenerateRequest {
	    templatePath: string;
	    outputPath: string;
	    clientName: string;
	    city: string;
	    state: string;
	    includeClient: boolean;
	    includeLocation: boolean;
	    includeDate: boolean;
	    completionDate: string;
	    images: string[];
	
	    static createFrom(source: any = {}) {
	        return new GenerateRequest(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.templatePath = source["templatePath"];
	        this.outputPath = source["outputPath"];
	        this.clientName = source["clientName"];
	        this.city = source["city"];
	        this.state = source["state"];
	        this.includeClient = source["includeClient"];
	        this.includeLocation = source["includeLocation"];
	        this.includeDate = source["includeDate"];
	        this.completionDate = source["completionDate"];
	        this.images = source["images"];
	    }
	}

}

