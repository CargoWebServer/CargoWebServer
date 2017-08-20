/*
 * (C) Copyright 2016 Mycelius SA (http://mycelius.com/).
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 * 
 */

/**
 * @fileOverview SVG element wrapper.
 * @author Dave Courtois
 * @version 1.0
 */

/**
 * A simple wrapper around SVG element.
 * @extends EventHub
 * @param {Element} parent The parent of the svg element.
 * @param {string} id Identifier, must be unique in it context.
 * @param {string} className The syle sheet class name.
 * @param {string} type The element type to create.
 * @constructor
 */
var SVG_Element  = function(parent, id, className, type){

    // Call the parent constructor.
    Element.call(this, parent, null)

    // The namespace of sgv utility.
    this.ns = "http://www.w3.org/2000/svg";
    this.type = type;

    this.id = id;
    if(this.id == undefined){
        this.id = randomUUID()
    }

    this.className = className;

    // The underlying svg element.
    this.element = null; // This will link the svg element in the page.
    this.element = document.createElementNS(this.ns, type);
    this.element.setAttribute("id", id);
    this.element.setAttribute("class", className);

    this.subElements = {}

    // If the parent is not null I will append the element in it.
    if(parent!=null){
        this.parentElement.appendElement(this)
    }

    return this;
}

SVG_Element.prototype = Object.create(Element.prototype);
SVG_Element.constructor = SVG_Element;

/**
 * Return the parent canvas.
 * @returns {*} The parent canvas where the element is render.
 */
SVG_Element.prototype.getCanvas = function(){
    if(this.type == "svg"){
        return this
    }else{
        return this.parentElement.getCanvas()
    }
}

/**
 * Append a child element.
 * @returns {Element} Return the element itself, usefull to link function call.
 */
SVG_Element.prototype.appendElement = function(e){
    this.subElements[e.id] = e;
    this.element.appendChild(e.element);
    
    // Call function on the parent.
    return Element.prototype.appendElement.call(this, e)
}

/**
 * Set attribute value.
 * @param {string} name The attribute name
 * @param {*} value The attribute value
 */
SVG_Element.prototype.setSvgAttribute = function(name, value){
    this.element.setAttribute(name, value);
}

/**
 * Return the attribute value.
 * @param {string} name The name of the attribute to return
 * @returns The attribute value
 */
SVG_Element.prototype.getSvgAttribute = function(name){
    return this.element.getAttribute(name);
}

////////////////////////////////////////////////////////////////////////////
//  Transformations
////////////////////////////////////////////////////////////////////////////

/**
 * Translate the element.
 * @param {int} x The x distance
 * @param {int} y The y distance
 */
SVG_Element.prototype.translate = function(x, y){
    var transform = this.element.getAttribute("transform");
    if(transform == null){
        transform = "";
    }
    transform += "translate("+ x + " ," + y +") ";
    this.element.setAttribute("transform", transform);
}

/**
 * Rotate the element by a given angle from a given point.
 * @param {float} deg The angle in degree to rotate
 * @param {int} x The x coordinate
 * @param {int} y The y coordinate
 */
SVG_Element.prototype.rotate = function(deg, cx, cy){
    var transform = this.element.getAttribute("transform");
    if(transform == null){
        transform = "";
    }
    if(cx != undefined && cy != undefined){
        transform += "rotate("+ deg + "," + cx + "," + cy +") ";
    }else{
        transform += "rotate("+ deg +") "; 
    }

    this.element.setAttribute("transform", transform);
}

/**
 * Scale the element.
 * @param {float} fx The x axis scale factor.
 * @param {float} fy The y axis scale factor.
 */
SVG_Element.prototype.scale = function(fx, fy){
    var transform = this.element.getAttribute("transform");
    if(transform == null){
        transform = "";
    }
    transform += "scale("+ fx + " " + fy +") ";
    this.element.setAttribute("transform", transform);
}

/**
 * Skew a given element in x axis.
 * @param {float} x The angle of skewniess.
 */
SVG_Element.prototype.skewX = function(x){
    var transform = this.element.getAttribute("transform");
    if(transform == null){
        transform = "";
    }
    transform += "skewX("+ x +") ";
    this.element.setAttribute("transform", transform);
}

