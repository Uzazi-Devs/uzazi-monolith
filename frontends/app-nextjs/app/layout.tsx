import "./globals.css";

export const metadata = { title: "uzazi" };

export default function RootLayout({ children }: { children: React.ReactNode }) {
  return (
    <html lang="en">
      <body className="m-0">{children}</body>
    </html>
  );
}
