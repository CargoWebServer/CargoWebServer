Dim xlPath, xlApp, fso, xlWb
Dim csvPath, csvFile

Set fso = CreateObject("Scripting.FileSystemObject")
Set xlApp = CreateObject("Excel.Application")
curDir = fso.GetParentFolderName(WScript.ScriptFullName)

' The the excel file 
xlPath = WScript.Arguments(0)
' The csv file
csvPath = WScript.Arguments(1)

set xlWb = xlApp.Workbooks.Open(xlPath,,1)
xlApp.DisplayAlerts = false
call xlWb.saveas(csvPath, 6, 0, 0, 0, 0, 0, 0, 0, 0, 0, local)
xlApp.Quit

' Return the path of the file.
Dim console
Set console = WScript.StdOut
console.WriteLine csvPath















