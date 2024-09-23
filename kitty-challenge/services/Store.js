// Do not change this!
const Store = {
  kitties: [
    {
      id: "ky-0",
      name: "Kitty 1",
      thumbnail: "https://www.cryptokitties.co/images/landing-kitty01.svg",
      price: 150,
    },
    {
      id: "ky-1",
      name: "Kitty 2",
      thumbnail: "https://www.cryptokitties.co/images/landing-kitty03.svg",
      price: 250,
    },
    {
      id: "ky-2",
      name: "Kitty 3",
      thumbnail: "https://www.cryptokitties.co/images/landing-kitty05.svg",
      price: 250,
    },
    {
      id: "ky-3",
      name: "Kitty 4",
      thumbnail: "https://www.cryptokitties.co/images/landing-kitty07.svg",
      price: 250,
    },
    {
      id: "ky-4",
      name: "Kitty 5",
      thumbnail: "https://www.cryptokitties.co/images/landing-kitty09.svg",
      price: 250,
    },
    {
      id: "ky-5",
      name: "Kitty 6",
      thumbnail: "https://www.cryptokitties.co/images/landing-kitty11.svg",
      price: 250,
    },
    {
      id: "ky-6",
      name: "Kitty 7",
      thumbnail: "https://www.cryptokitties.co/images/landing-kitty13.svg",
      price: 250,
    },
    {
      id: "ky-7",
      name: "Kitty 8",
      thumbnail: "https://www.cryptokitties.co/images/landing-kitty15.svg",
      price: 250,
    },
  ],
  selectedKitty: null,
};

const proxiedStore = new Proxy(Store, {
  set(target, property, value) {
    target[property] = value;
    if (property === "selectedKitty") {
      if (value) {
        window.dispatchEvent(new Event("kitty-selected"));
      } else {
        window.dispatchEvent(new Event("kitty-clear"));
      }
    }
    return true;
  },
  get(target, property) {
    return target[property];
  },
});

export default proxiedStore;
