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
	export class FileHash {
	    fileName: string;
	    directory: string;
	    hash: string;
	
	    static createFrom(source: any = {}) {
	        return new FileHash(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.fileName = source["fileName"];
	        this.directory = source["directory"];
	        this.hash = source["hash"];
	    }
	}
	export class UpdateCheckResult {
	    needsUpdate: boolean;
	    currentVersion: string;
	    serverVersion: string;
	    filesToUpdate: FileHash[];
	
	    static createFrom(source: any = {}) {
	        return new UpdateCheckResult(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.needsUpdate = source["needsUpdate"];
	        this.currentVersion = source["currentVersion"];
	        this.serverVersion = source["serverVersion"];
	        this.filesToUpdate = this.convertValues(source["filesToUpdate"], FileHash);
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

