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
	export class UserProfile {
	    user_id: string;
	    username: string;
	    email: string;
	    role: string;
	    profile_image: string;
	    created_at: string;
	
	    static createFrom(source: any = {}) {
	        return new UserProfile(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.user_id = source["user_id"];
	        this.username = source["username"];
	        this.email = source["email"];
	        this.role = source["role"];
	        this.profile_image = source["profile_image"];
	        this.created_at = source["created_at"];
	    }
	}

}

