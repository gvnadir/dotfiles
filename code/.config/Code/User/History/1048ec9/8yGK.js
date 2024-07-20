import { posts } from "../data.js";

export function load({ params }) {
  const post = posts.find((post) => post.slug === params.slug);

  return {
    post,
  };
}
