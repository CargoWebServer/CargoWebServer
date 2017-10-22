/* Logical Process:
1- Regarder s'il y a de nouveaux fichiers
    - S'il y en a on les ajoutes en retirant toutes les informations de son filename, puis de toutes les données qu'il contient
    - Insert les nouveaus fichiers en gardant son nouveau ID et insère son data. 
*/

function LaunchImportNewSQLData_Process() {
    // TODO Change the path to the real path here...
    var dir = "\\\\mon-filer-01\\data\\Departement Commun\\Buffer\\Pierre-Olivier\\15-SQL Proactiv\\02-Documentations\\SILMA_Export-Sample";
    //%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%
    //@param(dirPath, SuccesCallBack,errorCallBack, caller)    
    //%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%
    server.fileManager.readDir(dir,

        //%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%
        //Success d'obtention de tous les chemins de fichier
        //%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%
        function (results, caller) {

            var filenames = results[0];

            var paths = [];
            for (var fID = 0; fID < filenames.length; fID++) {
                paths[fID] = caller.dir + "\\" + filenames[fID];
            }

            //Obtient tous les fichiers dans la base de donnees
            var q = "SELECT ID, FilePath FROM [Eng_Dwg_Features].[dbo].[ProActiv_Files]"
            var fields = new Array()
            fields[0] = "int" // ID of the Entry
            fields[1] = "string" //[File Property] Chemin Absolue du fichier

            server.dataManager.read("Eng_Dwg_Features", q, fields, [],
                // Success Callback
                function (results, params) {

                    var initCallback = params[0];
                    var paths = params[1];
                    if (initCallback != null) {
                        initCallback(paths, results[0]);
                    }

                },
                // Progress Callback
                function () { },
                // Error Callback,
                function () { },
                [ProcessAnalyseImportation, paths])
        },
        function (errObj, caller) {

        }, { "dir": dir })
}

