const formWrapper = document.getElementById("form-wrapper");

const editForm = document.getElementById("edit-form");
editForm.addEventListener("click", () => {
  formWrapper.style.display = "block";
});

const form = document.getElementById("identity-form");
form.addEventListener("submit", (e) => {
  e.preventDefault();
  let fieldValues = document.getElementsByClassName("form-control");
  let identityValues = document.getElementsByClassName("identity-value");
  Array.from(fieldValues).forEach((element, index) => {
    identityValues[index].innerHTML = element.value;
  });
  formWrapper.style.display = "none";
});
