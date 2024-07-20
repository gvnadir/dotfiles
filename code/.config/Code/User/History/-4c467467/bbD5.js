// this component is converted in HTML and sent to the browser from the Server
export default function Home() {
  console.log("You won't see me in the browser, but in the server console!");
  return (
    <main>
      <img src="/logo.png" alt="A server surrounded by magic sparkles." />
      <h1>Welcome to this NextJS Course!</h1>
      <p>🔥 Let&apos;s get started! 🔥</p>
      <a href="/about">About Us</a>
    </main>
  );
}