//===============================================================================================================================================
// Anlayse s'il y a de nouvelles donnees (fichiers) a etre importer
//===============================================================================================================================================
function ProcessAnalyseImportation(ExistigFilePath, sqlFilePaths) {
    //Loop: regarde chaque fichier contenu dans la base de donnee SQL et compare s'il est existant dans le dossier. si NON: on l'ajoute
    for (var i = 0; i < ExistigFilePath.length; i++) {

        //Le chemin n'existant pas, on l'INSERT dans le SQL
        var isExist = false
        for (var j = 0; j < sqlFilePaths.length; j++) {
            if (sqlFilePaths[j].indexOf(ExistigFilePath[i]) != -1) {
                isExist = true
                break
            }
        }
        if (!isExist) {

            //%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%
            //Process l'extraction des valeurs relier au nom du fichier
            //%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%
            var fPath = ExistigFilePath[i];
            var filename = fPath.split(".")[0];
            filename = filename.split("\\")[filename.split("\\").length - 1];


            //Searched Values
            var operation = workOrder = partType = aircraft = partnumber = "";  //From Filename
            var fileAuthors = fModifiedDateTime = "";   //From File Properites

            //Normally: [0] = Op, [1] = WorkOrder, [2] = specs, [3] = Date, [4] = Time, 
            var filename_Values = filename.split("_");

            //Operation
            if (0 <= filename_Values.length - 1) { operation = filename_Values[0]; }
            //WorkOrder
            if (1 <= filename_Values.length - 1) { workOrder = filename_Values[1]; }

            //Specs: Aircraft_Product + PartNumber
            if (2 <= filename_Values.length - 1) {
                //Separe les specifications
                var specs = filename_Values[2].split(" ");
                if (0 <= specs.length - 1) { partType = specs[0]; }
                if (1 <= specs.length - 1) { aircraft = specs[1]; }
                if (2 <= specs.length - 1) { partnumber = specs[2]; }
                //Le reste ne nous intéresse pas! car Souvent pas rempli!
            }

            //fileAuthors
            fileAuthors = "dummyAuthor";
            //!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!! TO DO!: Waiting for the Function
            //!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
            //!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!

            //fModifiedDateTime
            fModifiedDateTime = new Date();
            //!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!! TO DO!: Waiting for the Function
            //!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
            //!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!

            //----------------------------------------
            //Garde l'information sur le fichier
            //----------------------------------------
            //newFiles's Columns_____________
            //0- [FilePath]
            //1- [WorkOrder]
            //2- [Operation]
            //3- [Aircraft_Product]
            //4- [PartNumber]
            //5- [PartType]
            //6- [File_Author]
            //7- [File_ModifiedDateTime]

            var newFileInfo;       //Array de toutes les données/propritétés sur le fichier Excel
            var newFilesData = [];   //Array de toutes les données à l'intérieur du fichier Excel (fileID, Data[rowID][colID])

            newFileInfo = [1, fPath, workOrder, operation, aircraft, partnumber, partType, fileAuthors, fModifiedDateTime];

            var fileToExport_Caller = { "fileInfo": newFileInfo, "fileData": newFilesData };


            //%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%
            //Process l'extraction des valeurs à l'intérieur du fichier Excel
            //%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%

            //Vérifie l'Extension
            if (fPath.endsWith(".XLS") || fPath.endsWith(".xls") || fPath.endsWith(".XLSX") || fPath.endsWith(".xlsx")) {

                // Here I will use vb script to convert xls to csv.
                var excelPath = fPath
                var csvPath = fPath.replace("XLS", "csv").replace("xls", "csv").replace("XLSX", "csv").replace("xlsx", "csv")

                /*getFileInfos(fPath,
                    function () {
                        return function (fileInfos) {
                            console.log(fileInfos)
    
                        }
                    }())*/

                //@param(dirPath, exeParam, SuccesCallBack,errorCallBack, caller)    
                server.executeVbSrcript("xlsx2csv.vbs", [excelPath, csvPath],

                    //SUCESS:
                    //La conversion a été un succes et on obtient le chemin du CSV créé.
                    function (csvPath) {
                        return function (results, caller) {
                            // So here I will read the csv file.
                            server.fileManager.readCsvFile(csvPath,

                                //SUCESS:
                                function (csvPath) {
                                    return function (results, caller) {

                                        // TODO Process the result here...
                                        var fRawData = results[0]; //Data brute contenu dans le fichier Excel

                                        //------------------------------
                                        //Transforme le data Brute en contenu compact et plus clair: enlève espace supperflu
                                        //------------------------------
                                        var bIsLineEmpty = true;
                                        var feature, tolName, entryName, actual, nominal, minTol, maxTol, dev, out;

                                        //Loop au travers de chaque ligne pour savoir si elle est completement vide
                                        for (var rID = 0; rID < fRawData.length; rID++) {

                                            lineData = fRawData[rID]; //.split(";");

                                            if (lineData.length > 0) {

                                                //Use Regex to know if LIKE "Feature"
                                                if (/FEATURE.*/.test(lineData[0].toUpperCase())) {

                                                    feature = lineData[1];     //Obtient le nom du feature dans 2ere Colonne

                                                    //Obtient le nom de la tolérance (identification) qui sur la ligne du dessous
                                                    rID++;
                                                    lineData = fRawData[rID]; //.split(";");
                                                    var myRe = new RegExp('.+\\.\\d+', 'g');  //Obtient le nom avant les ".###"
                                                    var myArray = myRe.exec(lineData[1]);   //2e Colonne
                                                    tolName = myArray[0];

                                                    //Va chercher les valeurs associées à cette tolérance
                                                    //= actuanl, nominal, min, max, dev, out
                                                    rID++;
                                                    var limit = rID;
                                                    for (; rID <= limit + 5; rID++) {

                                                        lineData = fRawData[rID]; //.split(";");
                                                        entryName = tolName
                                                        var typeCote = ""
                                                        if (lineData[2] != " ") {
                                                            entryName = entryName + "_" + lineData[2];
                                                            typeCote = lineData[2];
                                                        }

                                                        if (lineData[3] != " ") {
                                                            actual = parseFloat(lineData[3]);
                                                        } else {
                                                            actual = 0.0;
                                                        }
                                                        if (lineData[4] != " ") {
                                                            nominal = parseFloat(lineData[4]);
                                                        } else {
                                                            nominal = 0.0;
                                                        }
                                                        if (lineData[5] != " ") {
                                                            minTol = parseFloat(lineData[5]);
                                                        } else {
                                                            minTol = 0.0;
                                                        }
                                                        if (lineData[6] != " ") {
                                                            maxTol = parseFloat(lineData[6]);
                                                        } else {
                                                            maxTol = 0.0;
                                                        }
                                                        if (lineData[7] != " ") {
                                                            dev = parseFloat(lineData[7]);
                                                        } else {
                                                            dev = 0.0;
                                                        }
                                                        if (lineData[8] != " ") {
                                                            out = parseFloat(lineData[8]);
                                                        } else {
                                                            out = 0.0;
                                                        }

                                                        if (lineData[3] != " ") {
                                                            var tolData = [entryName, feature, actual, nominal, minTol, maxTol, dev, out];
                                                            caller.fileData.push(tolData);
                                                        } else {
                                                            break
                                                        }
                                                    }

                                                }   //END: Feature Line Detected
                                            }   //END: C'est une ligne aux colonnes Valides
                                        }   //END: Finish Reading the CSV File

                                        //%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%
                                        //INSERT les données sur le Fichier Excel
                                        //%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%
                                        var query = "INSERT INTO Eng_Dwg_Features.dbo.ProActiv_Files (ID_Machine,[FilePath],[WorkOrder],"
                                            + "[Operation],[Aircraft_Product],[PartNumber],[PartType],[File_Author],[File_ModifiedDateTime]) VALUES (?,?,?,?,?,?,?,?,?)";

                                        server.dataManager.create("Eng_Dwg_Features", query, caller.fileInfo,

                                            //SUCCESS: File Added
                                            function (results, caller) {
                                                //%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%
                                                //INSERT les données contenues dans le fichier Excel
                                                //%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%
                                                var fileID = results[0];
                                                var fileData = caller.fileData

                                                var dataToExport = [fileID, "my name is TolName"];
                                                for (var i = 0; i < fileData.length; i++) {
                                                    fileData[i].unshift(fileID)
                                                    /* append this values ,[AddedDate],[LastModifiedDate], [Type_Cote]*/
                                                    query = "INSERT INTO Eng_Dwg_Features.dbo.ProActiv_Entry ([ID_File],[Entry_Name],[Entry_Feature],[Cote_Actual],[Cote_Nominal],[Cote_Tol_Min],[Cote_Tol_Max],[Cote_Dev],[Cote_Out] ) VALUES (?,?,?,?,?,?,?,?,?)";
                                                    server.dataManager.create("Eng_Dwg_Features", query, fileData[i],
                                                        //SUCCESS: Tolerance Added
                                                        function (results, caller) {
                                                            var newTolID = results[0];
                                                            console.log("newTolID Added = " + newTolID)
                                                        },
                                                        // Error callback here...
                                                        function (errorMsg, caller) {

                                                        },
                                                        {});
                                                }

                                            },
                                            // Error callback here...
                                            function (errorMsg, caller) {

                                            },

                                            //CALLER:
                                            caller)

                                        // Remove the temp CSV File
                                        //------------------------------
                                        server.fileManager.removeFile(csvPath)
                                    } //END: Succes Function Reading CSV File

                                }(csvPath)
                                ,

                                //FAILED
                                function (errObj, caller) {
                                    console.log(errObj)
                                },

                                //CALLER:
                                caller)//Autre méthode: {"path":csvPath})
                        }
                    }(csvPath)
                    ,

                    //FAILED/ERROR CALLBACK:
                    function (errObj, caller) {
                    },

                    //CALLER:
                    fileToExport_Caller)
            }

        } //END: si le fichier n'est pas dans SQL
    }
}

