document.getElementById("button").addEventListener("click", () => {
  
  let result = "";
  document.querySelectorAll("ul > li:first-child").forEach(({ innerHTML }) => result += innerHTML );
  alert(result);

});