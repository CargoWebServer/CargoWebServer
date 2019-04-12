
// Must be global variables.
var intervals = []
var intervalsMap = {}

var TaskInstancesExplorer = function (parent) {

    this.parent = parent
    this.panel = parent.appendElement({ "tag": "div", "class": "data_explorer" }).down()

    // set the scrolling shadow...
    this.panel.element.onscroll = function (header) {
        return function () {
            var position = this.scrollTop;
            if (this.scrollTop > 0) {
                if (header.className.indexOf(" scrolling") == -1) {
                    header.className += " scrolling"
                }
            } else {
                header.className = header.className.replaceAll(" scrolling", "")
            }
        }
    }(this.parent.element.firstChild)

    this.parent.element.firstChild.style.borderBottom = "1px solid gray"

    // Set the resize event.
    window.addEventListener('resize',
        function (taskInstancesExplorer) {
            return function () {
                taskInstancesExplorer.resize()
            }
        }(this), true);

    // Event listener
    // Here I will connect the listeners... cancel and new, end will be manage by the setTaskTimer function...
    server.configurationManager.attach(this, NewTaskEvent, function (evt, taskInstancesExplorer) {
        // The task information to display.
        taskInstancesExplorer.setTask(evt.dataMap.taskInfos, intervals, intervalsMap, intervals.length)
    })

    // Update the task informations...
    server.configurationManager.attach(this, UpdateTaskEvent,
        function (evt, taskInstancesExplorer) {
            var instanceInfos = evt.dataMap.taskInfos
            var row = document.getElementById(instanceInfos.UUID + "_task_instance")
            if (row != undefined) {
                if (instanceInfos.EndTime != undefined || instanceInfos.CancelTime != undefined) {
                    if (intervalsMap[instanceInfos.TaskId] != undefined) {
                        clearInterval(intervalsMap[instanceInfos.TaskId]);
                        intervals = intervals.splice(intervalsMap[instanceInfos.TaskId], 1)
                        delete intervalsMap[instanceInfos.TaskId]
                    }
                }
                row.parentNode.removeChild(row)
                taskInstancesExplorer.setTask(instanceInfos, intervals, intervalsMap, intervals.length)
            }
        })

    // Initialyse the list of instance.
    // Here I will get the list of active task on the server and display it in the list.
    server.configurationManager.getTaskInstancesInfos(
        /** Success callback */
        function (results, caller) {
            // Here I will set the list of active task...
            for (var i = 0; i < results.length; i++) {
                caller.taskInstancesExplorer.setTask(results[i], intervals, intervalsMap, i)
            }
        },
        /** Error callback */
        function (errObj, caller) { },
        /** caller */
        { "taskInstancesExplorer": this })

    return this
}

TaskInstancesExplorer.prototype.resize = function () {
    var height = this.parent.element.offsetHeight - this.parent.element.firstChild.offsetHeight;
    this.panel.element.style.height = height - 10 + "px"
}

/**
  * Set the timer function.
  * @param {*} row 
  * @param {*} countdown 
  * @param {*} startTime 
  * @param {*} intervals 
  * @param {*} index 
  */
TaskInstancesExplorer.prototype.setTaskTimer = function(instanceInfos, row, countdown, startTime, intervals, intervalsMap, index) {
    // Get todays date and time
    var now = new Date().getTime();
    // Find the distance between now an the count down date
    var distance = startTime - now;
    // Time calculations for days, hours, minutes and seconds
    var days = Math.floor(distance / (1000 * 60 * 60 * 24));
    var hours = Math.floor((distance % (1000 * 60 * 60 * 24)) / (1000 * 60 * 60));
    var minutes = Math.floor((distance % (1000 * 60 * 60)) / (1000 * 60));
    var seconds = Math.floor((distance % (1000 * 60)) / 1000);

    // Display the result in the element with id="demo"
    var delay = ""
    if (days > 0) {
        delay += days + "d "
    }
    if (hours > 0) {
        delay += hours + "h "
    }
    if (minutes > 0) {
        delay += minutes + "m "
    }

    delay += seconds + "s "

    countdown.element.innerHTML = delay
    // If the count down is finished, write some text 
    if (distance < 0) {
        clearInterval(intervals[index]);
        intervals = intervals.splice(index, 1)
        delete intervalsMap[instanceInfos.TaskId]
        try{
            row.element.parentNode.removeChild(row.element)
        }catch(err){
            // nothing to do here.
        }
        // Set it back with it new status and delete callback.
        instanceInfos.EndTime = new Date().getTime() / 1000;
        this.setTask(instanceInfos, intervals, intervalsMap, intervals.length)
    }
}

