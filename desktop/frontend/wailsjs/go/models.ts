export namespace viewModels {
	
	export class UpdateCustomerVM {
	    id: number;
	    name: string;
	    surname: string;
	    phone: string;
	    address: string;
	
	    static createFrom(source: any = {}) {
	        return new UpdateCustomerVM(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.name = source["name"];
	        this.surname = source["surname"];
	        this.phone = source["phone"];
	        this.address = source["address"];
	    }
	}
	export class UserResponse {
	    user_id: number;
	    user_name: string;
	    token: string;
	
	    static createFrom(source: any = {}) {
	        return new UserResponse(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.user_id = source["user_id"];
	        this.user_name = source["user_name"];
	        this.token = source["token"];
	    }
	}

}

