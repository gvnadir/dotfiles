function fetchData(callback) {
  setTimeout(() => callback("calling from the future"), 2000);
}

// fetchData(value => {
//   return value;
// });

// callback hell

fetchData((res) => {
  console.log(res);
  fetchData((res) => {
    console.log(res);
    fetchData((res) => {
      console.log(res);
    });
  });
});

// const fetchData = () => {
//   const promise = new Promise((resolve, reject) => {
//     setTimeout(() => {
//       resolve("calling from the future");
//     }, 2000);
//   });
//   return promise;
// };

// solution to callback hell

// fetchData()
//   .then((text) => {
//     console.log(text);
//     return fetchData();
//   })
//   .then((text) => {
//     console.log(text);
//   });
