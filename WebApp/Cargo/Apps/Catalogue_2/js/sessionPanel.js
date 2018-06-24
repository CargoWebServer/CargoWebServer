var sessionPanelText = {
    "en": {
        "connect-lbl": " Login",
        "online-opt": " online",
        "offline-opt": " offline",
        "away-opt": " away",
        "login-username": "user name",
        "login-pwd" :"password",
        "login-btn" :"login",
        "remember-me-lbl" : "Remember me"
    },
    "fr": {
        "connect-lbl": " Se connecter",
        "online-opt": " en ligne",
        "offline-opt": " hors ligne",
        "away-opt": " absent",
        "login-username": "nom d'utilisateur",
        "login-pwd" :"mot de passe",
        "login-btn" :"connection",
        "remember-me-lbl" : "Rester connecter"
    },
    "es" : {
        "connect-lbl" : "  conectarse",
        "online-opt": " en línea",
        "offline-opt": " fuera de línea",
        "away-opt": " ausente",
        "login-username": "nombre del usuario",
        "login-pwd" :" contraseña",
        "login-btn" :"conexión",
        "remember-me-lbl" : "mantenerse en contacto"
    }
};


// Internationalisation.
server.languageManager.appendLanguageInfo(sessionPanelText);

var SessionPanel = function(parent){
    
    // keep track of the active session.
    this.activeSession = null
    
    this.panel = new Element(parent, {"tag":"div", "class":"dropdown", "style":"margin-right: 5px;"})
    this.btn = this.panel.appendElement({"tag":"button", "class":"btn btn-outline-secondary dropdown-toggle", "data-toggle":"dropdown", "style":""}).down();
    
    // dropdown-toggle
    // Now I will set the text when the state is disconnected.
    this.disconnected = new Element(null, {"tag":"a"})
    this.disconnected.appendElement({"tag":"span", "id":"connect-lbl"})
        .appendElement({"tag":"span", "class":"fa fa-sign-in", "style":"padding-left: 5px;"})
        
    // set the disconnected state at first.
    this.btn.element.appendChild(this.disconnected.element)
    
    // where the menu goes...
    this.menu = this.panel.appendElement({"tag":"div", "class":"dropdown-menu bg-dark", "role":"menu", "style":"font-size: 0.8rem;"}).down()
    
    // The login panel...
    this.loginPanel = new Element(null, {"tag":"div","style":"min-width: 215px; padding: 2px"});
    
    this.loginPanel.appendElement({"tag":"div", "class":"input-group", "style":"padding: 2px"}).down()
        .appendElement({"tag":"div", "class":"input-group-prepend"}).down()
        .appendElement({"tag":"span", "class":"input-group-text"}).down()
        .appendElement({"tag":"i", "class":"fa fa-user"}).up().up()
        .appendElement({"tag":"input", "id":"login-username", "type":"text", "autofocus":"autofocus", "class":"form-control", "name":"username"})
    
    this.loginPanel.appendElement({"tag":"div", "class":"input-group", "style":"padding: 2px"}).down()
        .appendElement({"tag":"div", "class":"input-group-prepend"}).down()
        .appendElement({"tag":"span", "class":"input-group-text"}).down()
        .appendElement({"tag":"i", "class":"fa fa-lock"}).up().up()
        .appendElement({"tag":"input", "id":"login-pwd", "class":"form-control", "type":"password", "name":"pwd"});
        
    // Now the login button.
    this.loginPanel.appendElement({"tag":"div", "class":"input-group", "style":"padding: 2px"}).down()
        .appendElement({"tag":"a","id":"login-btn", "class":"btn btn-success", "style":"width: 100%;"}).down();
        
    // The remember me button.
    this.loginPanel.appendElement({"tag":"div", "class":"input-group"}).down()
        .appendElement({"tag":"div", "class":"checkbox", "style":"padding-left: 5px; padding-top: 5px;"}).down()
        .appendElement({"tag":"label", "style":"margin-bottom: 0px;"}).down()
        .appendElement({"tag":"input", "id":"remember_me", "type":"checkbox", "name":"remember"})
        .appendElement({"tag":"span", "id":"remember-me-lbl"})
    
    this.menu.element.appendChild(this.loginPanel.element);
    
    // The sessionPanel
    this.sessionPanel = null;
    
    // The user sessions.
    this.userSessionsDiv = null;
    this.otherUserSessionsDiv = null;
    
    // The connected session information.
    this.conencted = null;
    
    // Now the action...
    var loginBtn = this.loginPanel.getChildById("login-btn")
    loginBtn.element.onclick = function(sessionPanel){
        return function(evt){
            // not remove the dropdown...
            evt.stopPropagation()
            server.sessionManager.login(sessionPanel.loginPanel.getChildById("login-username").element.value, sessionPanel.loginPanel.getChildById("login-pwd").element.value, "SafranLdap",
            // success callback.
            function(session, sessionPanel){
                // So here the user is connected.
                // I will display the user sessions.
                sessionPanel.btn.element.click();
                sessionPanel.openSession(session)
                
            },
            // error callback
            function(errObj, caller){

                if(errObj.dataMap.errorObj.M_id == "ENTITY_ID_DOESNT_EXIST_ERROR"){
                    var loginUserName = sessionPanel.loginPanel.getChildById("login-username");
                    loginUserName.element.disabled = true
                    sessionPanel.loginPanel.getChildById("login-pwd").element.value = "";
                    // remove the message if it exist...
                    var previousMessage = document.getElementById("ENTITY_ID_DOESNT_EXIST_ERROR")
                    if(previousMessage !== null){
                        previousMessage.parentNode.removeChild(previousMessage)
                    }
                    var errorDisplay = loginUserName.parentElement.parentElement.appendElement({"tag":"div", "id":"ENTITY_ID_DOESNT_EXIST_ERROR", "class":"text-danger", "innerHtml": errObj.message}).down()
                    setTimeout(function(loginUserName, errorDisplay) {
                        return function(){
                            try{
                                errorDisplay.element.parentNode.removeChild(errorDisplay.element)
                                loginUserName.element.disabled = false
                                loginUserName.element.className = "form-control is-invalid"
                                loginUserName.element.focus();
                                loginUserName.element.setSelectionRange(0, loginUserName.element.value.length);
                            }catch(e){
                                
                            }
                        }
                    }(loginUserName, errorDisplay), 2500);
                }else  if(errObj.dataMap.errorObj.M_id == "PASSWORD_MISMATCH_ERROR"){
                    var loginPwd = sessionPanel.loginPanel.getChildById("login-pwd");
                    loginPwd.element.value = "";
                    loginPwd.element.disabled = true;
                    // remove the message if it exist...
                    var previousMessage = document.getElementById("PASSWORD_MISMATCH_ERROR");
                    if(previousMessage !== null){
                        previousMessage.parentNode.removeChild(previousMessage);
                    }
                    var errorDisplay = loginPwd.parentElement.parentElement.appendElement({"tag":"div", "id":"PASSWORD_MISMATCH_ERROR", "class":"text-danger", "innerHtml": errObj.message}).down()
                    setTimeout(function(loginPwd, errorDisplay) {
                        return function(){
                            try{
                                errorDisplay.element.parentNode.removeChild(errorDisplay.element);
                                loginPwd.element.disabled = false;
                                loginPwd.element.className = "form-control is-invalid";
                                loginPwd.element.focus();
                            }catch(e){
                                
                            }
                        }
                    }(loginPwd, errorDisplay), 3500);
                }
            }, 
            // The sesion panel
            sessionPanel)
        }
    }(this)
    
    // The login event.
	server.sessionManager.attach(this, LoginEvent, function(evt, sessionPanel){
	    if(server.account === undefined){
	        return
	    }
	    
	    // if the session dosent exist I will create it.
	    var session = evt.dataMap.sessionInfo;
	    if(session.M_accountPtr === server.account.UUID){
	        sessionPanel.diplayUserSessions([session]);
	    }else{
	        sessionPanel.diplayOtherUserSessions([session])
	    }
	})
	
	// The logout event.
	server.sessionManager.attach(this, LogoutEvent, function(evt, sessionPanel){
	    var session = evt.dataMap.sessionInfo;
	    clearInterval(timers[session.UUID])
	    delete timers[session.UUID];
	    var sessionDiv = document.getElementById(session.M_id)
	    if(sessionDiv !== undefined){
	        try{
	            sessionDiv.parentNode.removeChild(sessionDiv);
	        }catch(err){
	            
	        }
	    }
	})
	
	// The sate change event.
	server.sessionManager.attach(this, StateChangeEvent, function(evt, sessionPanel){
	    var session = evt.dataMap.sessionInfo;
	    entities[session.UUID] = session;
	    var sessionRows = document.getElementsByName(session.M_id)
	    for(var i=0; i < sessionRows.length; i++){
	        if(session.M_sessionState == 1){
	            sessionRows[i].className = sessionRows[i].className.replace("text-warning","text-primary");
	            sessionRows[i].className = sessionRows[i].className.replace("text-secondary","text-primary");
	        }else if(session.M_sessionState == 2){
	            sessionRows[i].className = sessionRows[i].className.replace("text-primary","text-warning");
	            sessionRows[i].className = sessionRows[i].className.replace("text-secondary","text-warning");
	        }else if(session.M_sessionState == 3){
	            sessionRows[i].className = sessionRows[i].className.replace("text-primary","text-secondary");
	            sessionRows[i].className = sessionRows[i].className.replace("text-warning","text-secondary");
	        }
	    }
	})
	
    return this;
}

