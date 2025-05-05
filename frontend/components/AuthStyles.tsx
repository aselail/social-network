"use client";

import React from "react";

export const AuthInput = ({
  label,
  id,
  type = "text",
  value,
  onChange,
  required = false,
  ...props
}: {
  label: string;
  id: string;
  type?: string;
  value: string;
  onChange: (e: React.ChangeEvent<HTMLInputElement>) => void;
  required?: boolean;
  [key: string]: any;
}) => {
  return (
    <div>
      <label htmlFor={id} className="block text-sm font-medium text-gray-700">
        {label}
      </label>
      <input
        id={id}
        type={type}
        value={value}
        onChange={onChange}
        required={required}
        className="mt-1 block w-full rounded-md border border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm"
        {...props}
      />
    </div>
  );
};

export const AuthSelect = ({
  label,
  id,
  value,
  onChange,
  options,
  required = false,
  ...props
}: {
  label: string;
  id: string;
  value: string;
  onChange: (e: React.ChangeEvent<HTMLSelectElement>) => void;
  options: { value: string; label: string }[];
  required?: boolean;
  [key: string]: any;
}) => {
  return (
    <div>
      <label htmlFor={id} className="block text-sm font-medium text-gray-700">
        {label}
      </label>
      <select
        id={id}
        value={value}
        onChange={onChange}
        required={required}
        className="mt-1 block w-full rounded-md border border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm"
        {...props}
      >
        {options.map((option) => (
          <option key={option.value} value={option.value}>
            {option.label}
          </option>
        ))}
      </select>
    </div>
  );
};

export const AuthTextarea = ({
  label,
  id,
  value,
  onChange,
  rows = 2,
  required = false,
  ...props
}: {
  label: string;
  id: string;
  value: string;
  onChange: (e: React.ChangeEvent<HTMLTextAreaElement>) => void;
  rows?: number;
  required?: boolean;
  [key: string]: any;
}) => {
  return (
    <div>
      <label htmlFor={id} className="block text-sm font-medium text-gray-700">
        {label}
      </label>
      <textarea
        id={id}
        value={value}
        onChange={onChange}
        rows={rows}
        required={required}
        className="mt-1 block w-full rounded-md border border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm"
        {...props}
      />
    </div>
  );
};

export const AuthButton = ({
  children,
  type = "button",
  onClick,
  className,
  ...props
}: {
  children: React.ReactNode;
  type?: "button" | "submit" | "reset";
  onClick?: () => void;
  className?: string;
  [key: string]: any;
}) => {
  return (
    <button
      type={type}
      onClick={onClick}
      className={`w-full flex justify-center py-2 px-4 border border-transparent rounded-md shadow-sm text-sm font-medium text-white bg-indigo-600 hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500 ${
        className || ""
      }`}
      {...props}
    >
      {children}
    </button>
  );
};

export const AuthToggle = ({
  checked,
  onChange,
  label,
  ...props
}: {
  checked: boolean;
  onChange: (checked: boolean) => void;
  label: string;
  [key: string]: any;
}) => {
  return (
    <div className="flex items-center">
      <button
        type="button"
        onClick={() => onChange(!checked)}
        className={`${
          checked ? "bg-indigo-600" : "bg-gray-200"
        } relative inline-flex h-5 w-10 items-center rounded-full transition-colors focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:ring-offset-2`}
        {...props}
      >
        <span
          className={`${
            checked ? "translate-x-5" : "translate-x-1"
          } inline-block h-3 w-3 transform rounded-full bg-white transition-transform`}
        />
      </button>
      <span className="ml-2 text-sm text-gray-700">{label}</span>
    </div>
  );
};

export const AuthContainer = ({ children }: { children: React.ReactNode }) => {
  return (
    <div className="min-h-screen bg-gray-50 flex flex-col justify-center py-8 px-4">
      {children}
    </div>
  );
};

export const AuthCard = ({ children }: { children: React.ReactNode }) => {
  return (
    <div className="mt-6 mx-auto w-full max-w-sm">
      <div className="bg-white py-6 px-4 shadow rounded-lg">{children}</div>
    </div>
  );
};

export const AuthHeader = ({
  title,
  subtitle,
  link,
  linkText,
}: {
  title: string;
  subtitle?: string;
  link?: string;
  linkText?: string;
}) => {
  return (
    <div className="mx-auto w-full max-w-sm">
      <h2 className="text-center text-2xl font-bold text-gray-900">{title}</h2>
      {subtitle && (
        <div className="mt-2 text-center">
          {link ? (
            <a
              href={link}
              className="text-sm font-medium text-indigo-600 hover:text-indigo-500"
            >
              {linkText || subtitle}
            </a>
          ) : (
            <p className="text-sm text-gray-500">{subtitle}</p>
          )}
        </div>
      )}
    </div>
  );
};
