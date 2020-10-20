document.addEventListener('DOMContentLoaded', function () {
    var ctx = document.getElementById('myCloudChart').getContext('2d');
    var color = {
        blue: "rgb(54, 162, 235)",
        green: "rgb(75, 192, 192)",
        grey: "rgb(201, 203, 207)",
        orange: "rgb(255, 159, 64)",
        purple: "rgb(153, 102, 255)",
        red: "rgb(255, 99, 132)",
        yellow: "rgb(255, 205, 86)"
    };
    var myChart = new Chart(ctx, {
        type: 'pie',
        data: {
            datasets: [{
                data: [
                    25,
                    25,
                    25,
                    25,
                    25,
                ],
                backgroundColor: [
                    color.red,
                    color.orange,
                    color.yellow,
                    color.green,
                    color.blue,
                ],
                label: 'Dataset 1'
            }],
            labels: [
                'Red',
                'Orange',
                'Yellow',
                'Green',
                'Blue'
            ]
        },
        options: {
            responsive: true
        }
    });
});