// ---------------------------

// EXAMPLE 1
// simple callback
function fetchData(callback) {
  callback("calling from the future");
}

fetchData((res) => {
  console.log(res);
});

// ---------------------------

// ---------------------------
// EXAMPLE 2

// real callback (executed in the future)
// function fetchData(callback) {
//   setTimeout(() => callback("calling from the future"), 2000);
// }

// fetchData((res) => {
//   console.log(res);
// });

// ---------------------------

// ---------------------------
// EXAMPLE 3
// callback hell

function fetchData(callback) {
  setTimeout(() => callback("calling from the future"), 2000);
}

fetchData((res) => {
  console.log(res);
  fetchData((res) => {
    console.log(res);
    fetchData((res) => {
      console.log(res);
    });
  });
});
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
