google.charts.load('current', {packages: ['corechart']});
google.charts.setOnLoadCallback(drawChart);

function drawChart() {
	// hard coded, we could also make this a template and import from the json (or some other technique)
	var points = [
		{
			x: 63,
			y: -72,
		},
		{
			x: -94,
			y: 89,
		},
		{
			x: -30,
			y: -38,
		},
		{
			x: -43,
			y: -65,
		},
		{
			x: 88,
			y: -74,
		},
		{
			x: 73,
			y: 89,
		},
		{
			x: 39,
			y: -78,
		},
		{
			x: 51,
			y: 8,
		},
		{
			x: -92,
			y: -93,
		},
		{
			x: -77,
			y: -45,
		},
		{
			x: -41,
			y: 29,
		},
		{
			x: 54,
			y: -94,
		},
		{
			x: 43,
			y: -50,
		},
		{
			x: 83,
			y: 66,
		},
		{
			x: 79,
			y: 39,
		},
		{
			x: 7,
			y: -44,
		},
		{
			x: 14,
			y: 50,
		},
		{
			x: -29,
			y: -99,
		},
		{
			x: 94,
			y: -60,
		},
		{
			x: 78,
			y: 8,
		},
		{
			x: -13,
			y: -29,
		},
		{
			x: -61,
			y: -45,
		},
		{
			x: 95,
			y: -14,
		},
		{
			x: -74,
			y: -77,
		},
		{
			x: -3,
			y: -76,
		},
		{
			x: -9,
			y: -12,
		},
		{
			x: 54,
			y: -33,
		},
		{
			x: -89,
			y: 86,
		},
		{
			x: 17,
			y: 37,
		},
		{
			x: -69,
			y: -4,
		},
		{
			x: -80,
			y: 41,
		},
		{
			x: -25,
			y: 60,
		},
		{
			x: 58,
			y: -8,
		},
		{
			x: 47,
			y: -51,
		},
		{
			x: 80,
			y: -83,
		},
		{
			x: -89,
			y: 69,
		},
		{
			x: -42,
			y: 97,
		},
		{
			x: -26,
			y: -80,
		},
		{
			x: -41,
			y: -75,
		},
		{
			x: -3,
			y: -29,
		},
		{
			x: 16,
			y: 62,
		},
		{
			x: -7,
			y: -59,
		},
		{
			x: -6,
			y: -10,
		},
		{
			x: -47,
			y: 71,
		},
		{
			x: -32,
			y: 79,
		},
		{
			x: 74,
			y: 65,
		},
		{
			x: -82,
			y: 55,
		},
		{
			x: 62,
			y: -57,
		},
		{
			x: 36,
			y: 86,
		},
		{
			x: -38,
			y: -59,
		},
		{
			x: 18,
			y: -3,
		},
		{
			x: -31,
			y: 63,
		},
		{
			x: 76,
			y: 42,
		},
		{
			x: -44,
			y: 75,
		},
		{
			x: -17,
			y: 96,
		},
		{
			x: 98,
			y: -86,
		},
		{
			x: -42,
			y: -92,
		},
		{
			x: -20,
			y: 2,
		},
		{
			x: -32,
			y: -84,
		},
		{
			x: -46,
			y: 45,
		},
		{
			x: 83,
			y: -20,
		},
		{
			x: -46,
			y: 67,
		},
		{
			x: 27,
			y: 1,
		},
		{
			x: 64,
			y: 17,
		},
		{
			x: -64,
			y: -33,
		},
		{
			x: -65,
			y: -37,
		},
		{
			x: 90,
			y: 43,
		},
		{
			x: 37,
			y: -33,
		},
		{
			x: 91,
			y: 49,
		},
		{
			x: 9,
			y: 49,
		},
		{
			x: 2,
			y: -8,
		},
		{
			x: -44,
			y: -65,
		},
		{
			x: 30,
			y: 26,
		},
		{
			x: -77,
			y: 93,
		},
		{
			x: -88,
			y: -72,
		},
		{
			x: -61,
			y: 60,
		},
		{
			x: -60,
			y: 74,
		},
		{
			x: 8,
			y: 52,
		},
		{
			x: -84,
			y: -2,
		},
		{
			x: -3,
			y: 52,
		},
		{
			x: 19,
			y: 35,
		},
		{
			x: -36,
			y: 41,
		},
		{
			x: -98,
			y: 74,
		},
		{
			x: 84,
			y: -71,
		},
		{
			x: 74,
			y: 37,
		},
		{
			x: 92,
			y: -32,
		},
		{
			x: 96,
			y: 64,
		},
		{
			x: -13,
			y: -72,
		},
		{
			x: -25,
			y: 11,
		},
		{
			x: -60,
			y: 16,
		},
		{
			x: -100,
			y: 84,
		},
		{
			x: 84,
			y: -33,
		},
		{
			x: 28,
			y: 95,
		},
		{
			x: -55,
			y: 29,
		},
		{
			x: -73,
			y: 60,
		},
		{
			x: -24,
			y: 63,
		},
		{
			x: 29,
			y: 55,
		},
		{
			x: -50,
			y: -61,
		},
		{
			x: -5,
			y: 95,
		},
		{
			x: -59,
			y: 38,
		},
	];

	var input = [['X', 'Y']];
	points.forEach(function (v) {
		input.push([v.x, v.y]);
	});

	var data = google.visualization.arrayToDataTable(input);

	var options = {
		title: 'Points',
		hAxis: {title: 'X'},
		vAxis: {title: 'Y'},
		legend: 'none',
	};

	var chart = new google.visualization.ScatterChart(document.getElementById('chart_div'));

	chart.draw(data, options);
}