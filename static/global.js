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

global.add_notification = function (id, title, prog) {
    var notificationNumber = document.getElementById('notificationNum');
    var notificationList = document.getElementById('notificationList');
    var dataIdElements = document.querySelectorAll('[data-id="' + id + '"]');
    var num = 0;
    var lastNumber = parseInt(notificationNumber.innerHTML);
    var isDiffNumber = false;
    var addRemoveButton = function (newNotification, id, prog) {
        var elem = document.querySelector('[data-remove-btn="' + id + '"]');
        if (!elem) {
            newNotification.innerHTML += '<a data-remove-btn="' + id + '" href="#!" onclick="this.parentElement.remove()" class="secondary-content black-text" ' + (prog === 100 ? 'style ="display:none!important;"' : '') + '><i class="material-icons">close</i></a>';
        } else {
            if (elem.hasAttribute('style')) {
                elem.removeAttribute('style');
            }
        }
    };
    var addLoading = function (newNotification, id, prog) {
        newNotification.innerHTML += '<div class="progress" ' + (prog ? 'style ="display:none!important;"' : '') + '> <div data-prog-id="' + id + '" class="determinate" style="width: 0%"></div></div>';
    };
    var setLoading = function (id, prog) {
        var elem = document.querySelector('[data-prog-id="' + id + '"]');
        if (elem) {
            if (elem.parentElement.hasAttribute('style')) {
                elem.parentElement.removeAttribute('style');
            }

            elem.setAttribute('style', 'width: ' + prog + '%;');
        }
    };

    if (lastNumber !== '0') {
        num = lastNumber;
    }

    if (dataIdElements.length === 0) {
        num++;
        isDiffNumber = (num !== lastNumber);
    } else {
        isDiffNumber = true;
    }

    if (isDiffNumber) {
        notificationNumber.classList.add('new');
        notificationNumber.classList.add('grey');
        notificationNumber.innerHTML = num;
    }

    var newNotification = document.createElement('li');
    if (dataIdElements.length !== 0) {
        newNotification = dataIdElements[0];
        newNotification.querySelector('[data-id-title="' + id + '"]').innerHTML = title;

        if (prog) {
            setLoading(id, prog);

            if (prog === 100) {
                addRemoveButton(newNotification, id, prog);
            }
        }
    } else {
        newNotification.classList.add('collection-item');
        newNotification.setAttribute('data-id', id);
        newNotification.innerHTML = '<span data-id-title="' + id + '">' + title + '</span>';
        addLoading(newNotification, id, prog);
        addRemoveButton(newNotification, id, prog);

        notificationList.appendChild(newNotification);
    }
};


document.addEventListener('DOMContentLoaded', function () {
    var notificationElems = document.querySelectorAll('.notification-modal');

    M.Modal.init(notificationElems, {
        onOpenEnd: function () {
            var notificationElem = document.getElementById('notificationNum');

            notificationElem.classList.remove('new');
            notificationElem.classList.remove('grey');
            notificationElem.innerHTML = '0';
        }
    })
});