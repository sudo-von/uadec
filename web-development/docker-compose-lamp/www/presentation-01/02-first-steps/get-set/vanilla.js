const showButton = document.getElementById("show-button");
const changeButton = document.getElementById("change-button");
const animal = document.getElementById("animal");

showButton.addEventListener("click", () => {
  alert(`Text: ${animal.innerText}`);
});

changeButton.addEventListener("click", () => {
  const newAnimal = animal.innerText === "🐕" ? "🐈" : "🐕"; 
  animal.innerText = newAnimal;
});