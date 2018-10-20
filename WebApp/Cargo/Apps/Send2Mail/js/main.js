// Keep the list of users by id.
var usersByLetter = {};

var MainPage = function(){
    // So here I will create the split section...
    var left = new Element(document.getElementsByTagName("body")[0], {"tag":"div", "class":"split left"})
    var right = new Element(document.getElementsByTagName("body")[0], {"tag":"div", "class":"split right"})
    
    // Now I will split the right div in tow.
    var top = right.appendElement({"tag":"div", "class":"split", "style":"flex-direction: row; height: 50%;"}).down()
        .appendElement({"tag":"div", "class" :"item-wrapper"}).down();
        
    var bottom = right.appendElement({"tag":"div", "class":"split", "style":"flex-direction: row; height: 50%; top: 50%;border-top: 4px solid #353535;"}).down()
        .appendElement({"tag":"div", "class" :"item-wrapper"}).down();
        
    // So now I will display the list of users to the screen.
    var letters = Object.keys(usersByLetter).sort();
    var lettersDiv = left.appendElement({"tag":"div", "class" :"item-wrapper"}).down();
    for(var i=0; i < letters.length; i++){
        var letterDiv = lettersDiv.appendElement({"tag":"div", "id":letters[i], "class":"item", "innerHtml":letters[i]}).down()
        letterDiv.element.onclick = function(users, lettersDiv, left, top){
            return function(){
                lettersDiv.element.style.display = "none";
                var usersDiv = left.appendElement({"tag":"div", "class" :"item-wrapper"}).down();
                var backDiv = usersDiv.appendElement({"tag":"div", "class":"item", "style":"width: 160px; height:160px; line-height: 30px; font-size: 12pt;"}).down()
                backDiv.appendElement({"tag":"i", "class":"fa fa-arrow-left", "style":"font-size: 18pt; margin-top: 50px;"})
                    .appendElement({"tag":"div", "innerHtml":"retour"})
                   
                backDiv.element.onclick = function(usersDiv, lettersDiv){
                    return function(){
                        usersDiv.element.parentNode.removeChild(usersDiv.element);
                        lettersDiv.element.style.display = "";
                    }
                }(usersDiv, lettersDiv) 

                for(var i=0; i < users.length; i++){
                    var userDiv = usersDiv.appendElement({"tag":"div", "class":"item", "style":"width: 160px; height:160px; line-height: 30px; font-size: 12pt; position: relative;"}).down()
                    userDiv
                        .appendElement({"tag":"i", "class": "fa fa-times", "style":"font-size: 12pt; position: absolute; top: 6px; right: 6px; display: none"})
                        .appendElement({"tag":"i", "class": "fa fa-user", "style":"font-size: 20pt; margin-top: 10px;"})
                        .appendElement({"tag":"div", "innerHtml": users[i].M_id})
                        .appendElement({"tag":"div", "innerHtml": users[i].M_lastName + " " + users[i].M_firstName})
                        
                    // Now on clik user div event.
                    userDiv.element.onclick = function(usersDiv, top){
                        return function(){
                            var index = Array.prototype.indexOf.call(this.parentNode.children, this);
                            this.parentNode.removeChild(this)
                            top.element.appendChild(this)
                            this.firstChild.style.display = "";
                            this.firstChild.onclick = function(userDiv, parentNode, top, index){
                                return function(evt){
                                    evt.stopPropagation();
                                    userDiv.parentNode.removeChild(userDiv);
                                    usersDiv.element.insertBefore(userDiv, usersDiv.element.children[index]);
                                    this.style.display = "none";
                                }
                            }(this, usersDiv, top, index)
                        }
                    }(usersDiv, top)
                }
            }
        }(usersByLetter[letters[i]], lettersDiv, left, top)
    }
}

function main(){
    server.entityManager.getEntities("CargoEntities.User", "CargoEntities", null, 0, -1, ["M_lastName"], true, true,
    function(index,total, caller){
        
    },
    function(results, caller){
        users = results.sort(function(a, b){
            return a.M_lastName.localeCompare(b.M_lastName)
        })
        
        for (i = 0; i < 26; i++) {
            var letter =  (i+10).toString(36).toUpperCase()
            usersByLetter[letter] = []
            while(users.length > 0){
                if(users[0].M_lastName.length > 0){
                    if(!users[0].M_lastName.startsWith(letter)){
                        // remove it the list...
                        break; // exit while loop.
                    }else{
                        usersByLetter[letter].push(users[0])
                    }
                }
                users.shift()
            }
            
        }
        
        // Now I will compact letters whit few user.
        var minUserByLetter = 20;
        var combinedUsers = {};
        for(var i=0; i < Object.keys(usersByLetter).length; i++){
            var letters = letter = Object.keys(usersByLetter)[i];
            if(letter.length == 1){
                var users = usersByLetter[letter]
                while(users.length < minUserByLetter){
                    // Here I must regroup the users with the next letter.
                    if(i < Object.keys(usersByLetter).length -1){
                        i++
                        var letter_ = Object.keys(usersByLetter)[i]
                        users = users.concat(usersByLetter[letter_])
                        letters +=  letter_
                    }else{
                        break
                    }
                }
                if(letters.length > 1){
                    combinedUsers[letters] = users;
                }
            }
            
        }
        
        for(var letter in combinedUsers){
            for(var i=0; i < letter.length; i++){
                delete usersByLetter[letter[i]]
            }
            usersByLetter[letter] = combinedUsers[letter]
        }

        new MainPage();
    },
    function(error, caller){
        
    }, {})
 
}