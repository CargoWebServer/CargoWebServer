Dim snPath, xlApp, fso, xlWb, strLine
Dim csvPath, csvFile, txtPath, txtFile

Set fso = CreateObject("Scripting.FileSystemObject")
Set xlApp = CreateObject("Excel.Application")
curDir = fso.GetParentFolderName(WScript.ScriptFullName)
csvPath = curDir & "\sn.csv"
snPath = "\\mon-filer-01\Data\Departement Commun\Num" & chr(233) & "ros de s" & chr(233) & "rie vs Bons de travail\SN.XLS"

set xlWb = xlApp.Workbooks.Open(snPath,,1)
xlApp.DisplayAlerts = False
call xlWb.saveas(csvPath, 6, 0, 0, 0, 0, 0, 0, 0, 0, 0, local)
xlApp.Quit

Dim console
Set console = WScript.StdOut
console.WriteLine csvPath















