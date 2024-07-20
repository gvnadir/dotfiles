import Image from "next/image";
import classes from "./page.module.css";

export default function MealDetailPage() {
  return (
    <header className={classes.header}>
      <div className={classes.image}></div>
      <Image fill />

      <div className={classes.headerText}>
        <h1>TITLE</h1>
        <p className={classes.creator}></p>
        <p></p>
      </div>
    </header>
  );
}
