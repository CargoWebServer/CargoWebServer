var languageInfo = {
	"en": {
		"SessionState_Online": "Online",
		"SessionState_Away": "Away",
		"SessionState_Offline": "Offline",
		"Session_Close": "Close"
	},
	"fr": {
		"SessionState_Online": "En ligne",
		"SessionState_Away": "Occupé",
		"SessionState_Offline": "Hors ligne",
		"Session_Close": "Déconnexion"
	}
}

// Set the text...
server.languageManager.appendLanguageInfo(languageInfo)

/*
 * The session panel contain information about the current session
 * It also contain notification panel.
 **/
var SessionPanel = function (parent, sessionInfo) {
	/* The current session info **/
	this.id = randomUUID()

	// The current session information.
	this.sessionInfo = sessionInfo

	// The current user 
	this.userInfo = null

	// The current account
	this.accountInfo = null

	/* The div content **/
	this.div = parent.appendElement({ "tag": "div", "class": "session_panel" }).down()

	/* Now the button **/
	this.nameDisplay = this.div.appendElement({ "tag": "span" }).down()

	this.notificationBtn = this.div.appendElement({ "tag": "div", "class": "innactive" }).down()
	this.notificationBtn.appendElement({ "tag": "i", "class": "fa fa-bell" })

	this.userSessionBtn = this.div.appendElement({ "tag": "div", "class": "active" }).down()
	this.userSessionBtn.appendElement({ "tag": "i", "class": "fa fa-user" })

	// Contain information about the current session...
	this.sessionDisplayPanel = new Element(document.getElementsByTagName("body")[0], { "tag": "div", "class": "session_display_panel" })
	this.currentSessionInfoPanel = null

	// This contain the settings...
	this.settingsDisplayPanel = new Element(document.getElementsByTagName("body")[0], { "tag": "div", "class": "session_display_panel" })

	// Contain notifications...
	this.notificationsDisplayPanel = new Element(document.getElementsByTagName("body")[0], { "tag": "div", "class": "session_display_panel" })

	// Contain information about other sessions...
	this.otherSessionPanel = null
	this.otherSessionPanels = {}

	// the logout button.
	this.logoutBtn = null

	// Here I will display the current session information...
	// I will get the account pointer...
	this.setAccountInfo(sessionInfo.M_accountPtr)

	// Set the user information.
	if (sessionInfo.M_accountPtr.M_userRef != undefined) {
		// Set the user information if founded.
		this.setUserInfo(sessionInfo.M_accountPtr.M_userRef)
	} else {
		// Set the account name if there no user associated.
		this.nameDisplay.element.innerHTML = sessionInfo.M_accountPtr.M_name
	}

	/////////////////////////////////////////////////////////////////////////////////////////////
	// 	Set the action here
	/////////////////////////////////////////////////////////////////////////////////////////////
	this.userSessionBtn.element.addEventListener("click", function (sessionDisplayPanel, notificationsDisplayPanel, settingsDisplayPanel, userSessionBtn) {
		return function () {
			if (sessionDisplayPanel.element.style.display != "block") {
				sessionDisplayPanel.element.style.display = "block"
				var viewportOffset = userSessionBtn.element.parentNode.getBoundingClientRect();
				var top = viewportOffset.bottom
				var left = viewportOffset.right - sessionDisplayPanel.element.offsetWidth
				sessionDisplayPanel.element.style.top = top + "px"
				sessionDisplayPanel.element.style.left = left + "px"

			} else {
				sessionDisplayPanel.element.style.display = "none"
			}

			notificationsDisplayPanel.element.style.display = "none"
			settingsDisplayPanel.element.style.display = "none"
		}
	} (this.sessionDisplayPanel, this.notificationsDisplayPanel, this.settingsDisplayPanel, this.userSessionBtn))


	this.notificationBtn.element.addEventListener("click", function (sessionDisplayPanel, notificationsDisplayPanel, settingsDisplayPanel, userSessionBtn) {
		return function () {
			if (notificationsDisplayPanel.element.style.display != "block") {
				notificationsDisplayPanel.element.style.display = "block"
				var viewportOffset = userSessionBtn.element.parentNode.getBoundingClientRect();
				var top = viewportOffset.bottom
				var left = viewportOffset.right - notificationsDisplayPanel.element.offsetWidth
				notificationsDisplayPanel.element.style.top = top + "px"
				notificationsDisplayPanel.element.style.left = left + "px"
			} else {
				notificationsDisplayPanel.element.style.display = "none"
			}

			sessionDisplayPanel.element.style.display = "none"
			settingsDisplayPanel.element.style.display = "none"
		}
	} (this.sessionDisplayPanel, this.notificationsDisplayPanel, this.settingsDisplayPanel, this.userSessionBtn))


	this.nameDisplay.element.addEventListener("click", function (sessionDisplayPanel, notificationsDisplayPanel, settingsDisplayPanel, userSessionBtn) {
		return function () {
			if (settingsDisplayPanel.element.style.display != "block") {
				settingsDisplayPanel.element.style.display = "block"
				var viewportOffset = userSessionBtn.element.parentNode.getBoundingClientRect();
				var top = viewportOffset.bottom
				var left = viewportOffset.right - settingsDisplayPanel.element.offsetWidth
				settingsDisplayPanel.element.style.top = top + "px"
				settingsDisplayPanel.element.style.left = left + "px"
			} else {
				settingsDisplayPanel.element.style.display = "none"
			}

			notificationsDisplayPanel.element.style.display = "none"
			sessionDisplayPanel.element.style.display = "none"
		}
	} (this.sessionDisplayPanel, this.notificationsDisplayPanel, this.settingsDisplayPanel, this.userSessionBtn))

	return this
}

