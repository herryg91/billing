import { createContext, useState, useContext, useEffect, ReactNode } from 'react'
import { ApiErrorResponse } from '../api/response';
import { useAuth } from './useAuth';
import { AuthenticatedUser } from '@/repositories/auth-api/entity';
import { AuthApi } from '@/repositories/auth-api';

type user_ctx = {
    user?: AuthenticatedUser
    mutate: () => void
};

const UserContext = createContext<user_ctx>({mutate: () => {}});

export const UserProvider = (props: {
    children?: ReactNode 
}) => {
    const { isAuthenticated, logout } = useAuth()
    const [currentUser, setCurrentUser] = useState<AuthenticatedUser>()

    const fetch_current_user = async () => {
        try {
            const resp = await AuthApi.GetAuthenticated()
            setCurrentUser(resp.data)
          } catch (error) {
            if(error as ApiErrorResponse){
                const api_err = error as ApiErrorResponse
                if(api_err.http_status === 404 && isAuthenticated) {
                    logout()
                }
            } else {
              console.log("Unknown error:",error);
            }
          } finally {
          }
    }

    useEffect(() => {
        if(isAuthenticated){
            fetch_current_user()
        }
    }, [isAuthenticated])

    return (
        <UserContext.Provider value={{ user: currentUser, mutate: fetch_current_user }}>
            {props.children}
        </UserContext.Provider>
    )
}

export const useUser = () => useContext(UserContext)
