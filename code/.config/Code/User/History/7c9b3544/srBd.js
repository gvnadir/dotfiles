import logoImg from "@/assets/logo.png";
import Link from "next/link";

export default function MainHeader() {
  return (
    <header>
      <Link href="/">
        <img src={logoImg.src} alt="A plate with food on it" />
        NextLevel Food
      </Link>
    </header>
  );
}
