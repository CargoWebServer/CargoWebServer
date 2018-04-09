/**
 * This file create a theme file for the dashboard. The cosmetics for the
 * different charts are defined here.
 *
 * More information on theming can be found at
 * http://docs.fusioncharts.com/tutorial-configuring-your-chart-theme-manager.html
 */
FusionCharts.register('theme', {
    name: 'rating-meter',
    theme: {
        hlineargauge: {
            chart: {
                "captionAlignment": "left",
                "bgColor": "#ff0000",
                "bgAlpha": 0,
                "showBorder": 0,
                "showValue": "0",
                "lowerLimit": "0",
                "upperLimit": "5",
                "editMode": 1,
                "pointerOnTop": "0",
                "useRoundEdges": "0",
                "showTickMarks": 1,
                "trendValueDistance": 100,
                "ticksBelowGauge": "0",
                "placeTicksInside": "0",
                "showGaugeLabels": 0,
                "pointerBgColor": "#000000",
                "showBorderOnHover": "1",
                "pointerBorderHoverColor": "#333333",
                "pointerBorderHoverThickness": "1",
                "majorTMHeight": 5,
                "minorTMThickness": 0,
                "connectorThickness": 0,
                "showGaugeBorder": 0,
                "gaugeBorderThickness": 2,
                "gaugeBorderColor": "#000000",
                "showShadow": "0"
            }
        },
        angulargauge: {
            chart: {
                "bgAlpha": 0,
                "showBorder": 0,
                "showTickMarks": 1,
                "placeTicksInside": 0,
                "lowerLimit": "0",
                "upperLimit": "5",
                "gaugeStartAngle": 245,
                "gaugeEndAngle": -65,
                "gaugeOuterRadius": 150,
                "gaugeInnerRadius": 149,
                "gaugeFillMix": "#ff0000",
                "showGaugeBorder": 1,
                "gaugeBorderThickness": 5,
                "pivotRadius": 4,
                "pivotFillColor": "#ffffff",
                "pivotFillMix": "{light}",
                "showPivotBorder": "0",
                "connectorThickness": 0,
                "minorTMHeight": "0",
                "minorTMThickness": "0",
                "majorTMHeight": "10",
                "showShadow": "0",
                "showToolTip": "0",
                "animationDuration": "0.5"
            }
        }
    }
});