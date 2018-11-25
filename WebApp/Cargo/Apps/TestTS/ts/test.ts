function Greeter(greething: string){
    this.greething = greething;
}

Greeter.prototype.greet = function(){
    return "Hello, " + this.greething;
}

let greether = new Greeter("world");

let button = document.createElement('button');

button.textContent = "Say Hello";

button.onclick = function(){
    alert(greether.greet());
}

document.body.appendChild(button)