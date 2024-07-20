const fetchData = (callback) => {
  setTimeout(callback.bind(null, "calling from the future", 42), 2000);
};

const printMessageFromTheFuture = (text, secretNumber) => {
  console.log(text);
  console.log("We got a secret number too: ", secretNumber);
};

fetchData(printMessageFromTheFuture);

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
// ---------------------------
// EXAMPLE 1

// Simple callback

function fetchData(callback) {
  callback("calling from the future");
}

fetchData((res) => {
  console.log(res);
});

// ---------------------------

// ---------------------------
// EXAMPLE 2

// Real callback use-case (executed in the future)
// Bind for pre-configuration

// The callback is pre-configured using "bind". This allows to pass an argument to the function without calling it

// function fetchData(callback) {
//   setTimeout(callback.bind(null, "calling from the future"), 2000);
// }

// fetchData((res) => {
//   console.log(res);
// });

// ---------------------------

// ---------------------------
// EXAMPLE 3

// Real callback use-case (executed in the future)
// Arrow function to make it easier

// The arrow function allows to pass the argument to the function without calling it
// This is actually achieved thanks to anonymous function, but it's better to stick with arrow functions
// to avoid undesired behavious when working with "this"

// function fetchData(callback) {
//   setTimeout(() => callback("calling from the future"), 2000);
// }

// fetchData((res) => {
//   console.log(res);
// });

// ---------------------------

// ---------------------------
// EXAMPLE 4

// The problem: callback hell

// function fetchData(callback) {
//   setTimeout(() => callback("calling from the future"), 2000);
// }

// fetchData((res) => {
//   console.log(res);
//   fetchData((res) => {
//     console.log(res);
//     fetchData((res) => {
//       console.log(res);
//       fetchData((res) => {
//         console.log(res);
//         fetchData((res) => {
//           console.log(res);
//           fetchData((res) => {
//             console.log(res);
//           });
//         });
//       });
//     });
//   });
// });

// ---------------------------

// ---------------------------
// EXAMPLE 5

// Fix callback hell with Promise
// "then"

// function fetchData() {
//   const promise = new Promise((resolve, reject) => {
//     setTimeout(() => {
//       resolve("calling from the future");
//     }, 2000);
//   });
//   return promise;
// }

// fetchData()
//   .then((res) => {
//     console.log(res);
//     return fetchData();
//   })
//   .then((res) => {
//     console.log(res);
//     return fetchData();
//   })
//   .then((res) => {
//     console.log(res);
//   });

// ---------------------------

// ---------------------------
// EXAMPLE 6

// Fix callback hell with Promise
// "async", "await"

// function fetchData() {
//   const promise = new Promise((resolve, reject) => {
//     setTimeout(() => {
//       resolve("calling from the future");
//     }, 2000);
//   });
//   return promise;
// }

// async function asyncFun() {
//   console.log(await fetchData());
//   console.log(await fetchData());
//   console.log(await fetchData());
//   console.log(await fetchData());
// }

// asyncFun();

// ---------------------------

// ---------------------------
// EXAMPLE 7

// Real use-case with some work behind the scene

// const count = {
//   value: 0,
//   update: function(callback) {
//     const currentValue = this.value;
//     const newValue = callback(currentValue);
//     this.value = newValue;
//     console.log("Value updated to:", newValue);
//   }
// };

// function increment() {
//   count.update((n) => n + 1); // Increment by 1
// }

// function decrement() {
//   count.update((n) => n - 1); // Decrement by 1
// }

// increment(); // Output: Value updated to: 1
// decrement(); // Output: Value updated to: 0

// ---------------------------
