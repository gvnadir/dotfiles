function fetchData(callback) {
	setTimeout(() => {
		callback("data");

	}, [2000])
}

fetchData((res) => {
  console.log(res);
});
