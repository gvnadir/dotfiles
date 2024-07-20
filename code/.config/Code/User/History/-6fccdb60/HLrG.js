// ---------------------------

// EXAMPLE 1

// simple callback

// function fetchData(callback) {
//   callback("calling from the future");
// }

// fetchData((res) => {
//   console.log(res);
// });

// ---------------------------

// ---------------------------
// EXAMPLE 2

// real callback (executed in the future)
// bind for pre-configuration

function fetchData(callback) {
  // I preconfigured the callback by passing the argument without executing it!
  setTimeout(callback.bind(null, "calling from the future"), 2000);
}

fetchData((res) => {
  console.log(res);
});

// ---------------------------

// ---------------------------
// EXAMPLE 3

// real callback (executed in the future)
// arrow function to make it easier

function fetchData(callback) {
  // the arrow function allows to pass the argument to the function without calling it
  setTimeout(() => callback("calling from the future"), 2000);
}

fetchData((res) => {
  console.log(res);
});

// ---------------------------

// ---------------------------
// EXAMPLE 3
// the problem: callback hell

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
// EXAMPLE 4
// fix callback hell with promise

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
