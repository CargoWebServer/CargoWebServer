/**
 * This file creates a global variable called "ratingMeterData"
 * It is an object the groups the data based on the type of charts we will
 * be using.
 * 
 * ratingMeterData = {
 *  ratingAreas: [{
 *      // actual data goes here
 *  }.{
 *      // actual data goes here
 *  }],
 *  overallRatings: [{
 *      // actual data goes here
 *  }, {
 *       // actual data goes here
 *  }]
 * }
 */
var ratingMeterData = {
    "ratingAreas": [{
        "id": "productQuality",
        "dataSource": {
            "chart": {
                "caption": "Product Quality",
                "theme": "rating-meter"
            },
            "pointers": {
                "pointer": [{
                    "value": "0"
                }]
            },
            "annotations": {
        "groups": [
            {
                "id": "Grp2",
                "showbelow": "1",
                "items": [
                    {
                        "type": "rectangle",
                        "x": "$gaugeStartX - 10",
                        "y": "$gaugeStartY - 3",
                        "tox": "$gaugeEndX + 10",
                        "toy": "$gaugeEndY + 3",
                        "radius": "10",
                        "fillcolor": "dd0017,e11800,ff7800,ffbb00,eed800,99CF1B"
                    },
                    {
                        "type": "rectangle",
                        "x": "$gaugeStartX + 2 - 10",
                        "y": "$gaugeStartY - 1",
                        "tox": "$gaugeEndX - 2 + 10",
                        "toy": "$gaugeEndY + 1",
                        "radius": "7",
                        "fillalpha": "0",
                        "showborder": "1",
                        "bordercolor": "#f8f8f8",
                        "borderthickness": "3"
                    }
                ]
            }
        ]
    }
        }
    }, {
        "id": "features",
        "dataSource": {
            "chart": {
                "caption": "Features",
                "theme": "rating-meter"
            },
            "pointers": {
                "pointer": [{
                    "value": "0"
                }]
            },
            "annotations": {
        "groups": [
            {
                "id": "Grp2",
                "showbelow": "1",
                "items": [
                    {
                        "type": "rectangle",
                        "x": "$gaugeStartX - 10",
                        "y": "$gaugeStartY - 3",
                        "tox": "$gaugeEndX + 10",
                        "toy": "$gaugeEndY + 3",
                        "radius": "10",
                        "fillcolor": "dd0017,e11800,ff7800,ffbb00,eed800,99CF1B"
                    },
                    {
                        "type": "rectangle",
                        "x": "$gaugeStartX + 2 - 10",
                        "y": "$gaugeStartY - 1",
                        "tox": "$gaugeEndX - 2 + 10",
                        "toy": "$gaugeEndY + 1",
                        "radius": "7",
                        "fillalpha": "0",
                        "showborder": "1",
                        "bordercolor": "#f8f8f8",
                        "borderthickness": "3"
                    }
                ]
            }
        ]
    }
        }
    }, {
        "id": "documentation",
        "dataSource": {
            "chart": {
                "caption": "Documentation",
                "theme": "rating-meter"
            },
            "pointers": {
                "pointer": [{
                    "value": "0"
                }]
            },
            "annotations": {
        "groups": [
            {
                "id": "Grp2",
                "showbelow": "1",
                "items": [
                    {
                        "type": "rectangle",
                        "x": "$gaugeStartX - 10",
                        "y": "$gaugeStartY - 3",
                        "tox": "$gaugeEndX + 10",
                        "toy": "$gaugeEndY + 3",
                        "radius": "10",
                        "fillcolor": "dd0017,e11800,ff7800,ffbb00,eed800,99CF1B"
                    },
                    {
                        "type": "rectangle",
                        "x": "$gaugeStartX + 2 - 10",
                        "y": "$gaugeStartY - 1",
                        "tox": "$gaugeEndX - 2 + 10",
                        "toy": "$gaugeEndY + 1",
                        "radius": "7",
                        "fillalpha": "0",
                        "showborder": "1",
                        "bordercolor": "#f8f8f8",
                        "borderthickness": "3"
                    }
                ]
            }
        ]
    }
        }
    }, {
        "id": "support",
        "dataSource": {
            "chart": {
                "caption": "Support",
                "theme": "rating-meter"
            },
            "pointers": {
                "pointer": [{
                    "value": "0"
                }]
            },
           "annotations": {
        "groups": [
            {
                "id": "Grp2",
                "showbelow": "1",
                "items": [
                    {
                        "type": "rectangle",
                        "x": "$gaugeStartX - 10",
                        "y": "$gaugeStartY - 3",
                        "tox": "$gaugeEndX + 10",
                        "toy": "$gaugeEndY + 3",
                        "radius": "10",
                        "fillcolor": "dd0017,e11800,ff7800,ffbb00,eed800,99CF1B"
                    },
                    {
                        "type": "rectangle",
                        "x": "$gaugeStartX + 2 - 10",
                        "y": "$gaugeStartY - 1",
                        "tox": "$gaugeEndX - 2 + 10",
                        "toy": "$gaugeEndY + 1",
                        "radius": "7",
                        "fillalpha": "0",
                        "showborder": "1",
                        "bordercolor": "#f8f8f8",
                        "borderthickness": "3"
                    }
                ]
            }
        ]
    }
        }
    }],
    "overallRatings": [{
        "id": "overallRating",
        "dataSource": {
            "chart": {
                "theme": "rating-meter"
            },
            "dials": {
                "dial": [{
                    "bgColor": "#04476C",
                    "borderColor": "#04476C",
                    "borderThickness": "0",
                    "rearExtension": 10,
                    "baseWidth": 10,
                    "radius": 120,
                    "showValue": "0"
                }]
            },
            "annotations": {
                "groups": [{
                    "id": "dialStrip",
                    "items": [{
                        "id": "redOrange",
                        "type": "arc",
                        "x": "176",
                        "y": "186",
                        "radius": 130,
                        "innerRadius": 125,
                        "startAngle": 75,
                        "endAngle": 245,
                        "fillColor": "#dd0017,#FFBB00",
                        "fillAngle": "45",
                        "fillPattern": "linear"
                    }, {
                        "id": "orangeGreen",
                        "type": "arc",
                        "x": "176",
                        "y": "186",
                        "radius": 130,
                        "innerRadius": 125,
                        "startAngle": "-65",
                        "endAngle": "76",
                        "fillColor": "#99CF1B,#FFBB00",
                        "fillAngle": "90",
                        "fillPattern": "linear"
                    }]
                }]
            }
        }
    }]
}