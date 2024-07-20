// This module only ever runs on the server, including for client-side navigations.
// In a real scenario I would probably replace the data.json with a GET request to a REST API
// and put the request inside the load() function

import { error } from "@sveltejs/kit";
import { posts } from "../data.js";

export function load({ params }) {
  const post = posts.find((post) => post.slug === params.slug);

  if (!post) {
    throw error(404);
  }

  return {
    post,
  };
}
