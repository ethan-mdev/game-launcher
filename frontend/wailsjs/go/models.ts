export namespace backend {
	
	export class AuthResponse {
	    access_token: string;
	    refresh_token: string;
	    expires_in: number;
	
	    static createFrom(source: any = {}) {
	        return new AuthResponse(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.access_token = source["access_token"];
	        this.refresh_token = source["refresh_token"];
	        this.expires_in = source["expires_in"];
	    }
	}

}

