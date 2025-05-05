"use client";

import "../auth.css";

export default function AuthLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  return <div className="auth-pages">{children}</div>;
}
