[![Review Assignment Due Date](https://classroom.github.com/assets/deadline-readme-button-22041afd0340ce965d47ae6ef1cefeee28c7c493a6346c4f15d667ab976d596c.svg)](https://classroom.github.com/a/MAPz3zZh)
# The Kitty Store Instructions

---

Using web components, generate the components indicated in `index.html` so that the application matches the screenshots found in the corresponding folder. Don't forget to use `template`, `class`, `defer`, `module=type`, and `DOMContentLoaded`.

## Views

There will be only one view, which will display all the kittens available in the store, ordered as follows:

| Size          | Columns   |
| ------------- | --------- |
| 0-424px       | 1 column  |
| 425px-767px   | 2 columns |
| 768px-1023px  | 3 columns |
| 1024px-1199px | 4 columns |
| 1200px-max    | 5 columns |

The maximum width for the container of the kitten list is 1200px.
[screenshots](./screenshots)

> 💡 Use CSS variables available at base.css.

## Behaviors

When a kitten is clicked, a modal with its details will be displayed. If the screen width is less than 768px, the format will be vertical, 90% wide by 80% high; otherwise, it will be horizontal in a rectangle of 500px wide by 320px high.

## services/Store

You have a `store` that contains the necessary properties and data to render the corresponding views.

This store is part of a global variable called `app`, which means you can use it in any script as the source of truth.

### Events

This store triggers two events that respond to the following conditions:

- `kitty-selected`: When a value is assigned to the `selectedKitty` property.
- `kitty-clear`: When a null value is assigned to the `selectedKitty` property.

> 💡 It is recommended to use these events in order to show/hide the modal.
