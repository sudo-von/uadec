$("#show-button").click(function(){
  alert("Text: " + $("#animal").text());
});

$("#change-button").click(function(){
  const newAnimal = $("#animal").text() === "🐕" ? "🐈" : "🐕"; 
  $("#animal").text(newAnimal);
});