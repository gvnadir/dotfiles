# Next.js

## What is Next.js

- Next.js is a **fullstack** React framework
- **Fullstacks Apps**: It vastly simplifies the process of building fullstack applications with React (FE and BE are part of the same project)
- **File-based Routing**: Routes are configured via the filesystem (folders + files)
- **Server-side Rendering**: By default, NextJS (pre-) renders all pages on the server. The finished HTML page is sent to the client (great for SEO)

## Two approaches for building NextJS apps

1. Pages Router
   - Has been around for many years
   - Very stable
2. App Router
   - Introduces with NextJS 13
   - Still realatively new and partially buggy

## Understand file-based routing and React server components

- **React Server Components**: Rendered only on the server, **never** on the client
  - It looks like a regular component, but NextJS runs it on the server (e.g. adding a `console.log` will only be shown in the server console, not the browser)
- NextJS relies on reserved, special filenames:
  - `page.js`: Define page content
  - `layout.js`: Define wrapper around pages
  - `not-found.js`: Define "Not Found" fallback page
  - `error.js`: Define "Error" fallback page
  - `icon.png`: The favicon

## Pre-fetching

- It is possible to use both client-side code and server-side code using NextJS

```JavaScript
// A simple anchor button that will reload the entire page by requesting an HTML page for the specified route

<a href="/about">About Us</a>

// Using Link from NextJS will allow to pre-fetching the content of the page when hovering over the button

// Furthemore, it will use client-side rendering for showing the new page, it means that just the JavaScript code for showing that page will be requested and showed without refreshing the entire page with a new HTML page

import Link from "next/link";

<Link href="/about">About Us</Link>
```

## Import from the root folder

- It is possible to point to the root folder of the project using the `@` symbol

> e.g. `import Header from "@/components/header";`

## Dynamic route syntax

- When we have a dynamic path without knowing its value, it is possible to define a placeholder using the `[]` syntax when defining the folder containing the page.js

> e.g.
>
> - app
>   - blog --> localhost/blog
>     - [slug]
>       - page.js --> localhost/blog/post-1

- The dynamic path can be retrieved using the `params` prop

```JavaScript
// blog/[slug]/page.js

function BlogPost( { params }) {
	return (
		<main>
			<h1>Blog Post</h1>
			<p>{params.slug}</p>
		</main>
	)
}
```

## Output images in a more optimized way

- Achieved by using the NextJS `<Image>` component
- Automatically _Lazy loading_ images under the hood (the image is loaded only when necessary)
- Use the `priority` prop to the `Image` component to ensure the image is loaded as fast as possibile (e.g. an image in the header - always visible)
- Use the `fill` prop when you need to display some images without knowing its dimention in advance

## Server Components vs Client Components

- NextJS components (routes, components, layouts etc) are server-side components by default (they are only rendered on the server)
- Components that uses effects, events or every peace of code that runs on the browser, must be declared as client-side components by declaring `'use client'` in the JS/TS file.

## Server Actions

1. Applied using the `'use server'` directive
2. Used inside functions to tell NextJS that the function is only executed on a server
3. The `async` keyboard is also necessary
4. The component in which the function is declared must be a server component

```JavaScript
async function serverFun() {
	'use server';
	// some code ...
}
```

- It's also a best practice to outsource the server side function to a separate file.
- By declaring `'use server'` at the top of the file, every functions in the file will be a server function.
