import { useEffect } from 'react';
import { useRouter } from 'next/router';

const API_BASE_URL = 'http://localhost:8080'

export async function login(email: string, password: string): Promise<{  // Consider more specific type based on API response
    token: string;
    userData?: any; // Optional: Return user data from the API
    error?: string;
}> {
    try {
        const response = await fetch(`${API_BASE_URL}/api`, {  // Adjust endpoint
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({action: 'login', email, password}),
        });

        const data = await response.json();
        localStorage.setItem('token', data)
        console.log(data)

        if (!response.ok) {
            // Handle API errors (e.g., invalid credentials)
            return { error: data.message || 'Login failed' }; // Or a more helpful error message
        }

        if (!data.token) {
            return { error: 'Token missing in the response' };
        }


        return { token: data.token, userData: data.user }; // Assuming the API returns a token and user info
    } catch (error: any) {
        console.error('Login error:', error);
        return { error: error.message || 'An unexpected error occurred during login.' };
    }
}

export function logout(): void {
    localStorage.clear();
}


export function isUserLoggedIn(): boolean { // Allow optional context for server-side
    const cookie = localStorage.getItem('token');
    return !!cookie;
}


export function getAuthToken(context: any = null): string | null {
    return localStorage.getItem('token');
}

export function withAuth(Component: React.ComponentType<any>) { // Generic type for props

    return function AuthComponent(props: any) { // Generic type for props
        const router = useRouter();

        useEffect(() => {
            if (!isUserLoggedIn()) {
                router.replace('/login');
            }
        }, [router]);

        if (!isUserLoggedIn()) {
            return null; // Or a loading indicator
        }
        return <Component {...props} />;
    };
}

export async function apiRequest<T>(endpoint: string, body?: any, method: string = 'POST'): Promise<T> { // Generic type for return value
    const authToken = getAuthToken();
    const headers: { [key: string]: string } = {
        'Content-Type': 'application/json',
        ...(authToken ? { 'Authorization': `Bearer ${authToken}` } : {}),
    };

    try {
        const response = await fetch(`${API_BASE_URL}${endpoint}`, { // Adjust endpoint
            method,
            headers,
            body: body ? JSON.stringify(body) : undefined,
        });

        if (!response.ok) {
            // Handle API errors appropriately (e.g., 401 Unauthorized, 403 Forbidden)
            const errorData = await response.json();
            throw new Error(errorData.message || `API request failed with status ${response.status}`);
        }

        return await response.json();

    } catch (error: any) {
        console.error('API request error:', error);
        throw error; // Re-throw the error to be handled by the calling component
    }
}