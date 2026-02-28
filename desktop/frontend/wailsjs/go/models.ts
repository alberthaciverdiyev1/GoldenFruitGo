export namespace viewModels {
	
	export class LoginResult {
	    success: boolean;
	    message: string;
	    user?: string;
	    token?: string;
	
	    static createFrom(source: any = {}) {
	        return new LoginResult(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.success = source["success"];
	        this.message = source["message"];
	        this.user = source["user"];
	        this.token = source["token"];
	    }
	}

}

