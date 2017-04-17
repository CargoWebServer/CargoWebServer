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
 * @fileOverview Language manipulation.
 * @author Dave Courtois, Eric Kavalec
 * @version 1.0
 */

/**
 * The language manager is used to manage text translations.
 * @constructor
 */
var LanguageManager = function () {
	/**
	 * @property The language map that contains the language informations.
	 */
	this.languageInfo = {}

	/**
	 * @property The references to the text elements containing text managed by this class.
	 */
	this.textElements = {}

	/**
	 * @property The current language.
	 */
	this.language = null

	var language = navigator.language || navigator.userLanguage;

	if (language != undefined) {
		this.language = language.split("-")[0]
	}
	this.elements = {}

	return this
}

/**
 * Merge a map containning the language information with the member language info of this object.
 * @param {*} languageInfo The map containning the information to append.
 */
LanguageManager.prototype.appendLanguageInfo = function (languageInfo) {
	for (var languageId in languageInfo) {
		if (this.languageInfo[languageId] == undefined) {
			this.languageInfo[languageId] = languageInfo[languageId]
		} else {
			for (var textId in languageInfo[languageId]) {
				this.languageInfo[languageId][textId] = languageInfo[languageId][textId]
			}
		}
	}
}

/**
 * Associate an element (span, div, textarea.) with a text information with a given id.
 */
LanguageManager.prototype.registerElementText = function (element, textId) {
	if (this.textElements[textId] == undefined) {
		this.textElements[textId] = []
	}

	this.textElements[textId].push(element)

	// update the element in the array.
	this.elements[element.id] = element
}

/**
 * Set the current language and update all text elements.
 * @param {} language The language to set. {'fr', 'en', 'sp'.}
 */
LanguageManager.prototype.setLanguage = function (language) {
	this.language = language
	for (var textElementsId in this.textElements) {
		for (var i = 0; i < this.textElements[textElementsId].length; i++) {
			var element = this.textElements[textElementsId][i]
			if (text != undefined) {
				this.setElementText(element, textElementsId)
			}
		}
	}
}

/**
 * Set the text of an element with the text in the languageInfo map with a given id.
 * @param {Element} element The element that contains the text.
 * @param {string} textId The id of an element in the languageInfo map.
 */
LanguageManager.prototype.setElementText = function (element, textId) {
	if (this.languageInfo[this.language][textId] == undefined) {
		return
	}
	this.registerElementText(element, textId)
	if (this.elements[element.id].element.tagName == "SPAN") {
		this.elements[element.id].element.textContent = this.languageInfo[this.language][textId]
	} else if (this.elements[element.id].element.tagName == "INPUT"){
		if(this.elements[element.id].element.getAttribute("data-match-error") != null){
			// Set the error text here.
			this.elements[element.id].element.attributes["data-match-error"].nodeValue = this.languageInfo[this.language][textId]
		}
	} else {
		this.elements[element.id].setAttribute("innerHTML", this.languageInfo[this.language][textId])
	}
}

/**
 * Get the text of an item with a given id for the current language.
 * @param {string} textId The text item id.
 * @returns {string} The corresponding text.
 */
LanguageManager.prototype.getText = function (textId) {
	return this.languageInfo[this.language][textId]
}
