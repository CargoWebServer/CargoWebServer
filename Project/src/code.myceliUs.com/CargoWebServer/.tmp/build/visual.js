var __extends = (this && this.__extends) || (function () {
    var extendStatics = function (d, b) {
        extendStatics = Object.setPrototypeOf ||
            ({ __proto__: [] } instanceof Array && function (d, b) { d.__proto__ = b; }) ||
            function (d, b) { for (var p in b) if (b.hasOwnProperty(p)) d[p] = b[p]; };
        return extendStatics(d, b);
    }
    return function (d, b) {
        extendStatics(d, b);
        function __() { this.constructor = d; }
        d.prototype = b === null ? Object.create(b) : (__.prototype = b.prototype, new __());
    };
})();
/*
 *  Power BI Visualizations
 *
 *  Copyright (c) Microsoft Corporation
 *  All rights reserved.
 *  MIT License
 *
 *  Permission is hereby granted, free of charge, to any person obtaining a copy
 *  of this software and associated documentation files (the ""Software""), to deal
 *  in the Software without restriction, including without limitation the rights
 *  to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 *  copies of the Software, and to permit persons to whom the Software is
 *  furnished to do so, subject to the following conditions:
 *
 *  The above copyright notice and this permission notice shall be included in
 *  all copies or substantial portions of the Software.
 *
 *  THE SOFTWARE IS PROVIDED *AS IS*, WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 *  IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 *  FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 *  AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 *  LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 *  OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
 *  THE SOFTWARE.
 */
var powerbi;
(function (powerbi) {
    var extensibility;
    (function (extensibility) {
        var visual;
        (function (visual) {
            "use strict";
            var DataViewObjectsParser = powerbi.extensibility.utils.dataview.DataViewObjectsParser;
            var VisualSettings = /** @class */ (function (_super) {
                __extends(VisualSettings, _super);
                function VisualSettings() {
                    var _this = _super !== null && _super.apply(this, arguments) || this;
                    _this.dataPoint = new dataPointSettings();
                    return _this;
                }
                return VisualSettings;
            }(DataViewObjectsParser));
            visual.VisualSettings = VisualSettings;
            var dataPointSettings = /** @class */ (function () {
                function dataPointSettings() {
                    // Default color
                    this.defaultColor = "";
                    // Show all
                    this.showAllDataPoints = true;
                    // Fill
                    this.fill = "";
                    // Color saturation
                    this.fillRule = "";
                    // Text Size
                    this.fontSize = 12;
                }
                return dataPointSettings;
            }());
            visual.dataPointSettings = dataPointSettings;
        })(visual = extensibility.visual || (extensibility.visual = {}));
    })(extensibility = powerbi.extensibility || (powerbi.extensibility = {}));
})(powerbi || (powerbi = {}));
/*
 *  Power BI Visual CLI
 *
 *  Copyright (c) Microsoft Corporation
 *  All rights reserved.
 *  MIT License
 *
 *  Permission is hereby granted, free of charge, to any person obtaining a copy
 *  of this software and associated documentation files (the ""Software""), to deal
 *  in the Software without restriction, including without limitation the rights
 *  to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 *  copies of the Software, and to permit persons to whom the Software is
 *  furnished to do so, subject to the following conditions:
 *
 *  The above copyright notice and this permission notice shall be included in
 *  all copies or substantial portions of the Software.
 *
 *  THE SOFTWARE IS PROVIDED *AS IS*, WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 *  IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 *  FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 *  AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 *  LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 *  OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
 *  THE SOFTWARE.
 */
var powerbi;
(function (powerbi) {
    var extensibility;
    (function (extensibility) {
        var visual;
        (function (visual) {
            "use strict";
            var Visual = /** @class */ (function () {
                function Visual(options) {
                    console.log('Visual constructor', options);
                    this.target = options.element;
                    this.updateCount = 0;
                    if (typeof document !== "undefined") {
                        var new_p = document.createElement("p");
                        new_p.appendChild(document.createTextNode("Update count:"));
                        var new_em = document.createElement("em");
                        this.textNode = document.createTextNode(this.updateCount.toString());
                        new_em.appendChild(this.textNode);
                        new_p.appendChild(new_em);
                        this.target.appendChild(new_p);
                    }
                }
                Visual.prototype.update = function (options) {
                    this.settings = Visual.parseSettings(options && options.dataViews && options.dataViews[0]);
                    console.log('Visual update', options);
                    if (typeof this.textNode !== "undefined") {
                        this.textNode.textContent = (this.updateCount++).toString();
                    }
                };
                Visual.parseSettings = function (dataView) {
                    return visual.VisualSettings.parse(dataView);
                };
                /**
                 * This function gets called for each of the objects defined in the capabilities files and allows you to select which of the
                 * objects and properties you want to expose to the users in the property pane.
                 *
                 */
                Visual.prototype.enumerateObjectInstances = function (options) {
                    return visual.VisualSettings.enumerateObjectInstances(this.settings || visual.VisualSettings.getDefault(), options);
                };
                return Visual;
            }());
            visual.Visual = Visual;
        })(visual = extensibility.visual || (extensibility.visual = {}));
    })(extensibility = powerbi.extensibility || (powerbi.extensibility = {}));
})(powerbi || (powerbi = {}));
//# sourceMappingURL=visual.js.map