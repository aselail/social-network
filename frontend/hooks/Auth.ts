'use client'

import {NextRouter, useRouter} from 'next/router'

const API_BASE_URL = 'http://localhost:8080'

type UserInfo = {
  id: number
  nickname: string
  email: string
  firstName: string
  lastName: string
  gender: string
  dob: string
}

export const UserInfo: UserInfo = {
  id: 0,
  nickname: '',
  email: '',
  firstName: '',
  lastName: '',
  gender: '',
  dob: '',
}

export async function Login(email: string, password: string): Promise<{UserInfo?: UserInfo; error?: string}> {
  try {
    const response = await fetch(`${API_BASE_URL}/api`, {
      // Adjust endpoint
      method: 'POST',
      headers: {'Content-Type': 'application/json'},
      body: JSON.stringify({action: 'login', email, password}),
    })

    const data = await response.json()
    localStorage.setItem('token', JSON.stringify(data))
    console.log(data)

    if (!response.ok) {
      // Handle API errors (e.g., invalid credentials)
      return {error: data.message || 'Login failed'} // Or a more helpful error message
    }

    if (!data.token) {
      return {error: 'Token missing in the response'}
    }

    return {UserInfo: data.user}
  } catch (error: any) {
    console.error('Login error:', error)
    return {error: error.message || 'An unexpected error occurred during login.'}
  }
}

type RegisterParams = {
  User: {
    Public: boolean
    Email: string
    FirstName: string
    LastName: string
    Gender?: number
    Dob: string
    Nickname: string
    About: string
  }
  Password: string
  ImageFilename?: string
  ImageMimetype?: string
  ImageData?: string
}

export async function Register(request: RegisterParams): Promise<{UserInfo?: UserInfo; error?: string}> {
  try {
    request = Object.assign(request, {action: 'signup'})

    const response = await fetch(`${API_BASE_URL}/api`, {
      // Adjust endpoint
      method: 'POST',
      headers: {'Content-Type': 'application/json'},
      body: JSON.stringify(request),
    })

    if (!response!.ok) {
      // Handle API errors (e.g., invalid credentials)
      return {error: (await response.text()) || 'Register failed'} // Or a more helpful error message
    }

    const data = await response!.json()
    localStorage.setItem('token', JSON.stringify(data))
    console.log(data)

    if (!data.token) {
      return {error: 'Token missing in the response'}
    }

    return {UserInfo: data.user}
  } catch (error: any) {
    console.error('Register error:', error)
    return {error: error.message || 'An unexpected error occurred during login.'}
  }
}

export function Logout(): void {
  localStorage.clear()
  cache = null
}

export function isUserLoggedIn(): boolean {
  return !!getAuthToken()
}

let cache: any = null

export function getAuthToken(): string | null {
  if (cache) return cache
  if (typeof window != 'undefined') cache = localStorage.getItem('token')

  return cache
}

export function getUserInfo(): UserInfo | null {
  const token = getAuthToken()

  return JSON.parse(token!).user
}

export async function apiRequest<T>(endpoint: string, body?: any, method: string = 'POST'): Promise<T> {
  // Generic type for return value
  const authToken = getAuthToken()
  const headers: {[key: string]: string} = {
    'Content-Type': 'application/json',
    ...(authToken ? {Authorization: `Bearer ${authToken}`} : {}),
  }

  try {
    const response = await fetch(`${API_BASE_URL}${endpoint}`, {
      // Adjust endpoint
      method,
      headers,
      body: body ? JSON.stringify(body) : undefined,
    })

    if (!response.ok) {
      // Handle API errors appropriately (e.g., 401 Unauthorized, 403 Forbidden)
      const errorData = await response.json()
      throw new Error(errorData.message || `API request failed with status ${response.status}`)
    }

    return await response.json()
  } catch (error: any) {
    console.error('API request error:', error)
    throw error // Re-throw the error to be handled by the calling component
  }
}

function useEffect(arg0: () => void, arg1: NextRouter[]) {
  throw new Error('Function not implemented.')
}
