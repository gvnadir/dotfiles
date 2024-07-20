// this module only ever runs on the server, including for client-side navigations.

import { posts } from "./data";

export function load() {
  return {
    summaries: posts.map((post) => ({
      slug: post.slug,
      title: post.title,
    })),
  };
}
