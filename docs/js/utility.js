//////////////////////////////////////////////////////////////////////////
// Cago hljs html element creation function.
//////////////////////////////////////////////////////////////////////////

/**
 * Append html meta tag element at a given offset.
 */
function appendHtmlMeta(parent, meta, offset) {
    // append space...
    if (offset == undefined) {
        offset = 0;
    }

    for (var i = 0; i < offset; i++) {
        parent.appendElement({ "tag": "span", innerHtml: "&nbsp" })
    }
    parent.appendElement({ "tag": "span", "class": "hljs-meta", "innerHtml": meta })
}

/**
 * Append a comment to a parent element at a given offset from the left.
 */
function appendHtmlComment(parent, comment, offset) {
    // append space...
    parent.appendElement({ "tag": "br" })
    if (offset == undefined) {
        offset = 0;
    }

    for (var i = 0; i < offset; i++) {
        parent.appendElement({ "tag": "span", innerHtml: "&nbsp" })
    }

    parent.appendElement({ "tag": "span", "class": "hljs-comment", "innerHtml": "&lt;!-- " + comment + " --&gt;" })
}

/**
 * Create html tag
 * offset from the left.
 * tag the name of the tag.
 * attributtes array of attribute of form [{"name":"attr0","value":"val0"}, ...]
 * opentTag if false the tag is considere a closing tag.
 * value if a value is given the tag will be printed on the same line.
 */
function appendHtmlTag(parent, offset, tag, attributes, openTag, value, sameLine) {

    if (sameLine == undefined) {
        parent.appendElement({ "tag": "br" })
    }

    var lt = "&lt"
    if (openTag == false) {
        lt += ";/"
    }

    // append space...
    for (var i = 0; i < offset; i++) {
        parent.appendElement({ "tag": "span", innerHtml: "&nbsp" })
    }

    // Open brace
    var tagElement = parent.appendElement({ "tag": "span", "class": "hljs-tag", "innerHtml": lt }).down()
    // Inner tag value
    tagElement.appendElement({ "tag": "span", "class": "hljs-name", innerHtml: tag })
    // attributes
    if (attributes != undefined) {

        for (var i = 0; i < attributes.length; i++) {

            tagElement.appendElement({ "tag": "span", "innerHtml": "&nbsp" })

            var attribute = attributes[i]
            tagElement.appendElement({ "tag": "span", "class": "hljs-attr", innerHtml: attribute.name })
                .appendElement({ "tag": "span", "innerHtml": "=" })
                .appendElement({ "tag": "span", "class": "hljs-string", innerHtml: '"' + attribute.value + '"' })
        }
    }
    // Close brace
    tagElement.appendElement({ "tag": "span", "class": "hljs-tag", "innerHtml": "&gt" }).down()

    // Here I will append the value and close the element.
    if (value != undefined) {
        parent.appendElement({ "tag": "span", "innerHtml": value })
        appendHtmlTag(parent, 0, tag, [], false, undefined, "")
    }
}