"use client";

import { useState } from "react";
import { useRouter } from "next/navigation";
import Link from "next/link";
import {
  AuthContainer,
  AuthHeader,
  AuthCard,
  AuthInput,
  AuthButton,
  AuthSelect,
  AuthTextarea,
  AuthToggle,
} from "@/components/AuthStyles";
import ImageUpload from "@/components/ImageUpload";

const RegisterPage = () => {
  const [firstName, setFirstName] = useState("");
  const [lastName, setLastName] = useState("");
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [dateOfBirth, setDateOfBirth] = useState("");
  const [gender, setGender] = useState("");
  const [nickname, setNickname] = useState("");
  const [about, setAbout] = useState("");
  const [imagePreview, setImagePreview] = useState<string | null>(null);
  const [isPublic, setIsPublic] = useState(true);
  const [error, setError] = useState<string | null>(null);
  const router = useRouter();

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setError(null);
    try {
      const response = await fetch(
        `${process.env.NEXT_PUBLIC_API_URL}/register`,
        {
          method: "POST",
          headers: { "Content-Type": "application/json" },
          body: JSON.stringify({
            first_name: firstName,
            last_name: lastName,
            email,
            password,
            date_of_birth: dateOfBirth,
            gender,
            nickname,
            about,
            profile_public: isPublic,
            profile_image: imagePreview,
          }),
        }
      );
      if (!response.ok) {
        const errText = await response.text();
        throw new Error(errText);
      }
      router.push("/login");
    } catch (err: any) {
      setError(err.message);
    }
  };

  return (
    <AuthContainer>
      <AuthHeader
        title="Create your account"
        subtitle="Already have an account? Login"
        link="/login"
      />

      <AuthCard>
        <form className="space-y-4" onSubmit={handleSubmit}>
          {error && (
            <div className="bg-red-50 border-l-4 border-red-400 p-3">
              <p className="text-red-700 text-sm">{error}</p>
            </div>
          )}

          <div className="grid grid-cols-2 gap-3">
            <AuthInput
              label="First Name"
              id="firstName"
              value={firstName}
              onChange={(e) => setFirstName(e.target.value)}
              required
            />
            <AuthInput
              label="Last Name"
              id="lastName"
              value={lastName}
              onChange={(e) => setLastName(e.target.value)}
              required
            />
          </div>

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

          <div className="grid grid-cols-2 gap-3">
            <AuthInput
              label="Date of Birth"
              id="dateOfBirth"
              type="date"
              value={dateOfBirth}
              onChange={(e) => setDateOfBirth(e.target.value)}
            />
            <AuthSelect
              label="Gender"
              id="gender"
              value={gender}
              onChange={(e) => setGender(e.target.value)}
              options={[
                { value: "", label: "Select" },
                { value: "Male", label: "Male" },
                { value: "Female", label: "Female" },
              ]}
            />
          </div>

          <AuthInput
            label="Nickname"
            id="nickname"
            value={nickname}
            onChange={(e) => setNickname(e.target.value)}
          />

          <ImageUpload onImageChange={setImagePreview} />

          <AuthTextarea
            label="About Me"
            id="about"
            value={about}
            onChange={(e) => setAbout(e.target.value)}
            rows={2}
          />

          <AuthToggle
            checked={isPublic}
            onChange={setIsPublic}
            label={isPublic ? "Public Profile" : "Private Profile"}
          />

          <AuthButton type="submit">Sign Up</AuthButton>
        </form>
      </AuthCard>
    </AuthContainer>
  );
};

export default RegisterPage;
