document.addEventListener('readystatechange', function () {
    if (document.readyState !== 'complete') {
        global.startLoading();
    } else {
        global.hideLoading();
    }
});

window.addEventListener('beforeunload', function (evt) {
    global.startLoading();
});