/** 
 * Open a new session.
 */
SessionPanel.prototype.openSession = function(session){
    // So here the sesion is Open...
    this.activeSession = session;

    // Remove the login panel and display the session panel.
    this.menu.removeAllChilds();
    this.menu.element.innerHTML = ""
    this.btn.removeAllChilds();
    this.btn.element.innerHTML = ""
    this.btn.element.className = "btn dropdown-toggle bg-dark"
    this.sessionPanel = null;
    
    // display the current session information. 
    this.connected = new Element(null, {"tag":"a", "class":"text-primary", "name": session.M_id})
    
    this.connected.appendElement({"tag":"span", "class":"fa fa-user"})
        .appendElement({"tag":"span", "id":"accoutId", "style":"padding-left: 5px;"})
    
    // So here I will set the user information
    server.entityManager.getEntityByUuid(session.M_accountPtr, false, 
        // success callback
        function(account, sessionPanel){
            server.account = account;
            server.entityManager.getEntityByUuid(account.M_userRef, false,
                // success callback
                function(user, sessionPanel){
                    server.user = user;
                    sessionPanel.connected.getChildById("accoutId").element.innerHTML = user.M_firstName + " " + user.M_lastName;
                    sessionPanel.diplayUserSessions(server.account.M_sessions);

                
                    
                    // Now the network event.
                    var evtInfo = {
                        "TYPENAME": "Server.MessageData",
                        "Name": "welcomeEvent",
                        "Value": {
                            "user": user,
                            "sessionId" : server.sessionId
                        }
                    }
                    // Also broadcast the event over the network...
                    //server.eventHandler.broadcastEvent(evt)
                    server.eventManager.broadcastEventData(
                        welcomeEvent,
                        catalogMessage,
                        [evtInfo],
                        function () { },
                            function () { },
                            undefined
                        )
                    
                    
                    // Now I will display session of other user's.
                    // first I need to get the list of user's
                    server.entityManager.getEntities("CargoEntities.Session", "CargoEntities", "", 0, -1, [], true, false,
                        function(index, total, caller){
                            
                        },
                        function(sessions, caller){
                            sessionPanel.diplayOtherUserSessions(sessions)
                        },
                        function(errObj, caller){
                            
                        },
                        sessionPanel)
                    },
                    // error callback
                    function(){
                        
                    }, sessionPanel)
        },
        // error callcack
        function(error, sessionPanel){
            
        },this)
    
    // Here The current computer.
    this.btn.element.appendChild(this.connected.element)
}

