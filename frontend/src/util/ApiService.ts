class ApiService {

    public static instance: ApiService
    public static init() {

    }
    public static getInstance() {
        if (ApiService.instance) {
            return ApiService.instance
        } else {
            throw new Error("ApiService must be initialized first")
        }
    }

    private apiBase: string
    private bearerToken: string = ""
    private loginRoute: string

    constructor(apiBase: string, loginRoute: string) {
        this.apiBase = apiBase
        this.loginRoute = loginRoute
    }

    public async relogin(username: string, password: string) {
        const requestBody: unknown = {
            username,
            password
        }
        const response = await fetch(this.apiBase + this.loginRoute, {
            method: "POST",
            body: JSON.stringify(requestBody),
            headers: {
                'Content-Type': 'application/json'
            }
        })
        const data = await response.json()
        this.bearerToken = data.token
    }

    public async query<T>(route: string, method: string = "GET", body?: unknown) {
        const requestBody: unknown = body
        const response = await fetch(this.apiBase + route, {
            method: method,
            body: JSON.stringify(requestBody),
            headers: {
                'Content-Type': 'application/json',
                'Authorization': 'Bearer ' + this.bearerToken
            }
        })
        return await response.json() as T
    }
}

export {
    ApiService
}