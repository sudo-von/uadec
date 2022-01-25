let parents = [];
let element = document.getElementById("son");
while(element.parentNode) {
  parents.unshift(element.parentNode);
  element = element.parentNode;
}
let position = parents.length-1;

document.getElementById("button").addEventListener("click", function(){
  parents[position].style.border = "2px solid red";
  parents[position].style.color = "red";
  position--;
});