/**
 * Skew a given element in y axis.
 * @param {float} y The angle of skewniess.
 */
SVG_Element.prototype.skewY = function(y){
    var transform = this.element.getAttribute("transform");
    if(transform == null){
        transform = "";
    }
    transform += "skewY(" + y + ") ";
    this.element.setAttribute("transform", transform);
}

/**
 * The canvas contain all svg element, it can contain reference to other canvas.
 * @extends SVG_Element
 * @param {string} id The canvas id must be unique
 * @param {string} className  The Css class name if one is use.
 * @param {int} w  The width
 * @param {int} h  The height
 * @returns {*}
 * @constructor
 */
var SVG_Canvas = function(parent,id,className, w, h){
    SVG_Element.call(this, parent, id, className, "svg");

    // The svg element.
    this.element.setAttribute("width", w);
    this.element.setAttribute("height", h);
    
    // I will append the svg inside it's parent.
    this.parentElement.element.appendChild(this.element)

    return this;
}

SVG_Canvas.prototype = Object.create(SVG_Element.prototype);
SVG_Canvas.constructor = SVG_Canvas;

////////////////////////////////////////////////////////////////////////////
//  Basic shape
////////////////////////////////////////////////////////////////////////////

/**
 * Append a new line inside the canvas. The line style attribute must be set in the CSS.
 * @param parent The parent element, can be SVG_Canvas or SVG_Group
 * @param {string} id The id of the line
 * @param {string} className The class for the line
 * @param {int} startX The starting point of the line
 * @param {int} startY
 * @param {int} endX   The ending point of the line
 * @param {int} endY
 * @returns {*}
 * @extends SVG_Element
 * @constructor
 */
var SVG_Line = function( parent, id, className, startX, startY, endX, endY){
    SVG_Element.call(this, parent, id, className, "line");

    // The svg element.
    this.element.setAttribute("x1", startX);
    this.element.setAttribute("x2", endX);
    this.element.setAttribute("y1", startY);
    this.element.setAttribute("y2", endY);

    return this;
}

SVG_Line.prototype = Object.create(SVG_Element.prototype);
SVG_Line.constructor = SVG_Line;


/**
 * Draw a circle in a given canvas.
 * @param parent The parent element, can be SVG_Canvas or SVG_Group
 * @param {string} id The id of the circle
 * @param {string} className The class name of that circle.
 * @param {int} cx  The center x position
 * @param {int} cy  The center y position
 * @param {int} r   The radius
 * @returns {*}
 * @extends SVG_Element
 * @constructor
 */
var SVG_Circle = function( parent, id, className, cx, cy, r){

    SVG_Element.call(this,parent,id,className, "circle");

    // The svg element.
    this.element.setAttribute("cx", cx);
    this.element.setAttribute("cy", cy);
    this.element.setAttribute("r", r);

    return this;
}

SVG_Circle.prototype = Object.create(SVG_Element.prototype);
SVG_Circle.constructor = SVG_Circle;

/**
 * Draw an ellipse in a given canvas.
 * @param parent The parent element, can be SVG_Canvas or SVG_Group
 * @param {string} id The id of the circle
 * @param {string} className The class name of that circle.
 * @param {int} cx  The center x position
 * @param {int} cy  The center y position
 * @param {int} rx   The radius in x
 * @param {int} ry   The radius in y
 * @extends SVG_Element
 * @constructor
 */
var SVG_Ellipse = function(parent, id, className, cx, cy, rx, ry){

    SVG_Element.call(this, parent, id,className, "ellipse");

    // The svg element.
    this.element.setAttribute("cx", cx);
    this.element.setAttribute("cy", cy);
    this.element.setAttribute("rx", rx);
    this.element.setAttribute("ry", ry);

    return this;
}

SVG_Ellipse.prototype = Object.create(SVG_Element.prototype);
SVG_Ellipse.constructor = SVG_Ellipse;

/**
 * Draw a rectangle at a given position width.
 * @param parent The parent element, can be SVG_Canvas or SVG_Group
 * @param {string} id The id of the rectangle
 * @param {string} className the name of the class
 * @param {int} x The position at x from the right
 * @param {int} y The position from the top
 * @param {int} w The width
 * @param {int} h The height
 * @param {int} rx The radius of corner width
 * @param {int} ry The radius of corner
 * @extends SVG_Element
 * @constructor
 */
