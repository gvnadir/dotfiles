import Link from "next/link";
import Header from "./components/header";

// this component is converted in HTML and sent to the browser from the Server
export default function Home() {
  console.log("You won't see me in the browser, but in the server console!");
  return (
    <main>
      <Header />
      <p>🔥 Let&apos;s get started! 🔥</p>
      <Link href="/about">About Us</Link>
    </main>
  );
}
