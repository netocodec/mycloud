document.addEventListener('DOMContentLoaded', function () {
    console.log("HELLO");


    document.getElementById('submit_login').addEventListener('click', function () {
        global.makeRequest('/api/login');
    });
});