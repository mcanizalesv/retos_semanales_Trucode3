export default class kitties_collection extends HTMLElement {
  constructor() {
    super();
    this.root = this.attachShadow({ mode: "open" });
    const container = document.createElement("div");
    const kitties = app.store.kitties;

    fetch("components/kitties_collection.css")
      .then((res) => res.text())
      .then((text) => {
        const style = document.createElement("style");
        style.textContent = text;
        this.root.appendChild(style);
      });

    fetch("components/kitty_card.css")
      .then((res) => res.text())
      .then((text) => {
        kitties.forEach((kitty, index) => {
          const kittyCard = document.createElement("kitty-card"); //creo instancia de documento kitty-card
          const style = document.createElement("style");
          style.textContent = text;
          kittyCard.shadowRoot.appendChild(style);
          kittyCard.dataset.detail = JSON.stringify({ kitty, index });
          container.appendChild(kittyCard);
        });
      });

    this.root.appendChild(container);
  }
}

customElements.define("kitty-card-show", kitties_collection);
