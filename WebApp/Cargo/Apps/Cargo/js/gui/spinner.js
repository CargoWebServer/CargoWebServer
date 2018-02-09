/**
 * Display a spinner.
 * @param {*} parent 
 * @param {*} radius 
 */
var Spinner = function(parent, radius){
    this.panel = parent.appendElement({ "tag": "div", "id":"waitingDiv" }).down();

    var spinner = new SVG_Element(null, "spinner", "spinner", "svg");
    this.radius = radius;
    if(this.radius == undefined){
        this.radius = 30;
    }

    spinner.setSvgAttribute("viewBox", "0 0 " + (2  * this.radius + 6) + " " +  (2  * this.radius + 6) + "" );
    spinner.setSvgAttribute("width", 2  * this.radius + 5 + "px");
    spinner.setSvgAttribute("height", 2  * this.radius + 5 + "px");

    var ring = new SVG_Circle( spinner, "spin_circle", "ring", "50%", "50%", "45%");
    ring.setSvgAttribute("fill", "none");
    ring.setSvgAttribute("stroke-width", "10%");
    ring.setSvgAttribute("stroke-linecap", "butt");

    var circle = new SVG_Circle( spinner, "spin_circle", "path", "50%", "50%", "45%");
    circle.setSvgAttribute("fill", "none");
    circle.setSvgAttribute("stroke-width", "10%");
    circle.setSvgAttribute("stroke-linecap", "butt");
    
    this.panel.appendElement(spinner)
    return this
}