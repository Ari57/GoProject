function InsertName() {
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


function DeleteNames () {
    var xhr = new XMLHttpRequest();

    // Define the request method, URL, and set asynchronous to true
    xhr.open("GET", "https://goformtrigger.azurewebsites.net/api/FormTrigger?delete=yes", true);

    xhr.onload = function () {
      if (xhr.status >= 200 && xhr.status < 300) {
        // Request was successful
        console.log("Names deleted successfully.");
        window.location.href="https://goformtrigger.azurewebsites.net/api/FormTrigger?delete=yes"
      } else {
        // Request encountered an error
        console.error("Error deleting names:", xhr.status, xhr.statusText);
      }
    };

    // Send the request with an empty body
    xhr.send();
}