/*
 * Set the account that own the active session
 */
SessionPanel.prototype.setAccountInfo = function (accountInfo) {
	// The current account
	this.accountInfo = accountInfo

	// So here I will create the information about the current session.
	this.currentSessionInfoPanel = new SessionInfoPanel(this.sessionDisplayPanel, this.sessionInfo, this.accountInfo, true, false)

	// The othe session info panels...
	this.otherSessionPanel = this.sessionDisplayPanel.appendElement({ "tag": "div", "style": "display: none; heigth:auto; border-bottom: 1px solid #bbb; border-top: 1px solid #bbb" }).down()

	// Now the other session from the same user...
	if (this.accountInfo.M_sessions != undefined) {
		for (var i = 0; i < this.accountInfo.M_sessions.length; i++) {
			if (this.accountInfo.M_sessions[i].UUID != this.sessionInfo.UUID) {
				var session = this.accountInfo.M_sessions[i]
				var sessionInfoPanel = new SessionInfoPanel(this.otherSessionPanel, session, accountInfo, false, true)
				// I will keep the reference to the panel...
				this.otherSessionPanels[session.UUID] = sessionInfoPanel
				this.otherSessionPanel.element.style.display = "inline-block"
			}
		}
	}


	///////////////////////////////////////////////////////////////////////////////////////////////////////////
	// The event listeners...
	///////////////////////////////////////////////////////////////////////////////////////////////////////////
	var stateEventListener = function (evt, sessionPanel) {

		var session = evt.dataMap.sessionInfo
		var sessions = evt.dataMap.sessionsInfo
		var sessionInfoPanel = null
		if (session.M_accountPtr == sessionPanel.accountInfo.UUID) {
			// Update all sessions...
			sessionPanel.accountInfo.M_sessions = sessions
			if (session.UUID == sessionPanel.sessionInfo.UUID) {
				// Set the session info with the new value...
				sessionPanel.sessionInfo = session
				sessionInfoPanel = sessionPanel.currentSessionInfoPanel
				if (sessionPanel.sessionInfo.M_sessionState == 1) {
					sessionPanel.userSessionBtn.element.firstChild.className = "fa fa-user active"
				} else if (sessionPanel.sessionInfo.M_sessionState == 2) {
					sessionPanel.userSessionBtn.element.firstChild.className = "fa fa-user away"
				} else if (sessionPanel.sessionInfo.M_sessionState == 3) {
					sessionPanel.userSessionBtn.element.firstChild.className = "fa fa-user innactive"
				}
			} else {
				sessionInfoPanel = sessionPanel.otherSessionPanels[session.UUID]
			}

			sessionInfoPanel.sessionInfo = session
			sessionInfoPanel.accountInfo.M_sessions = sessions

			if (sessionInfoPanel.sessionInfo.M_sessionState == 1) {
				sessionInfoPanel.sessionState.element.firstChild.className = "fa fa-desktop active"
			} else if (sessionInfoPanel.sessionInfo.M_sessionState == 2) {
				sessionInfoPanel.sessionState.element.firstChild.className = "fa fa-desktop away"
			} else if (sessionInfoPanel.sessionInfo.M_sessionState == 3) {
				sessionInfoPanel.sessionState.element.firstChild.className = "fa fa-desktop innactive"
			}
		}
		// Todo append message in the growl panel...
	}

	var logoutEventListener = function (evt, sessionPanel) {
		// So here i will remove the session from the other sessions...
		var session = evt.dataMap["sessionInfo"]
		var toDelete = sessionPanel.otherSessionPanels[session.UUID]

		if (toDelete != undefined) {
			toDelete.parent.removeElement(toDelete.sessionPanel)
			delete sessionPanel.otherSessionPanels[session.UUID]
			if (Object.keys(sessionPanel.otherSessionPanels).length == 0) {
				sessionPanel.otherSessionPanel.element.style.display = "none"
			}
		}
	}

	var loginEventListener = function (logoutEventListener, stateEventListener) {
		return function (evt, sessionPanel) {
			// if the session dosent exist I will append it...
			var session = evt.dataMap["sessionInfo"]
			var sessions = evt.dataMap["sessionsInfo"]

			if (sessionPanel.accountInfo.UUID == session.M_accountPtr) {
				// Update the sessions
				sessionPanel.accountInfo.M_sessions = sessions

				if (sessionPanel.otherSessionPanels[session.UUID] == null) {
					// The panel dosent exist i will create it...
					var sessionInfoPanel = new SessionInfoPanel(sessionPanel.otherSessionPanel, session, sessionPanel.accountInfo, false, true)
					// Here I wil append the other session...
					sessionPanel.otherSessionPanels[session.UUID] = sessionInfoPanel
					sessionPanel.otherSessionPanel.element.style.display = "inline-block"
				} else {
					console.log("session already exist!")
				}
			}

		}
	} (logoutEventListener, stateEventListener)


	// The login event...
	server.sessionManager.attach(this, LoginEvent, loginEventListener)
	server.sessionManager.attach(this, LogoutEvent, logoutEventListener)
	server.sessionManager.attach(this, StateChangeEvent, stateEventListener)

	// Now the logout button...
	this.logoutBtn = this.sessionDisplayPanel.appendElement({ "tag": "div", "id": "logout_btn" }).down()
	server.languageManager.setElementText(this.logoutBtn, "Session_Close")

	// Now the click event...
	this.logoutBtn.element.onclick = function (sessionInfo) {
		return function () {
			//The event will be catch by the listener
			localStorage.removeItem('_remember_me_')
			server.sessionManager.logout(sessionInfo.M_id,
				function () {

				},
				function () {

				}, null)
		}
	} (this.sessionInfo)
}