var SVG_Rectangle = function( parent, id, className, x, y, w, h, rx, ry){

    SVG_Element.call(this, parent, id, className, "rect");

    // The svg element.
    this.element.setAttribute("x", x);
    this.element.setAttribute("y", y);
    this.element.setAttribute("width", w);
    this.element.setAttribute("height", h);

    // Set the radius in x and y
    if(rx != undefined){
        this.element.setAttribute("rx", rx);
    }

    if(ry != undefined){
        this.element.setAttribute("ry", ry);
    }

    return this;
}

SVG_Rectangle.prototype = Object.create(SVG_Element.prototype);
SVG_Rectangle.constructor = SVG_Rectangle;

/**
 * Draw a polyline from the serie of points.
 * @param parent The parent element, can be SVG_Canvas or SVG_Group
 * @param {string} id The id of the polyline
 * @param {string} className  The name of the class.
 * @param points The list of points.
 * @extends SVG_Element
 * @constructor
 */
var SVG_Polyline = function( parent, id, className, points){

    SVG_Element.call(this, parent, id,className, "polyline");

    // The svg element.
    var pointStr = "";
    for(var i = 0; i < points.length; ++i){
        pointStr += points[i];
        if(i % 2 == 0){
            pointStr += ",";
        }else{
            pointStr += " ";
        }
    }

    this.element.setAttribute("points", pointStr);
    return this;
}

SVG_Polyline.prototype = Object.create(SVG_Element.prototype);
SVG_Polyline.constructor = SVG_Polyline;

/**
 * Similar to the polyline, but a segment is draw from the first point
 * to the last.
 * @param parent The parent element, can be SVG_Canvas or SVG_Group
 * @param {string} id The id of the polygon
 * @param {string} className The class name
 * @param points the list of coordinate, x,y,x,y.
 * @returns {*}
 * @extends SVG_Element
 * @constructor
 */
var SVG_Polygon = function( parent, id, className, points){

    SVG_Element.call(this, parent, id,className, "polygon");

    // The svg element.
    var pointStr = "";
    for(var i = 0; i < points.length; ++i){
        pointStr += points[i];
        if(i % 2 == 0){
            pointStr += ",";
        }else{
            pointStr += " ";
        }
    }

    this.element.setAttribute("points", pointStr);
    return this;
}

SVG_Polygon.prototype = Object.create(SVG_Element.prototype);
SVG_Polygon.constructor = SVG_Polygon;

/**
 * The path element regroup a serie of simple command to draw path.
    <ul>
        <li>M = moveto</li>
        <li>L = lineto</li>
        <li>H = horizontal lineto</li>
        <li>V = vertical lineto</li>
        <li>C = curveto</li>
        <li>S = smooth curveto</li>
        <li>Q = quadratic Bézier curve</li>
        <li>T = smooth quadratic Bézier curveto</li>
        <li>A = elliptical Arc</li>
        <li>Z = closepath</li>
    </ul>
 * @param parent The parent element, can be SVG_Canvas or SVG_Group
 * @param {string} id The id of the polygon
 * @param {string} className The class name
 * @param {string} subcommands The list of commands
 * @extends SVG_Element
 * @constructor
 */
var SVG_Path = function( parent, id, className, subcommands){

    SVG_Element.call(this, parent, id,className, "path");

    // The svg element.
    var dataStr = "";

    for(var i = 0; i < subcommands.length; ++i){
        dataStr += subcommands[i]
        if(i != subcommands.length - 1){
            dataStr += " ";
        }
    }

    this.element.setAttribute("d", dataStr);

    return this;
}

SVG_Path.prototype = Object.create(SVG_Element.prototype);
SVG_Path.constructor = SVG_Path;

////////////////////////////////////////////////////////////////////////////
//  Image and Gradient
////////////////////////////////////////////////////////////////////////////