/**
  * Set task information.
  * @param {*} instanceInfos 
  * @param {*} intervals 
  * @param {*} intervalsMap 
  * @param {*} index 
  */
TaskInstancesExplorer.prototype.setTask = function (instanceInfos, intervals, intervalsMap, index) {
    var row = this.panel.appendElement({ "id": instanceInfos.UUID + "_task_instance", "tag": "div", "style": "display: table-row; width: 100%;", "class": "entity" }).down();
    row.appendElement({ "tag": "div", "style": "display: table-cell; padding: 2px;", "innerHtml": instanceInfos.TaskId })
    var countdown = row.appendElement({ "tag": "div", "style": "display: table-cell; padding: 2px 2px 2px 10px; width: 100%;" }).down()
    var cancelTaskBtn = row.appendElement({ "tag": "div", "class": "entity_panel_header_button" }).down()
        .appendElement({ "tag": "i", "class": "fa fa-close" }).down()

    countdown.element.innerHTML = instanceInfos.Status
    
    if (instanceInfos.StartTime != 0) {
        // In case that the task instance has a start time.
        // In case that the task in not already running
        var startTime = new Date(instanceInfos.StartTime * 1000).getTime()
        var now = new Date().getTime();
        if (startTime - now > 0 && instanceInfos.CancelTime == 0) {
            // Here if the user click the button it will cancel the task.
            cancelTaskBtn.element.onclick = function (instanceInfos) {
                return function () {
                    // Cancel the task...
                    server.configurationManager.cancelTask(instanceInfos.UUID,
                        /** The success callbacak */
                        function () {

                        },
                        /** The error callback */
                        function () {

                        }, {})
                }
            }(instanceInfos)

            intervals[index] = setInterval(function (instanceInfos, row, countdown, startTime, index, taskExplorer) {
                return function () {
                    taskExplorer.setTaskTimer(instanceInfos, row, countdown, startTime, intervals, intervalsMap, index)
                }
            }(instanceInfos, row, countdown, startTime, index, this), 1000);

            // keep the interval...
            intervalsMap[instanceInfos.TaskId] = intervals[index]

        } else {
            if (instanceInfos.EndTime == 0) {
                if (instanceInfos.Error != null) {
                    countdown.element.title = new Date(instanceInfos.EndTime * 1000).toISOString()
                } else {
                    // If the task failed.
                    countdown.element.title = instanceInfos.Error
                }
                cancelTaskBtn.element.onclick = function (row) {
                    return function () {
                        row.element.parentNode.removeChild(row.element)
                    }
                }(row)
            } else if (instanceInfos.CancelTime != 0) {
                countdown.element.title = new Date(instanceInfos.CancelTime * 1000).toISOString()
                cancelTaskBtn.element.onclick = function (row) {
                    return function () {
                        row.element.parentNode.removeChild(row.element)
                    }
                }(row)
            }
        }
    } else {
        if (instanceInfos.CancelTime != 0) {
            countdown.element.title = new Date(instanceInfos.CancelTime * 1000).toISOString()
            cancelTaskBtn.element.onclick = function (row) {
                return function () {
                    row.element.parentNode.removeChild(row.element)
                }
            }(row)
            cancelTaskBtn.element.onclick = function (row) {
                return function () {
                    row.element.parentNode.removeChild(row.element)
                }
            }(row)
        } else {
            // Here If the user click the button it will stop the task.
            cancelTaskBtn.element.onclick = function (instanceInfos) {
                return function () {
                    server.configurationManager.cancelTask(instanceInfos.UUID,
                        /** The success callbacak */
                        function () {

                        },
                        /** The error callback */
                        function () {

                        }, {})
                }
            }(instanceInfos)
        }
    }
}