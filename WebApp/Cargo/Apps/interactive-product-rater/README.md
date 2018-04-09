Rating meter is a dashboard to demonstrate the editable features of FusionCharts
Widgets

Charts used are
	- Real-time Angular (angulargauge)
	- Real-time Horizontal Linear (hlineargauge)

The dashboard has the following folder structure
rating-meter
	- index.html
	- readme.md
	- css (contains css files)
	- js (contains js files)
	- fusioncharts (contains FusionCharts library files)

# dashboard.js
This is the controller file which reads data from the object inside 
data.js and prepares the chart object.

# data.js
This file contains the input data object for all charts that are to be drawn. 
This data object is used by the dashboard.js for each chart while rendering it.

# fusioncharts.theme.rating-meter.js
This is the theme file which is used to define the cosmetic properties of chart
such as chart-padding, chart-margin etc.