/**
 * That class permit to display image inside a SVG canvas.
 * @param parent The parent element, can be SVG_Canvas or SVG_Group
 * @param {string} id The id of the polygon
 * @param {string} className The class name
 * @param {string} src The path of the image file
 * @param {int} x the x coordinate
 * @param {int} y the y coordinate
 * @param {int} w the width
 * @param {int} h the height
 * @extends SVG_Element
 * @constructor
 */
var SVG_Image = function(parent, id, className, src, x, y, w, h){

    SVG_Element.call(this,parent,id,className, "image");

    // The svg element.
    this.element.setAttributeNS("http://www.w3.org/1999/xlink", "href",  src);
    this.element.setAttribute("x",x);
    this.element.setAttribute("y",y);
    this.element.setAttribute("width",w);
    this.element.setAttribute("height",h);

    return this;
}

SVG_Image.prototype = Object.create(SVG_Element.prototype);
SVG_Image.constructor = SVG_Image;

/**
 * A gradient is a smooth transition from one color to another. In addition, several color transitions can be applied to the same element.
 * @param id The gradient identifier
 * @param className The className of the gradient.
 * @param type The type, can be Linear or Radial.
 * @extends SVG_Element
 * @constructor
 */
var SVG_Gradient = function(id, className, type){

    SVG_Element.call(this, id,className, type + "Gradient");
    this.stops = {};

    return this;
}

SVG_Gradient.prototype = Object.create(SVG_Element.prototype);
SVG_Gradient.constructor = SVG_Gradient;

/**
 * Apply a rotation on the radiant from it center point.
 * @param {float} The angle of rotation.
 */
SVG_Gradient.prototype.rotate = function(angle){
   var val =  "rotate(" + angle +",.5,.5)"
   this.element.setAttribute("gradientTransform",val);
}

/**
 * Set the stop color.
 * @param {int} The offset
 * @param {string} The color
 * @param {float} The opacity
 */
SVG_Gradient.prototype.appendStop = function(offset, color, opacity){
    var id = this.id + "_stop_" + offset;
    var className = this.className + "_stop";

    // Create a new element.
    var stop = new SVG_Element(id, className, "stop");

    stop.element.setAttribute("offset", offset);
    stop.element.setAttribute("stop-color", color);
    if(opacity != undefined){
        stop.element.setAttribute("stop-opacity", opacity);
    }

    this.appendElement(stop);
}


////////////////////////////////////////////////////////////////////////////
//  Grouping and referencing 
////////////////////////////////////////////////////////////////////////////

/**
 * The use element takes nodes from within the SVG document, and duplicates them somewhere else.
 * @param parent The parent element, can be SVG_Canvas or SVG_Group
 * @param {string} id The id of the polygon
 * @param {string} className The class name
 * @param {string} ref The element reference to be clone
 * @param {int} w The width
 * @param {int} h The height
 * @param {int} x The x coordinate
 * @param {int} y The y coordinate
 * @extends SVG_Element
 * @constructor 
 */
var SVG_Use = function(parent, id, className, ref, w, h, x, y){

    SVG_Element.call(this,parent,id,className, "use");

    // The svg element.
    this.element.setAttributeNS("http://www.w3.org/1999/xlink", "href","#" + ref);
    if(x != undefined){
        this.element.setAttribute("x",x);
    }

    if(y != undefined){
        this.element.setAttribute("y",y);
    }

    if(w != undefined){
        this.element.setAttribute("width",w);
    }

    if(h != undefined){
        this.element.setAttribute("height",h);
    }

    return this;
}

SVG_Use.prototype = Object.create(SVG_Element.prototype);
SVG_Use.constructor = SVG_Use;

/**
 * Group is a container element. It can be use to manipulate a group 
 * of element with a single objet.
 * @param parent The parent element, can be SVG_Canvas or SVG_Group
 * @param {string} id The id of the polygon
 * @param {string} className The class name
 * @param {int} cx The x coordinate
 * @param {int} cy The y coordinate
 * @param {int} rx The x radius
 * @param {int} ry The y radius
 * @extends SVG_Element
 * @constructor 
 */
var SVG_Group = function(parent,id, className, cx, cy, rx, ry){

    SVG_Element.call(this,parent,id,className, "g");
    this.element.setAttribute("cx",cx);
    this.element.setAttribute("cy",cy);
    if(rx != undefined){
        this.element.setAttribute("rx",rx);
    }
    if(ry != undefined){
        this.element.setAttribute("ry",ry);
    }
    return this;
}

