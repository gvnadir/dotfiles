// this component is converted in HTML and sent to the browser from the Server
export default function Home() {
  console.log("Hello from the server side!");
  return (
    <main>
      <img src="/logo.png" alt="A server surrounded by magic sparkles." />
      <h1>Welcome to this NextJS Course!</h1>
      <p>🔥 Let&apos;s get started! 🔥</p>
    </main>
  );
}
