Dim xlPath, csvPath, xlApp, fso, xlWb

Set fso = CreateObject("Scripting.FileSystemObject")
Set xlApp = CreateObject("Excel.Application")
curDir = fso.GetParentFolderName(WScript.ScriptFullName)

' The the excel file 
xlPath = WScript.Arguments(0)

' The csv file
csvPath = WScript.Arguments(1)

set xlWb = xlApp.Workbooks.Open(csvPath, Format:=4)
xlApp.DisplayAlerts = False
call xlWb.saveas(xlPath, FileFormat:=XlFileFormat.xlOpenXMLWorkbook)
xlApp.Quit

' Return the path of the file.
Dim console
Set console = WScript.StdOut
console.WriteLine xlPath