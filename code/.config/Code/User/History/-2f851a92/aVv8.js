// This module only ever runs on the server, including for client-side navigations.
// In a real scenario I would probably replace the data.json with a GET request to a REST API
// and put the request inside the load() function

import { posts } from "./data";

export function load() {
  return {
    summaries: posts.map((post) => ({
      slug: post.slug,
      title: post.title,
    })),
  };
}
