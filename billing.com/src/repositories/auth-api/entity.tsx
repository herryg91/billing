export type UserToken = {
	auth_token: string
	refresh_token: string
}

export type AuthenticatedUser = {
	id: number
	email: string
	name: string
}
