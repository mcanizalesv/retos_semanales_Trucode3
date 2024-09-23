export default class kitty_card extends HTMLElement {
  constructor() {
    super();
    this.root = this.attachShadow({ mode: "open" });
    const template = document.querySelector("#kitty-card-template");
    const content = template.content.cloneNode(true);
    this.root.appendChild(content);
  }

  connectedCallback() {
    const dataFromParent = JSON.parse(this.dataset.detail);
    const kittyData = dataFromParent.kitty;
    const index = dataFromParent.index;

    this.root.addEventListener("click", () => {
      app.store.selectedKitty = kittyData;
    });
    this.root.querySelector("h2").innerText = kittyData.name;
    this.root.querySelector("h2").style.color =
      index % 2 === 0 ? "var(--magenta)" : "var(--christalle)";
    this.root.querySelector("img").src = kittyData.thumbnail;
    this.root.querySelector("p").innerText = `$${kittyData.price}`;
  }
}

customElements.define("kitty-card", kitty_card);