// Display the list of active user session.
SessionPanel.prototype.diplayUserSessions = function(sessions){
    if(server.account == undefined){
        return;
    }
    
    if(this.userSessionsDiv === null){
        this.userSessionsDiv = this.menu.appendElement({"tag":"div", "class":"container", "style":"min-width: 265px;"}).down();
    }
    
    for(var i=0; i < sessions.length; i++){
        // I will create the buttons container.
        if(sessions[i] != undefined){
            if(server.sessionId == sessions[i].M_id){
            // put at first...
            this.diplayUserSession(sessions[i], this.userSessionsDiv.prependElement({"tag":"div", "name": sessions[i].M_id, "class":"row", "style":"justify-content: left; align-items: baseline; flex-wrap: nowrap;"}).down());
            }else{
                this.diplayUserSession(sessions[i], this.userSessionsDiv.appendElement({"tag":"div", "id": sessions[i].M_id, "name": sessions[i].M_id, "class":"row", "style":"justify-content: left; align-items: baseline; flex-wrap: nowrap;"}).down());
            }
        }
        
    }
}

// Display other users session
SessionPanel.prototype.diplayOtherUserSessions = function(sessions){
    if(server.account == undefined){
        return;
    }
    
    if(this.otherUserSessionsDiv === null){
        this.otherUserSessionsDiv = this.menu.appendElement({"tag":"div", "class":"container", "style":"min-width: 265px; border-top: 1px solid #2c3135; margin-top: 5px; padding-top: 5px;"}).down();
    }
    
    for(var i=0; i < sessions.length; i++){
        if(sessions[i].M_accountPtr != server.account.UUID){
            // I will create the buttons container.
            this.diplayUserSession(sessions[i], this.otherUserSessionsDiv.appendElement({"tag":"div", "id": sessions[i].M_id, "name": sessions[i].M_id, "class":"row", "style":"justify-content: left; align-items: baseline; flex-wrap: nowrap;"}).down());
        }
    }
}

