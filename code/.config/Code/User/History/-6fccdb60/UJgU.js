function fetchData(callback) {
  setTimeout(() => callback("calling from the future", 42), 2000);
}

fetchData((text, secretNumber) => {
  console.log(text, secretNumber);
});

// const fetchData = (callback) => {
//   setTimeout(callback.bind(null, "calling from the future", 42), 2000);
// };

// const printMessageFromTheFuture = (text, secretNumber) => {
//   console.log(text);
//   console.log("We got a secret number too: ", secretNumber);
// };

// fetchData(printMessageFromTheFuture);

// const fetchData = (callback) => {
//   setTimeout(() => {
//     callback("calling from the future", 42);
//   }, 2000);
// };

// fetchData((text, secretNumber) => {
//   console.log(text, secretNumber);
// });

// callback hell

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
