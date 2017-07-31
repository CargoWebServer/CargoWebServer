
/**
 * Keep the list of used bpmn elements.
 */
var entities = {}

/**
 * Initalyse defintions.
 * @param {*} definitions The list of defintions to intialise.
 */
function initDefinitions(definitions) {

    for (var i = 0; i < definitions.M_BPMNDiagram.length; i++) {
        // Here I will use function insted of reference, so there will no circulare ref at serialysation time...
        definitions.M_BPMNDiagram[i].getParentDefinitions = function (definitions) {
            return function () {
                return definitions
            }
        } (definitions)

        for (var j = 0; j < definitions.M_BPMNDiagram[i].M_BPMNPlane.M_DiagramElement.length; j++) {
            definitions.M_BPMNDiagram[i].M_BPMNPlane.M_DiagramElement[j].getParentPlane = function (parentPlane) {
                return function () {
                    return parentPlane
                }
            } (definitions.M_BPMNDiagram[i].M_BPMNPlane)

            definitions.M_BPMNDiagram[i].M_BPMNPlane.M_DiagramElement[j].getParentDiagram = function (parentDiagram) {
                return function () {
                    return parentDiagram
                }
            } (definitions.M_BPMNDiagram[i])
        }
    }
}
