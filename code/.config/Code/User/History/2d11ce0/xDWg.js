function fetchData() {
	return new Promise(resolve => {
		setTimeout(() => {resolve("data")}, [2000])
	})
}

fetchData().then(res => console.log(res))