/*
 * Set the user info tha own the acount.
 */
SessionPanel.prototype.setUserInfo = function (userInfo) {
	// The current user
	this.userInfo = userInfo
	this.nameDisplay.element.innerHTML = userInfo.M_firstName
}

/*
 * Display information about a given session.
 */
var SessionInfoPanel = function (parent, sessionInfo, accountInfo, displayStateMenu, diplayCloseSessionMenu) {
	this.parent = parent
	this.sessionInfo = sessionInfo
	this.accountInfo = accountInfo
	this.id = randomUUID()

	// So here I will create the information about the current session.
	this.sessionPanel = this.parent.appendElement({ "tag": "div", "class": "active_session" }).down()

	this.sessionState = this.sessionPanel.appendElement({ "tag": "div" }).down()
	this.sessionState.appendElement({ "tag": "i", "class": "fa fa-desktop active" }).down()

	// Set the computer name...
	if (this.sessionInfo.M_computerRef != undefined) {
		// The computer name...
		if (this.sessionInfo["set_M_computerRef_" + this.sessionInfo.M_computerRef + "_ref"] != undefined) {
			this.sessionInfo["set_M_computerRef_" + this.sessionInfo.M_computerRef + "_ref"](
				function (sessionPanel) {
					return function (computer) {
						computerName = computer.M_name
						if (computerName.length == 0) {
							computerName = "localhost"
						}
						sessionPanel.appendElement({ "tag": "div" }).down()
							.appendElement({ "tag": "span", "innerHtml": computerName })
					}
				} (this.sessionPanel))
		} else {
			this.sessionPanel.appendElement({ "tag": "div" }).down()
				.appendElement({ "tag": "span", "innerHtml": this.sessionInfo.M_computerRef })
		}
	}



	// Set the session state.
	if (this.sessionInfo.M_sessionState == 1) {
		this.sessionState.element.firstChild.className = "fa fa-desktop active"
	} else if (this.sessionInfo.M_sessionState == 2) {
		this.sessionState.element.firstChild.className = "fa fa-desktop away"
	} else if (this.sessionInfo.M_sessionState == 3) {
		this.sessionState.element.firstChild.className = "fa fa-desktop innactive"
	}

	// Display state select menu...
	if (displayStateMenu == true) {

		this.sessionState.element.onmouseover = function () {
			this.style.cursor = "pointer"
		}

		this.sessionState.element.onmouseleave = function () {
			this.style.cursor = "default"
		}

		// Here I will create the list of available state here...
		this.sessionStateMenu = new Element(document.getElementsByTagName("body")[0], { "tag": "div", "class": "session_state_menu" })
		this.sessionStateMenu.element.onclick = function () {
			this.style.display = "none"
		}

		this.sessionState.element.onclick = function (sessionPanel) {
			return function () {
				if (sessionPanel.sessionStateMenu.element.style.display != "block") {
					sessionPanel.sessionStateMenu.element.style.display = "block"
					var sessionStateMenuCoord = getCoords(this)
					sessionPanel.sessionStateMenu.element.style.top = sessionStateMenuCoord.top + this.offsetHeight + "px"
					sessionPanel.sessionStateMenu.element.style.left = sessionStateMenuCoord.left + "px"
				} else {
					sessionPanel.sessionStateMenu.element.style.display = "none"
				}
			}
		} (this)

		// The session states...
		this.sessionStateOnline = this.sessionStateMenu.appendElement({ "tag": "div", "class": "state_select", "style": "display:none;" }).down()
		this.sessionStateOnline.appendElement({ "tag": "div", "style": "display:inline;" }).down()
			.appendElement({ "tag": "i", "class": "fa fa-desktop active" }).up()
			.appendElement({ "tag": "span", "id": "SessionState_Online" })

		this.sessionStateAway = this.sessionStateMenu.appendElement({ "tag": "div", "class": "state_select", "style": "display:table;" }).down()
		this.sessionStateAway.appendElement({ "tag": "div", "style": "display:inline;" }).down()
			.appendElement({ "tag": "i", "class": "fa fa-desktop away" }).up()
			.appendElement({ "tag": "span", "id": "SessionState_Away" })

		this.sessionStateOffline = this.sessionStateMenu.appendElement({ "tag": "div", "class": "state_select", "style": "display:table;" }).down()
		this.sessionStateOffline.appendElement({ "tag": "div", "style": "display:inline;" }).down()
			.appendElement({ "tag": "i", "class": "fa fa-desktop innactive" }).up()
			.appendElement({ "tag": "span", "id": "SessionState_Offline" })

		// Set the element text...
		server.languageManager.setElementText(this.sessionStateMenu.getChildById("SessionState_Online"), "SessionState_Online")
		server.languageManager.setElementText(this.sessionStateMenu.getChildById("SessionState_Away"), "SessionState_Away")
		server.languageManager.setElementText(this.sessionStateMenu.getChildById("SessionState_Offline"), "SessionState_Offline")

		// Now I will set the action on those menu...
		this.sessionStateOnline.element.onclick = function (sessionStateOnline, sessionStateAway, sessionStateOffline) {
			return function () {
				sessionStateOnline.element.style.display = "none"
				sessionStateAway.element.style.display = "table"
				sessionStateOffline.element.style.display = "table"
				var state = 1
				server.sessionManager.updateSessionState(state,
					// Success callback.
					function () {
						/* The event will be catch by the listener **/
					},
					// Error callback.
					function () {
						/* Nothing todo here */
					}, null)
			}
		} (this.sessionStateOnline, this.sessionStateAway, this.sessionStateOffline)

		this.sessionStateAway.element.onclick = function (sessionStateOnline, sessionStateAway, sessionStateOffline) {
			return function () {
				sessionStateAway.element.style.display = "none"
				sessionStateOnline.element.style.display = "table"
				sessionStateOffline.element.style.display = "table"
				// Here I will update the session state...
				var state = 2
				server.sessionManager.updateSessionState(state,
					function () {
						/* The event will be catch by the listener **/
					}, null)
			}
		} (this.sessionStateOnline, this.sessionStateAway, this.sessionStateOffline)

		this.sessionStateOffline.element.onclick = function (sessionStateOnline, sessionStateAway, sessionStateOffline) {
			return function () {
				sessionStateOffline.element.style.display = "none"
				sessionStateOnline.element.style.display = "table"
				sessionStateAway.element.style.display = "table"

				// Here I will update the session state...
				var state = 3
				server.sessionManager.updateSessionState(state,
					function () {
						/* The event will be catch by the listener **/
					},
					function () {
						/* Nothing here */
					}
					, null)
			}
		} (this.sessionStateOnline, this.sessionStateAway, this.sessionStateOffline)
	}

	// That option is use to give the chance to the user to close other session
	// at distance.
	if (diplayCloseSessionMenu == true) {

		this.sessionState.element.onmouseover = function () {
			this.style.cursor = "pointer"
		}

		this.sessionState.element.onmouseleave = function () {
			this.style.cursor = "default"
		}

		// Here I will create the list of available state here...
		this.sessionStateMenu = this.sessionState.appendElement({ "tag": "div", "class": "session_state_menu", "style": "display:none;" }).down()
		var sessionMenuIsShow = false
		this.sessionState.element.onclick = function (sessionPanel, isShow) {
			return function () {
				if (!isShow) {
					sessionPanel.sessionStateMenu.element.style.display = "block"
				} else {
					sessionPanel.sessionStateMenu.element.style.display = "none"
				}

				isShow = !isShow
			}
		} (this, sessionMenuIsShow)

		// The session states...
		this.sessionCloseBtn = this.sessionStateMenu.appendElement({ "tag": "div", "class": "state_select" }).down()

		this.sessionCloseBtn.appendElement({ "tag": "div", "style": "display:inline;" }).down()
			.appendElement({ "tag": "i", "class": "fa fa-sign-out", "style": "color: darkgrey;" }).up()
			.appendElement({ "tag": "span", "id": "Session_Close_" + this.id })

		server.languageManager.setElementText(this.sessionStateMenu.getChildById("Session_Close_" + this.id), "Session_Close")

		// Now I will set the action on those menu...
		this.sessionCloseBtn.element.onclick = function (sessionInfo) {
			return function () {
				// So here i will call the logout event on the session...
				server.sessionManager.logout(sessionInfo.M_id,
					function () {
						//The event will be catch by the listener
					},
					function () {
						// nothing todo here...
					}
					, null)
			}
		} (this.sessionInfo)

	}

	// Now the elspase session time...
	this.sessionPanel.appendElement({ "tag": "div" }).down()
		.appendElement({ "tag": "i", "class": "fa fa-clock-o", "style": "color:#657383;" }).down()
	this.sessionTimer = this.sessionPanel.appendElement({ "tag": "span" }).down()

	// Here I will connect the time session display.
	// That time display the time of the current session time...
	setInterval(function (sessionPanel) {
		return function () {
			// Here I will calculate the session time...
			var sessionTimer = sessionPanel.sessionTimer
			var statusTime = sessionPanel.sessionInfo.M_statusTime
			var statusTimeStr = getTimeSinceStr(statusTime)
			sessionTimer.element.innerHTML = statusTimeStr
		}
	} (this), 1000)

	return this
}