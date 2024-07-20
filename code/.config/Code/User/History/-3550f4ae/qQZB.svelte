<script>
	let count = 0;

	function increment() {
		count += 1;
	}
</script>

<h1>Welcome to my Svelte web site!</h1>
<button on:click={increment}>
	Clicked {count}
</button>

