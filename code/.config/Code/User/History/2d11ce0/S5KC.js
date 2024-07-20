function fetchData() {
	return new Promise(resolve => {
		setTimeout(() => {resolve("data")}, [2000])
	})
}

// fetchData().then(res => console.log(res))

async function myFun() {
	let res = await fetchData();
	console.log(res);
}

myFun();