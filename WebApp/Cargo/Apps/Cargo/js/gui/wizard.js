
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
 */

/**
 * @fileOverview Generic wizard.
 * @author Dave Courtois
 * @version 1.0
 */

/**
 * The wizard page.
 */
var WizardPage = function (id, title, instruction, content, width, height) {
    // The id of the page.
    this.id = id
    // The current index of the page.
    this.index = 0
    // The title display in the progression pane.
    this.title = title
    // The instruction what the user need to enter as info...
    this.instruction = instruction
    // The page width
    this.width = width
    // the page height
    this.height = height

    this.panel = new Element(null, { "tag": "div", "class": "wizardPage" })

    return this
}

/**
 * The base wizard class.
 */
var Wizard = function (id, title) {
    if (title == undefined) {
        return
    }

    // The maps of page.
    this.pages = {}
    this.pagesByIndex = {}

    // The active page.
    this.currentPage = null

    // Set the id, or make it random...
    this.id = id
    if (this.id == undefined) {
        this.id = randomUUID()
    }

    // I will create a modal dialog to display the wizard.
    this.dialog = new Dialog(id, null, true, title)
    this.dialog.div.element.style.maxWidth = "auto";
    this.dialog.div.element.style.maxHeight = "auto";

    // The content...
    this.content = this.dialog.content.appendElement({ "tag": "div", "style": "display: table;" }).down()

    // the step panel width
    this.stepPanelWidth = 200
    // the instruction panel
    this.instructionPanelHeight = 50

    this.stepPanel = this.content.appendElement({ "tag": "div", "class": "stepPanel", "style": "width:" + this.stepPanelWidth + "px;" }).down()
    this.pagePanelContent = this.content.appendElement({ "tag": "div", "class": "pagePanelContent" }).down()

    // That grid will be use to slide inside the page layout grid...
    this.pagesDiv = this.pagePanelContent.appendElement({ "tag": "div", "class": "wizardPages" }).down()

    // Navigation buttons.
    this.okButton = this.dialog.ok
    this.okButton.element.style.display = "none"
    this.cancelButton = this.dialog.cancel
    this.nextButton = this.dialog.appendButton("next", "", "fa fa-arrow-right")
    this.previousButton = this.dialog.appendButton("previous", "", "fa fa-arrow-left")

    // Translate the page
    this.nextButton.element.onclick = function (wizard) {
        return function () {
            if (wizard.currentPage.index == Object.keys(wizard.pages).length - 1) {
                return
            }
            if (wizard.currentPage != null) {
                wizard.setPage(wizard.currentPage.index + 1)
            }
            // Calculate the offset...
            var offset = wizard.pagesDiv.element.offsetLeft + (-1 * wizard.currentPage.width)

            // Set the keyframe inner text.
            var keyframe = "100% { left:" + offset + "px;}"
            wizard.pagesDiv.animate(keyframe, "wizard_next_page_animation", 1,
                function (pageDiv, offset) {
                    return function () {
                        pageDiv.element.style.left = offset + "px"
                    }
                } (wizard.pagesDiv, offset))
        }
    } (this)

    this.previousButton.element.onclick = function (wizard) {
        return function () {
            if (wizard.currentPage.index == 0) {
                return
            }
            if (wizard.currentPage != null) {
                wizard.setPage(wizard.currentPage.index - 1)
            }
            var offset = wizard.pagesDiv.element.offsetLeft + wizard.currentPage.width
            var keyframe = "100% { left:" + offset + "px;}"
            wizard.pagesDiv.animate(keyframe, "wizard_next_page_animation", 1,
                function (pageDiv, offset) {
                    return function () {
                        pageDiv.element.style.left = offset + "px"
                    }
                } (wizard.pagesDiv, offset))
        }
    } (this)

    // Now the action...
    return this
}

Wizard.prototype.appendPage = function (page) {

    // Here I will set the content size in respect of the page size...
    if (this.content.element.offsetWidth < page.width + this.stepPanelWidth) {
        this.content.element.style.width = page.width + this.stepPanelWidth + "px"
        this.dialog.div.element.style.maxWidth = page.width + this.stepPanelWidth + "px"
        // put at center.
        this.dialog.setCentered()

    }

    if (this.content.element.offsetHeight < page.height) {
        this.stepPanel.element.style.height = page.height + this.instructionPanelHeight + "px"
        this.dialog.setCentered()
    }

    // Set the index
    page.index = Object.keys(this.pages).length

    // Keep the page reference...
    this.pages[page.id] = page
    this.pagesByIndex[page.index] = page
    this.nextButton.element.style.display = ""
    this.stepPanel.appendElement({ "tag": "div", "class": "wizardStep", "id": "step_panel_" + page.index, "style": "color: lightgrey;", "innerHtml": page.title }).down()

    // Set the page component here.
    this.pagesDiv.appendElement({ "tag": "div", "style": "display: table-cell;" }).down()
        .appendElement({ "tag": "div", "class": "pageInstruction", "style": "width: " + page.width + "px; height:" + this.instructionPanelHeight + "px;", "innerHtml": page.instruction })
        .appendElement({ "tag": "div", "style": "width: " + page.width + "px; height:" + page.height + "px; padding: 10px;" }).down()
        .appendElement(page.panel)

    if (page.index == 0) {
        this.setPage(page.index)
    }
}

Wizard.prototype.setPage = function (index) {
    if (index == 0) {
        this.previousButton.element.style.display = "none"
    } else {
        this.previousButton.element.style.display = ""
    }

    if (index == Object.keys(this.pages).length - 1) {
        this.nextButton.element.style.display = "none"
    } else {
        this.nextButton.element.style.display = ""
    }

    // Reset the previous current page...
    if (this.currentPage != undefined) {
        var previousStep = this.stepPanel.getChildById("step_panel_" + this.currentPage.index)
        previousStep.element.style.color = ""
        previousStep.element.style.backgroundColor = ""
    }

    // Set the current page.
    this.currentPage = this.pagesByIndex[index]
    var step = this.stepPanel.getChildById("step_panel_" + this.currentPage.index)
    step.element.style.color = "white"
    step.element.style.backgroundColor = "#657383"

    // here I will set the step panel action.
    step.element.onclick = function (wizard, page) {
        return function () {
            // here I will set the page offset...
            var offset = -1 * page.index * page.width
            var keyframe = "100% { left:" + offset + "px;}"
            wizard.setPage(page.index)
            wizard.pagesDiv.animate(keyframe, "wizard_next_page_animation", 1,
                function (pageDiv, offset) {
                    return function () {
                        pageDiv.element.style.left = offset + "px"
                    }
                } (wizard.pagesDiv, offset))
        }
    } (this, this.currentPage)
}