// keep track of interval function.
var timers = {}
// display user session.
SessionPanel.prototype.diplayUserSession = function(session, div){
    // The case the session dosent exist only
    if(document.getElementById(session.UUID) != undefined){
        return
    }
    
    // if the session is the current session.
    var isOwner = session.M_accountPtr == server.account.UUID
    
    // for local session.
    if(session.M_accountPtr == server.account.UUID){
        var computerDiv = div.appendElement({"tag":"div", "class":"col-sm-5"}).down()
        server.entityManager.getEntityByUuid(session.M_computerRef, false,
            // Success callback
            function(computer, caller){
                caller.sessionPanel.createStateSelector(caller.session, computer, caller.computerDiv)
            },
            // Error callback
            function(errObj, caller){
                var computer = new CargoEntities.Computer();
                computer.M_name = "unknow"
                caller.sessionPanel.createStateSelector(caller.session, computer, caller.computerDiv)
            }, 
            {"sessionPanel":this, "session":session, "computerDiv":computerDiv})
    }else{
        // for distant session...
        var userDiv = div.appendElement({"tag":"div", "class":"col-sm-5"}).down()
        server.entityManager.getEntityByUuid(session.M_accountPtr, false,
            // Success callback
            function(user, caller){
                caller.userDiv.appendElement({"tag":"span", "class":"fa fa-user"})
                    .appendElement({"tag":"span", "style": "padding-left: 5px; font-size: .7em;", "innerHtml":user.M_id.split("@")[0]})
            },
            // Error callback
            function(errObj, caller){
                
            }, 
            {"sessionPanel":this, "session":session, "userDiv":userDiv})
    }
    
    
    if(session.M_sessionState == 1){
        div.element.className = "row text-primary";
    }else if(session.M_sessionState == 2){
        div.element.className = "row text-warning";
    }else if(session.M_sessionState == 3){
        div.element.className = "row text-secondary";
    }
    
    var div_ = div.appendElement({"tag":"div", "class":"col-sm-6"}).down();
    div_.appendElement({ "tag": "span", "class": "fa fa-clock-o", "style": "padding-right: 5px;"}).down();
	var timerDiv  = div_.appendElement({"tag":"span","style": ""}).down();
	
	timers[session.UUID] = setInterval(function (sessionUUID, timerDiv) {
		return function () {
			// Here I will calculate the session time...
			var session = entities[sessionUUID]
	        timerDiv.element.innerHTML = getTimeSinceStr(session.M_statusTime)
		}
	} (session.UUID, timerDiv), 1000);
	
	if(isOwner === true){
	    var signOut = div.appendElement({"tag": "span", "class": "col-sm-1 fa fa-sign-out",  "style": "padding-left: 5px; text-align: right;" }).down()
	    signOut.element.onclick = function(session){
	        return function(evt){
	            evt.stopPropagation() // keep the panel open.
	            server.sessionManager.logout(session.M_id);
	        }
	    }(session);
	}
}

SessionPanel.prototype.createStateSelector = function(session, computer, div){
    if(server.sessionId == session.M_id){
        // Here I will display the session state selector...
        var sessionState = div.appendElement({"tag":"select","name":session.M_id, "style":"border: none;"}).down()
        sessionState.appendElement({"tag":"option", "id":"online-opt", "value":"1", "class":"text-primary"})
        sessionState.appendElement({"tag":"option", "id":"away-opt", "value":"2", "class":"text-warning"})
        sessionState.appendElement({"tag":"option", "id":"offline-opt", "value":"3", "class":"text-secondary"})
        // Set the selection style.
        sessionState.element.value = session.M_sessionState 
        if(session.M_sessionState == 1){
            sessionState.element.className = "selectpicker bg-dark text-primary";
        }else if(session.M_sessionState == 2){
            sessionState.element.className = "selectpicker bg-dark text-warning";
        }else if(session.M_sessionState == 3){
            sessionState.element.className = "selectpicker bg-dark text-secondary";
        }
        
        sessionState.element.value = session.M_sessionState;
        sessionState.element.onclick = function(evt){
            evt.stopPropagation()
        }
        
        // Now the session state select action.
        sessionState.element.onchange =  function(){
            server.sessionManager.updateSessionState(parseInt(this.value), function(){}, function(){}, {})
        }
    }else{
        // I will simply display the computer icone and it respective name.
        div.appendElement({"tag":"span", "class":"fa fa-desktop"})
            .appendElement({"tag":"span", "style": "padding-left: 5px; font-size: .7em;", "innerHtml":computer.M_name})
    }

}
