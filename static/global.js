var global = {};

global.startLoading = function () {
    document.getElementById('loading').classList.remove('hide');
    document.getElementById('loading_overlay').classList.remove('hide');
};

global.hideLoading = function () {
    document.getElementById('loading').classList.add('hide');
    document.getElementById('loading_overlay').classList.add('hide');
};

global.makeRequest = function (url, method, params, success_callback, failed_callback) {
    var xhr = new XMLHttpRequest();

    xhr.onreadystatechange = function () {
        var data_json = null;

        if (xhr.readyState === 4) {
            try {
                data_json = JSON.parse(xhr.responseText);
            } catch (ex) {
                data_json = xhr.responseText;
            }
        }

        if (xhr.readyState === 4 && xhr.status === 200) {
            success_callback(xhr, data_json);
        } else if (xhr.readyState === 4 && xhr.status !== 200) {
            failed_callback(xhr, data_json);
        }
    };

    if (failed_callback !== undefined) {
        xhr.ontimeout = function () {
            failed_callback();
        };
    }

    xhr.onloadstart = global.startLoading;
    xhr.onloadend = global.hideLoading;
    xhr.open(method, url, true);
    xhr.timeout = 30000;
    xhr.setRequestHeader('X-Requested-With', 'XMLHttpRequest');
    xhr.setRequestHeader('Content-Type', 'application/json;charset=UTF-8');

    var token = localStorage.getItem('mc_tok');
    if (token !== null) {
        xhr.setRequestHeader('Authorization', 'Bearer ' + token);
    }

    if (method === 'POST')
        xhr.send(JSON.stringify(params));
    else
        xhr.send();
};