SVG_Group.prototype = Object.create(SVG_Element.prototype);
SVG_Group.constructor = SVG_Group;

////////////////////////////////////////////////////////////////////////////
//  Text
////////////////////////////////////////////////////////////////////////////

/**
 * The Text element is used to define a text.
 * @param parent The parent element, can be SVG_Canvas or SVG_Group
 * @param {string} id The id of the polygon
 * @param {string} className The class name
 * @param {string} text The text to display
 * @param {int} x The x coordinate
 * @param {int} y The y coordinate
 * @param {int} dx The x distance
 * @param {int} dy The y distance
 * @extends SVG_Element
 * @constructor 
 */
var SVG_Text = function(parent, id, className, text, x, y,dx, dy){

    SVG_Element.call(this,parent,id,className, "text");

    // The svg element.
    this.element.textContent = text;
    this.element.setAttribute("x",x);
    this.element.setAttribute("y",y);

    if(dx != undefined){
        this.element.setAttribute("dx",dx);
    }

    if(dy != undefined){
        this.element.setAttribute("dy",dy);
    }
    return this;
}

SVG_Text.prototype = Object.create(SVG_Element.prototype);
SVG_Text.constructor = SVG_Text;

/**
 * The Text element can be arranged in any number of sub-groups with the tspan element. Each tspan element can contain different formatting and position.
 * @param parent The parent element, can be SVG_Canvas or SVG_Group
 * @param {string} id The id of the polygon
 * @param {string} className The class name
 * @param {string} text The text to display
 * @param {int} dx The x distance
 * @param {int} dy The y distance
 * @param {int} x The x coordinate
 * @param {int} y The y coordinate
 * @extends SVG_Element
 * @constructor 
 */
SVG_Text.prototype.appendSpan = function(parent, id, className, text, dx, dy,  x, y){
    var span = new SVG_Element(parent, id, className, "tspan");
    span.element.textContent = text;

    if(dx != undefined){
        span.element.setAttribute("dx",dx);
    }

    if(dy != undefined){
        span.element.setAttribute("dy",dy);
    }

    if(x != undefined){
        span.element.setAttribute("x",x);
    }

    if(y != undefined){
        span.element.setAttribute("y",y);
    }

    this.appendElement(span);
}

/**
 * Set the text.
 * @param {string} value The text value itself.
 */
SVG_Text.prototype.setText = function(value){
    this.element.textContent = value;
}

/**
 * The text can be follow a given path.
 * @param {string} pathId The path id
 * @param {text} text The text to display
 */
SVG_Text.prototype.setTextPath = function(pathId, text){
    var id = this.id + "_text_path_" + pathId;
    var className = this.className + "_text_path";
    var textPath = new SVG_Element(id, className, "textPath");
    textPath.element.setAttributeNS("http://www.w3.org/1999/xlink", "href","#" + pathId);
    textPath.element.textContent = text;

    this.appendElement(textPath);
}

/**
 * The textual content for a text can be either character data directly embedded within the <text> element or the character data content of a 
 * referenced element, where the referencing is specified with a tref element.
 * @param parent The parent element, can be SVG_Canvas or SVG_Group
 * @param {string} id The id of the polygon
 * @param {string} className The class name
 * @param {string} text The text to display
 * @param {string} ref The text element reference
 * @extends SVG_Element
 * @constructor 
 */
var SVG_tRef = function(parent, id, className, ref){

    SVG_Element.call(this,parent,id,className, "tref");

    // The svg element.
    this.element.setAttributeNS("http://www.w3.org/1999/xlink", "href","#" + ref);

    return this;
}

SVG_tRef.prototype = Object.create(SVG_Element.prototype);
SVG_tRef.constructor = SVG_tRef;

////////////////////////////////////////////////////////////////////////////
//  Text area
////////////////////////////////////////////////////////////////////////////

/**
 * ForeignObject is use to embedded html element inside a SVG element.
 * @param parent The parent element, can be SVG_Canvas or SVG_Group
 * @param {string} id The id of the polygon
 * @param {string} className The class name
 * @param {int} x The x coordinate
 * @param {int} y The y coordinate
 * @param {int} w The width
 * @param {int} h The height
 * @extends SVG_Element
 * @constructor 
 */
