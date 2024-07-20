<script>
  // Gets the data from +layout.server
  export let data;
</script>

<h1>blog</h1>

{#each data.summaries as { slug, title }}
  <li><a href="/blog/{slug}">{title}</a></li>
{/each}
