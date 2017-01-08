
/**
 * Set splitter action...
 */
function initSplitter(splitter, area) {

	var moveSplitter = function (splitter, area) {
		return function (e) {
			var isVertical = splitter.element.className.indexOf("vertical") > -1
			if (isVertical) {
				var isAtLeft = splitter.element.offsetLeft >= area.element.offsetLeft
				var offset = e.clientX - splitter.element.offsetLeft
				var w = area.element.offsetWidth
				if (isAtLeft) {
					w = area.element.offsetWidth + offset
				} else {
					w = area.element.offsetWidth - offset
				}
				area.element.style.width = w + "px"
			} else {
				var isAtTop = splitter.element.offsetTop >= area.element.offseTop
				var offset = e.clientY - splitter.element.offsetTop
				var h = area.element.offsetHeight
				if (isAtTop) {
					h = area.element.offsetHeight + offset
				} else {
					h = area.element.offsetHeight - offset
				}
				area.element.style.height = h + "px"
			}
		}
	} (splitter, area)

	function mouseUp() {
		window.removeEventListener('mousemove', moveSplitter);
		window.removeEventListener('mouseup', mouseUp);
	}

    splitter.element.onmousedown = function () {
		window.addEventListener('mousemove', moveSplitter);
		window.addEventListener('mouseup', mouseUp);
	}

}