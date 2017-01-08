/*
 * Created by Dave on 12/25/2014.
 */

// Error message.
var languageInfo = {
    "en":{
        "InvalidEmail": "Invalid Email",
        "PasswordConfirmError" : "The passwords you have entered don't match",
        "AccountAlreadyExists" : "Account already exists",
        "ACCOUNT_DOESNT_EXIST_ERROR": "Account doesn't exist",
        "PASSWORD_MISMATCH_ERROR" : "PasswordMissmatch"
    },
    "fr":{
        "InvalidEmail": "Email invalide",
        "PasswordConfirmError" : "Le mot de passe entré n'est pas le bon",
        "AccountAlreadyExists" : "Un acompte avec le même nom existe déjà",
        "ACCOUNT_DOESNT_EXIST_ERROR" : "Le nom d'usagé entré est inexistant",
        "PASSWORD_MISMATCH_ERROR" : "Le mot de passe ne correspond pas"
    }
}

var ErrorManager = function(id, name){
    // The event id.
    this.id = id

    // The name of the event
    this.name = name
    this.observers = {}

    return this
}

ErrorManager.prototype.attach = function(observer, errorId, reportErrorFct){
    if(this.observers[errorId]==undefined){
        this.observers[errorId] = []
    }
    this.observers[errorId].push(observer)
    if (observer.reportErrorFunctions == undefined){
        observer.reportErrorFunctions = {}
    }
    observer.reportErrorFunctions[errorId] = reportErrorFct
}

ErrorManager.prototype.detach = function(observer, errorId){
    delete observer.reportErrorFunctions[this.id]
    this.observers[errorId].pop(observer)
}

ErrorManager.prototype.onError = function(err){
    var errObj = err.dataMap["errorObj"]
    if (this.observers[errObj.M_id] != undefined){
        for (var i=0; i<this.observers[errObj.M_id].length; i++){
            if (this.observers[errObj.M_id][i].reportErrorFunctions[errObj.M_id] != undefined){
                this.observers[errObj.M_id][i].reportErrorFunctions[errObj.M_id](errObj)
            }
        }
    }
}
