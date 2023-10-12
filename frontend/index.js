let form = document.forms["name-form"]
form.addEventListener("submit", getData)


function getData() {
var nameValue = document.getElementById("fname").value
console.log(nameValue)
}