document.addEventListener('readystatechange', function () {
    console.log('CHANGE_STATE', document.readyState);
    if (document.readyState !== 'complete') {
        global.startLoading();
    } else {
        global.hideLoading();
    }
});

window.addEventListener('beforeunload', function (evt) {
    global.startLoading();
});
