const modal = document.querySelectorAll(".modal");
const openModal = document.querySelectorAll(".open-button");
const closeModal = document.querySelectorAll(".close-button");

const closeModal2 = document.querySelectorAll(".close-button");

// print hello world on the web page
console.log("Hello World");

openModal[0].addEventListener("click", () => {
  modal[0].showModal();
});

closeModal[0].addEventListener("click", () => {
  modal[0].close();
});

openModal[1].addEventListener("click", () => {
  modal[1].showModal();
});

closeModal[1].addEventListener("click", () => {
  modal[1].close();
});