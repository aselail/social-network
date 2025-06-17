'use client'

import {isUserLoggedIn, Login, Logout} from '@/hooks/Auth'
import {useRouter} from 'next/router'
import {createContext, ReactNode, useCallback, useContext, useEffect, useState} from 'react'

interface AuthContextType {
  isAuthenticated: boolean
  // login: (email: string, password: string) => void
  // logout: () => void
  update: () => void
}

const AuthContext = createContext<AuthContextType | undefined>(undefined)

export const AuthProvider = ({children}: {children: ReactNode}) => {
  const [isAuthenticated, setIsAuthenticated] = useState<boolean>(isUserLoggedIn())

  const login = useCallback((email: string, password: string) => {
    Login(email, password)
    setIsAuthenticated(true)
  }, [])

  const logout = useCallback(() => {
    Logout()
    setIsAuthenticated(false)
  }, [])

  const update = useCallback(() => {
    setIsAuthenticated(isUserLoggedIn())
  }, [])

  return <AuthContext.Provider value={{isAuthenticated, /* login, logout, */ update}}>{children}</AuthContext.Provider>
}

export const useAuth = (): AuthContextType => {
  const context = useContext(AuthContext)
  if (context === undefined) {
    throw new Error('useAuth must be used within an AuthProvider')
  }
  return context
}

export function WithAuth(Component: React.ComponentType<any>) {
  return function AuthComponent(props: any) {
    const {isAuthenticated} = useAuth()
    const router = useRouter()

    useEffect(() => {
      if (!isAuthenticated) {
        router.replace('/login')
      }
    }, [router])

    if (!isUserLoggedIn()) {
      return null
    }
    return <Component {...props} />
  }
}
