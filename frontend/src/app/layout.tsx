import type { Metadata } from "next";
import { Geist, Geist_Mono } from "next/font/google";
import Link from "next/link";
import "./globals.css";

const geistSans = Geist({
  variable: "--font-geist-sans",
  subsets: ["latin"],
});

const geistMono = Geist_Mono({
  variable: "--font-geist-mono",
  subsets: ["latin"],
});

export const metadata: Metadata = {
  title: "Claude Task App",
  description: "シンプルなタスク管理アプリ",
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="ja">
      <body className={`${geistSans.variable} ${geistMono.variable} antialiased bg-gray-50 min-h-screen`}>
        <nav className="border-b border-gray-200 bg-white px-6 py-3">
          <div className="flex items-center gap-6">
            <Link href="/" className="text-lg font-semibold text-gray-900 hover:text-blue-600">
              Claude Task App
            </Link>
            <Link href="/tasks" className="text-sm text-gray-600 hover:text-gray-900">
              タスク一覧
            </Link>
            <Link href="/tasks/new" className="text-sm text-gray-600 hover:text-gray-900">
              タスク登録
            </Link>
            <Link href="/users/new" className="text-sm text-gray-600 hover:text-gray-900">
              ユーザー登録
            </Link>
          </div>
        </nav>
        <main>{children}</main>
      </body>
    </html>
  );
}
