import classes from "./page.module.css";

export default function MealsPage() {
  return (
    <>
      <header className={classes.header}>
        <h1>
          Delicious meals, create{" "}
          <span className={classes.hightlight}>by you</span>
        </h1>
        <p>
          Choose your favorite receipe and cook it yourself. It is easy and fun!
        </p>
      </header>
      <main className={classes.main}></main>
    </>
  );
}
