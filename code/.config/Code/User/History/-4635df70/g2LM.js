import classes from "./page.module.css";

export default function MealsPage() {
  return (
    <>
      <header className={classes.header}>
        <h1>
          Delicious meals, create{" "}
          <span className={classes.hightlight}>by you</span>
        </h1>
      </header>
      <main></main>
    </>
  );
}
