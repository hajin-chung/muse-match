const dropDownButton = document.querySelector("#dropdown-button");
const dropDownImage = document.querySelector("#dropdown-button > img");
const dropDownMenu = document.querySelector("#dropdown-menu")

dropDownButton.addEventListener("click", () => {
  const display = dropDownMenu.style.display;
  if (display === "none" || display === "") {
    dropDownMenu.style.display = "flex";
    dropDownImage.src = "/icons/chevron-up.svg";
  } else {
    dropDownMenu.style.display = "none";
    dropDownImage.src = "/icons/chevron-down.svg";
  }
})
