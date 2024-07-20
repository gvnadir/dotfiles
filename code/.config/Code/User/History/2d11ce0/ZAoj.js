function fetchData(callback) {
	setTimeout(() => {callback("data")}, [2000])
}

fetchData().then()