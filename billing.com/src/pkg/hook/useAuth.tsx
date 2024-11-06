import React, { createContext, useState, useContext, useEffect, ReactNode } from 'react'
import { UserToken } from '@/repositories/usertoken-api/param';
import { useRouter } from 'next/router';
import { deleteCookie, getCookie, setCookie } from 'cookies-next';

const KEY_AUTH_TOKEN = "auth_token"
const KEY_REFRESH_TOKEN = "refresh_token"

type auth_ctx = {
    isAuthenticated: boolean
    token?: UserToken
    login: (token: UserToken) => void
    logout: () => void
};

const AuthContext = createContext<auth_ctx>({ 
    isAuthenticated: false, 
    login: async (token) => { 
        await setCookie(KEY_AUTH_TOKEN, token.auth_token)
        await setCookie(KEY_REFRESH_TOKEN, token.refresh_token)
    }, 
    logout: () => {
        deleteCookie(KEY_AUTH_TOKEN)
        deleteCookie(KEY_REFRESH_TOKEN)
    } });

export const AuthProvider = (props: {
    children?: ReactNode 
    config?: {
        login_path?: string,
        redirect_to?: string,
    },
}) => {
    const router = useRouter()
    const { redirect } = router.query
    
    const [isAuth, setIsAuth] = useState(false)
    const [token, setToken] = useState<UserToken>()

    const login_path =  props.config ? props.config.login_path ? props.config.login_path : '/login' : '/login'
    const redirect_path = props.config ? props.config.redirect_to ? props.config.redirect_to : '/' : '/'
    
    
    useEffect(() => {
        const auth_token = getCookie(KEY_AUTH_TOKEN)
        const refresh_token = getCookie(KEY_REFRESH_TOKEN)
    
        if (auth_token) {
            setIsAuth(true)
            setToken({auth_token: auth_token, refresh_token: refresh_token ?? ''})
            if (router.pathname === login_path) {
                router.push(redirect?.toString() ?? redirect_path)
            } else if(router.pathname === "/register") {
                router.push(redirect?.toString() ?? redirect_path)
            }
        } else {
            if(router.pathname.startsWith("/app")) {
                router.push(login_path +'?redirect='+encodeURI(location.pathname))
            } 
        }
    }, [])

    const login = async (token: UserToken) => {
        setCookie(KEY_AUTH_TOKEN, token.auth_token)
        setCookie(KEY_REFRESH_TOKEN, token.refresh_token)
        setIsAuth(true)
        setToken(token)
        router.push(redirect?.toString() ?? redirect_path)
    }

    const logout = () => {
        deleteCookie(KEY_AUTH_TOKEN)
        deleteCookie(KEY_REFRESH_TOKEN)
        setIsAuth(false)
        setToken(undefined)
        router.push(login_path)
    }

    return (
        <AuthContext.Provider value={{ 
            isAuthenticated: isAuth, 
            token:token,
            login, 
            logout, 
        }}>
            {props.children}
        </AuthContext.Provider>
    )
}

export const useAuth = () => useContext(AuthContext)
