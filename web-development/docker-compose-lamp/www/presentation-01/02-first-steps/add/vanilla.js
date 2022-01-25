const animal = document.getElementById("animal");

document.getElementById("append-button").addEventListener("click",function(){
  animal.append("🐈");
});

document.getElementById("prepend-button").addEventListener("click",function(){
  animal.prepend("🐈‍⬛");
});

document.getElementById("after-button").addEventListener("click",function(){
  animal.after("🐱");
});

document.getElementById("before-button").addEventListener("click",function(){
  animal.before("🙀");
});