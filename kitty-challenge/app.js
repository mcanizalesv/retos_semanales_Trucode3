import Store from "./services/Store.js";
// Don't forget to import web components
import "./components/kitties_collection.js";
import "./components/kitty_card.js";
import "./components/kitty_modal.js";

document.addEventListener("DOMContentLoaded", () => {
  // WHat you think you could add in here?

  //Kitties container
  const main = document.querySelector("main");
  const cardEl = document.createElement("kitty-card-show");
  main.appendChild(cardEl);

  //Kitty modal

  const modal = document.createElement("kitty-modal");
  document.body.appendChild(modal);

  window.addEventListener("kitty-selected", () => {
    const selectedKitty = app.store.selectedKitty;
    const modal = document.querySelector("kitty-modal");

    if (selectedKitty) {
      modal.setKittyData(selectedKitty);
      modal.classList.add("show");
    }
  });
});

// Do not change this!
window.app = {};
app.store = Store;