/**
 * 
 * @param {*} path 
 * @param {*} callback 
 */
function getFileInfos(path, callback) {
    server.runCmd("cmd", ["/K", "dir /Q " + path],
        // Success callback
        function (results, caller) {
            var values = results["result"].split(/\s+/);
            var author = values[23]
            var path = caller.path.replaceAll("\\", "\\\\")
            server.runCmd("cmd", ["/K", "wmic datafile where name='" + path + "' list full"],
                // Success callback
                function (results, caller) {
                    var values = results["result"].split(/\s+/);
                    var fileInfos = {}
                    for (var i = 0; i < values.length; i++) {
                        if (values[i].indexOf("=") != -1) {
                            var infos = values[i].split("=")
                            var propertie = infos[0]
                            var value = infos[1]
                            if (propertie == "LastAccessed" || propertie == "LastModified" || propertie == "CreationDate" || propertie == "InstallDate") {

                                value = moment(value, "YYYYMMDDHHmmSSSS").toDate();
                            }
                            fileInfos[propertie] = value
                        }
                    }
                    fileInfos["Author"] = author
                    caller.callback(fileInfos)
                },
                // Error callback.
                function () {

                }, { "callback": caller.callback, "author": author });
        },
        // Error callback.
        function () {

        }, { "path": path, "callback": callback });
}