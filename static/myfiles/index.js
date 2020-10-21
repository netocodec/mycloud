document.addEventListener('DOMContentLoaded', function () {
    var currentDir = "/";
    var columnFilter = ['isDir', 'name', 'size'];
    var changeDir = function (dir) {
        var fileExplorerElem = document.getElementById('fileExplorer');

        if (dir !== undefined) {
            currentDir += dir;
        }

        global.makeRequest('/api/fshared/get/files?fname=' + currentDir, 'GET', undefined, function (xhr, data_json) {
            fileExplorerElem.innerHTML = '';

            var currentList = JSON.parse(data_json.folder_content);
            var fileTable = document.getElementById('fileExplorerTable');
            if (currentList.length === 0) {
                var newLineElement = document.createElement('tr');
                var newColumnElement = document.createElement('td');

                document.getElementById('fileExplorerTable').classList.add('centered');
                newColumnElement.innerHTML = '<b>There are no files to show!</b>';
                newColumnElement.setAttribute('colspan', '4');

                newLineElement.appendChild(newColumnElement);
                fileExplorerElem.appendChild(newLineElement);
            } else {
                if (fileTable.contains('centered')) {
                    fileTable.classList.add('centered');
                }

                for (var n = 0; n < currentList.length; n++) {
                    var item = currentList[n];
                    var newLineElement = document.createElement('tr');

                    for (var c = 0; c < columnFilter.length; c++) {
                        var cItem = columnFilter[c];
                        var newColumnElement = document.createElement('td');
                        var result = item[cItem];

                        if (c === 0) {
                            var fStatus = (result ? 'folder' : 'description');
                            result = '<i class="material-icons">' + fStatus + '</i>';
                        } else if (c === 2) {
                            result += ' mb';
                        }
                        newColumnElement.innerHTML = result;

                        newLineElement.appendChild(newColumnElement);
                    }

                    fileExplorerElem.appendChild(newLineElement);
                }
            }
        }, function (xhr, data_json) {
            if (xhr.status === 406) {
                M.toast({ html: '<i class="material-icons">error</i>&nbsp;' + data_json.message, classes: 'rounded' });
            } else {
                M.toast({ html: '<i class="material-icons">error</i>&nbsp; Cannot process this login at this time, try later!', classes: 'rounded' });
            }
        });
    };

    changeDir();
});
