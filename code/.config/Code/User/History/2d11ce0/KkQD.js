function fetchData(callback) {
  callback("data");
}
fetchData((res) => {
  console.log(res);
});
