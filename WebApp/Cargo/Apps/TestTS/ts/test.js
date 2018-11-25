"use strict";
function Greeter(greething) {
    this.greething = greething;
}
Greeter.prototype.greet = function () {
    return "Hello, " + this.greething;
};
var greether = new Greeter("world");
var button = document.createElement('button');
button.textContent = "Say Hello";
button.onclick = function () {
    alert(greether.greet());
};
document.body.appendChild(button);
