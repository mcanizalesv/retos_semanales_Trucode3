export default class kitty_modal extends HTMLElement {
  constructor() {
    super();
    this.root = this.attachShadow({ mode: "open" });
    const template = document
      .querySelector("#kitty-modal-template")
      .content.cloneNode(true);
    this.root.appendChild(template);

    //CSS styles application
    fetch("components/kitty_modal.css")
      .then((res) => res.text())
      .then((text) => {
        const style = document.createElement("style");
        style.textContent = text;
        this.root.appendChild(style);
      });
  }

  closeModal() {
    this.root.querySelector(".modal").style.display = "none";
    this.root.querySelector(".background").classList.remove("show");
  }

  connectedCallback() {
    const background = this.root.querySelector(".background");
    if (background) {
      background.addEventListener("click", () => {
        this.closeModal();
      });
    } else {
      console.error("Background element not found");
    }
  }

  setKittyData(kitty) {
    this.root.querySelector(".background").classList.add("show");
    this.root.querySelector("#kitty-name").textContent = kitty.name;
    this.root.querySelector("#kitty-img img").src = kitty.thumbnail;
    this.root.querySelector("#kitty-price").textContent = `$${kitty.price}`;

    this.root.querySelector(".modal").style.display = "block";
  }
}

customElements.define("kitty-modal", kitty_modal);
