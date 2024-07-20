function fetchData(callback) {
  setTimeout(() => callback("calling from the future", 42), 2000);
}

fetchData((text, secretNumber) => {
  return text + " " + secretNumber;
});

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

// fetchData((text) => {
//   console.log(text);
//   fetchData((text) => {
//     console.log(text);
//     fetchData((text) => {
//       console.log(text);
//     });
//   });
// });

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