var SVG_ForeignObject = function(parent, id, className, x, y, w, h){

    SVG_Element.call(this,parent,id,className, "foreignObject");

    // Here if the browser is IE 11 the foreignObject is not supported,
    // tank's to microsoft.
    if(getNavigatorName()=="IE"){
        var canvas = this.getCanvas()
        var parent = canvas.parentElement
        // I will set the parent as relative.
        //parent.element.style.position = "relative"
        this.element = parent.appendElement({"tag":"div", "style":"position:absolute;"}).down().element
        this.element.style.left = x + "px"
        this.element.style.top = y + "px"
        this.element.style.width = w + "px"
        this.element.style.height = h + "px"
    }else{
        this.element.setAttribute("x",x);
        this.element.setAttribute("y",y);
        this.element.setAttribute("width",w);
        this.element.setAttribute("height",h);
    }

    return this;
}

SVG_ForeignObject.prototype = Object.create(SVG_Element.prototype);
SVG_ForeignObject.constructor = SVG_ForeignObject;

////////////////////////////////////////////////////////////////////////////
//  Pattern
////////////////////////////////////////////////////////////////////////////

/**
 * SVG fill patterns are used to fill a shape with a pattern made up from images. This pattern can be made up from SVG images (shapes) or from bitmap images.
 * @param parent The parent element, can be SVG_Canvas or SVG_Group
 * @param {string} id The id of the polygon
 * @param {string} className The class name
 * @param {int} w The width
 * @param {int} h The height
 * @param {int} unit The pattern units
 * @extends SVG_Element
 * @constructor 
 */
var SVG_Pattern = function(parent, id, className, w, h, unit){

    SVG_Element.call(this,parent,id,className, "pattern");

    // The svg element.
    this.element.setAttribute("width", w);
    this.element.setAttribute("height", h);

    // Value are objectBoundingBox or userSpaceOnUse
    if(unit == undefined){
        this.element.setAttribute("patternUnits", "userSpaceOnUse");
    }else{
        this.element.setAttribute("patternUnits", unit);
    }

    return this;
}

SVG_Pattern.prototype = Object.create(SVG_Element.prototype);
SVG_Pattern.constructor = SVG_Pattern;

SVG_Pattern.prototype.applyTo = function(to){
    to.setSvgAttribute("fill", "url(#"+ this.id +")");
}

////////////////////////////////////////////////////////////////////////////
//  Clipping and masking
////////////////////////////////////////////////////////////////////////////

/**
 * The clipping path restricts the region to which paint can be applied.
 * @param parent The parent element, can be SVG_Canvas or SVG_Group
 * @param {string} id The id of the polygon
 * @param {string} className The class name
 * @extends SVG_Element
 * @constructor 
 */
var SVG_ClipPath = function(parent, id, className){
    SVG_Element.call(this,parent,id,className, "clipPath");
    return this;
}

SVG_ClipPath.prototype = Object.create(SVG_Element.prototype);
SVG_ClipPath.constructor = SVG_ClipPath;

/**
 * Apply the clip region to a given element
 * @param {string} to The reference of the element to apply this clip.
 */
SVG_ClipPath.prototype.applyTo = function(to){
    to.setSvgAttribute("clip-path", "url(#"+ this.id +")");
}

/**
 * The SVG masking feature makes it possible to apply a mask to an SVG shape. 
 * @param {string} id The id of the polygon
 * @param {string} className The class name
 * @extends SVG_Element
 * @constructor 
 */
var SVG_Mask = function(id, className){
    SVG_Element.call(this, id,className, "mask");
    return this;
}

SVG_Mask.prototype = Object.create(SVG_Element.prototype);
SVG_Mask.constructor = SVG_Mask;

/**
 * Apply the mask to a given element.
 * @param {string} to The reference of the element to apply this mask.
 */
SVG_Mask.prototype.applyTo = function(to){
    to.setSvgAttribute("mask", "url(#"+ this.id +")");
}

////////////////////////////////////////////////////////////////////////////
//  Animation
////////////////////////////////////////////////////////////////////////////
