"use client";

import Link from 'next/link';

export default function Home() {
  return (
    <main style={{ padding: "2rem", fontFamily: "Arial, sans-serif" }}>
      <h1>Welcome to Social Network</h1>
      <p>This is your feed. Explore posts, connect with friends, and chat in real time!</p>
      
      <nav style={{ marginTop: "1rem" }}>
        <ul style={{ listStyle: "none", padding: 0 }}>
          <li style={{ margin: "0.5rem 0" }}>
            <Link href="/login">Login</Link>
          </li>
          <li style={{ margin: "0.5rem 0" }}>
            <Link href="/register">Register</Link>
          </li>
          <li style={{ margin: "0.5rem 0" }}>
            <Link href="/profile">Profile</Link>
          </li>
          <li style={{ margin: "0.5rem 0" }}>
            <Link href="/chat">Chat</Link>
          </li>
        </ul>
      </nav>
    </main>
  );
}
