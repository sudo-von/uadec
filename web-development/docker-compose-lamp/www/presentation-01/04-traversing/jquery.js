let position = 0;

$("#button").click(function(){
  $("#son").parents()[position].style.border = "2px solid red";
  $("#son").parents()[position].style.color = "red";
  position++;
});