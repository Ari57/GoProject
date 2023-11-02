function UpdateFormAction() {
    // Get the input element by its id
    var nameInput = document.getElementById("name");
  
    // Get the value entered by the user
    var nameValue = nameInput.value;
  
    // Update the form action with the value from the input
    var form = document.forms["name-form"];
    form.action = "https://goformtrigger.azurewebsites.net/api/FormTrigger?name=" + nameValue;
  
    // Submit the form
    form.submit();
  }


/*
var nameContainer = document.querySelector(".current-name")

async function GetCounter() {
        let response = await fetch("https://cors-anywhere.herokuapp.com/https://countertrigger.azurewebsites.net/api/http_trigger")
        let counter = await response.json()
        return counter;
    }

nameContainer.innerHTML = "Ye"

GetCounter().then(counter => nameContainer.innerHTML = "Ye")
    
const check = (e) => {
    const form = new FormData(e.target)
    const name = form.get("fname")
    console.log(name)
    console.log("----------------------------------")

    return fetch("https://goformtrigger.azurewebsites.net/api/FormTrigger?name= " + name)
        .then(data => nameContainer.innerHTML = "Ye")
}
*/