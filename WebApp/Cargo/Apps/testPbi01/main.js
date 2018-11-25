// Load the interface.
function main(){
    var div = document.createElement("div");
    document.getElementsByTagName("body")[0].appendChild(div);
    var toto = new powerbi.extensibility.visual.Visual({"element":div});
    console.log("---> Im loaded!", toto);
}