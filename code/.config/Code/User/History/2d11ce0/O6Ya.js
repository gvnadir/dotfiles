function fetchData(callback) {
	setTimeout(() => {
		callback("data");
}

fetchData((res) => {
  console.log(res);
});
