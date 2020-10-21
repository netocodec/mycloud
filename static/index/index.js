document.addEventListener('DOMContentLoaded', function () {
    var error = document.getElementById('errorPage').value;
    document.getElementById('username').focus();

    if (error === 'ACCOUNT_NOT_AUTH') {
        M.toast({ html: '<i class="material-icons">error</i>&nbsp; Account session has been expired!', classes: 'rounded red' });
    }

    document.getElementById('submit_login').addEventListener('click', function () {
        global.makeRequest('/api/login', 'POST', {
            user: document.getElementById('username').value,
            pass: document.getElementById('password').value
        }, function (xhr, data_json) {
            window.location = '/member/dashboard';
        }, function (xhr, data_json) {
            if (xhr.status === 406) {
                M.toast({ html: '<i class="material-icons">error</i>&nbsp;' + data_json.message, classes: 'rounded' });
            } else {
                M.toast({ html: '<i class="material-icons">error</i>&nbsp; Cannot process this login at this time, try later!', classes: 'rounded' });
            }
        });
    });
});