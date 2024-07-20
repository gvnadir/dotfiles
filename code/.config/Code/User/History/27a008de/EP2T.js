import "./globals.css";

// metadata is a reserved name used for set metadata
// instead of adding the <head></head> section in the JSX below
export const metadata = {
  title: "NextJS Course App",
  description: "Your first NextJS app!",
};

export default function RootLayout({ children }) {
  return (
    <html lang="en">
      <body>{children}</body>
    </html>
  );
}
