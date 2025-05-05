"use client";
// This is a client-side component for the login page
import { useState } from "react";
import { useRouter } from "next/navigation";
import Link from "next/link";
import {
  AuthContainer,
  AuthHeader,
  AuthCard,
  AuthInput,
  AuthButton,
} from "@/components/AuthStyles";

const LoginPage = () => {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [error, setError] = useState<string | null>(null);
  const router = useRouter();

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setError(null);
    try {
      const response = await fetch(`${process.env.NEXT_PUBLIC_API_URL}/login`, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        // credentials set to include cookies
        credentials: "include",
        body: JSON.stringify({ email, password }),
      });
      if (!response.ok) {
        const errText = await response.text();
        throw new Error(errText);
      }
      // If successful, navigate to home or profile
      router.push("/");
    } catch (err: any) {
      setError(err.message);
    }
  };

  return (
    <AuthContainer>
      <AuthHeader
        title="Login to your account"
        subtitle="Need an account? Sign up"
        link="/register"
      />

      <AuthCard>
        <form className="space-y-4" onSubmit={handleSubmit}>
          {error && (
            <div className="bg-red-50 border-l-4 border-red-400 p-3">
              <p className="text-red-700 text-sm">{error}</p>
            </div>
          )}

          <AuthInput
            label="Email"
            id="email"
            type="email"
            value={email}
            onChange={(e) => setEmail(e.target.value)}
            required
          />

          <AuthInput
            label="Password"
            id="password"
            type="password"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
            required
          />

          <AuthButton type="submit">Login</AuthButton>
        </form>
      </AuthCard>
    </AuthContainer>
  );
};

export default LoginPage;
