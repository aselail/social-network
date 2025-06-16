"use client";

import { useSearchParams } from "next/navigation";
import { useEffect, useState } from "react";

interface Profile {
  id: string;
  first_name: string;
  last_name: string;
  email: string;
  avatar?: string;
  nickname?: string;
  about_me?: string;
  created_at: string;
}

const ProfilePage = () => {
  const [profile, setProfile] = useState<Profile | null>(null);
  const [editing, setEditing] = useState(false);
  const [form, setForm] = useState({
    first_name: "",
    last_name: "",
    avatar: "",
    nickname: "",
    about_me: "",
  });
  const [error, setError] = useState<string | null>(null);

  const fetchProfile = async () => {
    try {
      const searchParams = useSearchParams();
      const id = searchParams.get('id');  // TODO use it
      const res = await fetch(`${process.env.NEXT_PUBLIC_API_URL}/profile`, {
        credentials: "include",
      });
      if (!res.ok) {
        throw new Error(await res.text());
      }
      const data = await res.json();
      setProfile(data);
      setForm({
        first_name: data.first_name,
        last_name: data.last_name,
        avatar: data.avatar || "",
        nickname: data.nickname || "",
        about_me: data.about_me || "",
      });
    } catch (err: any) {
      setError(err.message);
    }
  };

  useEffect(() => {
    fetchProfile();
  }, []);

  const handleChange = (e: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>) => {
    setForm({...form, [e.target.name]: e.target.value });
  };

  const handleSave = async () => {
    try {
      const res = await fetch(`${process.env.NEXT_PUBLIC_API_URL}/profile`, {
        method: "PUT",
        headers: { "Content-Type": "application/json" },
        credentials: "include",
        body: JSON.stringify(form),
      });
      if (!res.ok) {
        throw new Error(await res.text());
      }
      const data = await res.json();
      setProfile(data);
      setEditing(false);
    } catch (err: any) {
      setError(err.message);
    }
  };

  if (!profile) return <p>Loading...</p>;

  return (
    <div style={{ maxWidth: 600, margin: "0 auto", padding: "1rem" }}>
      <h1>Profile</h1>
      {error && <p style={{ color: "red" }}>{error}</p>}
      {editing ? (
        <div>
          <input
            name="first_name"
            value={form.first_name}
            onChange={handleChange}
            placeholder="First Name"
          />
          <input
            name="last_name"
            value={form.last_name}
            onChange={handleChange}
            placeholder="Last Name"
          />
          <input
            name="avatar"
            value={form.avatar}
            onChange={handleChange}
            placeholder="Avatar URL"
          />
          <input
            name="nickname"
            value={form.nickname}
            onChange={handleChange}
            placeholder="Nickname"
          />
          <textarea
            name="about_me"
            value={form.about_me}
            onChange={handleChange}
            placeholder="About Me"
          />
          <button onClick={handleSave}>Save</button>
          <button onClick={() => setEditing(false)}>Cancel</button>
        </div>
      ) : (
        <div>
          <p>
            <strong>Name:</strong> {profile.first_name} {profile.last_name}
          </p>
          <p>
            <strong>Email:</strong> {profile.email}
          </p>
          {profile.avatar && <img src={profile.avatar} alt="Avatar" width={100} />}
          {profile.nickname && <p><strong>Nickname:</strong> {profile.nickname}</p>}
          {profile.about_me && <p><strong>About Me:</strong> {profile.about_me}</p>}
          <button onClick={() => setEditing(true)}>Edit Profile</button>
        </div>
      )}
    </div>
  );
};

export default ProfilePage;
