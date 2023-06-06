let burgerIsOpen = false;

function openBurger() {
  let burgerNavContainer = document.getElementById("hamburger-nav-container");
  if (!burgerIsOpen) {
    burgerNavContainer.style.display = "block";
    burgerIsOpen = true;
    console.log(burgerIsOpen);
    console.log(burgerNavContainer);
  } else {
    burgerNavContainer.style.display = "none";
    burgerIsOpen = false;
  }
}

