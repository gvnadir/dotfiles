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

### Understand file-based routing and React server components

- **React Server Components**: Rendered only on the server, **never** on the client
	- It looks like a regular component, but NextJS runs it on the server (e.g. adding a `console.log` will only be shown in the server console, not the browser)
- NextJS relies on reserved, special filenames:
	- `page.js`: Define page content
	- `layout.js`: Define wrapper around pages
	- `not-found.js`: Define "Not Found" fallback page
	- `error.js`: Define "Error" fallback page

### Pre-fetching

- It is possible to use both client-side code and server-side code using NextJS

```JavaScript
// A simple anchor button that will reload the entire page by requesting an HTML page for the specified route

<a href="/about">About Us</a>

// Using Link from NextJS will allow to pre-fetching the content of the page when hovering over the button

// Furthemore, it will use client-side rendering for showing the new page, it means that just the JavaScript code for showing that page will be requested and showed without refreshing the entire page with a new HTML page

import Link from "next/link";

<Link href="/about">About Us</Link>
```
