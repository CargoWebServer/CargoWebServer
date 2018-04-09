/**
 * This is the controller file for the dashboard. It reads data from the data.js
 * file and prepares the chart objects.
 *
 * The event handlers for events such as "realTimeUpdateComplete" are written
 * here.
 * 
 * More documentation on this can be found at 
 * http://docs.fusioncharts.com/FusionCharts.html
 */
FusionCharts.ready(function() {

    /**
     * Utility function to get charts of given type
     * Returns an array of given chart type
     * @param {string} chartType - The alias of the chart
     */
    function getAllChartsByType(chartType) {
        if (chartType) {
            var allCharts = FusionCharts.items,
                chart, hLinearGauge = [],
                chartId;
            for (chartId in allCharts) {
                if (allCharts.hasOwnProperty(chartId)) {
                    chart = allCharts[chartId];
                    if (chart.chartType() == chartType) {
                        hLinearGauge.push(chart);
                    }
                }
            }
            return hLinearGauge;
        } else {
            return [];
        }
    }

    /** 
     * Updates the new value in overall rating gauge and in the caption
     * @params {string} newValue - New value for the overall rating gauge
     */
    function updateOverallRating(newValue) {
        var overallRatingGauge = FusionCharts.items['overallRating'],
            overallRatingDom = document.getElementById('overallRatingValue'),
            fixedPointNewValue = (+newValue).toFixed(2);

        // first argument is the index of the pointer
        overallRatingGauge.setData("1", newValue);

        // update overall rating value html
        overallRatingDom.textContent = fixedPointNewValue;
    }

    /** 
     *
     * Handler for the realtime update complete.
     * Get all linear gauges in page. Loop over them to calculate the sum of
     * ratings given.
     * Calculate average rating
     * Call "updateOverallRating"
     * @params {object} event - realTimeUpdateComplete event
     * @params {object} eventData - extra data passed to the handler
     */
    function updateCompleleHandler(event, eventData) {
        var hlineargauges = getAllChartsByType("hlineargauge"),
            totalRating = 0,
            averageRating = 0;

        // calculate the sum of all ratings
        hlineargauges.forEach(function(chart) {
            var realtimeJSONData = chart.getDataJSON(),
                values = realtimeJSONData["values"];
            totalRating += (+values[0]);
        });

        // calculate the average rating
        averageRating = totalRating / (hlineargauges.length);

        // update the overall rating angular gauge
        updateOverallRating(averageRating);
    }

    /**
     * ratingMeterData is defined in data.js
     * Loop over the ratingAreas array to create hlineargauge for each
     * rating criterion
     */
    ratingMeterData.ratingAreas.forEach(function(ratingArea) {
        ratingArea.type = 'hlineargauge';
        ratingArea.width = 260;
        ratingArea.height = 77;
        ratingArea.containerBackgroundOpacity = 0;
        ratingArea.renderAt = ratingArea.id + "Container";
        ratingArea.events = {
            "realTimeUpdateComplete": updateCompleleHandler
        }
        FusionCharts.render(ratingArea);
    });

    /**
     * ratingMeterData is defined in data.js
     * Loop over the overallRatings array to create angulargauge for each
     * overall rating.
     */
    ratingMeterData.overallRatings.forEach(function(overallRating) {
        overallRating.type = 'angulargauge';
        overallRating.containerBackgroundOpacity = 0;
        overallRating.height = 352;
        overallRating.width = 352;
        overallRating.renderAt = overallRating.id + "Container";
        FusionCharts.render(overallRating);
    });

});
