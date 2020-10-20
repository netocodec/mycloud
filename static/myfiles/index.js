document.addEventListener('DOMContentLoaded', function () {
    var currentDir = "/";
    var columnFilter = ['isDir', 'name', 'size'];
    var currentList = [
        {
            isDir: 0,
            name: "test.txt",
            size: 2
        },
        {
            isDir: 1,
            name: "sIsNewFolder",
            size: 50
        }
    ];
    var changeDir = function (dir) {
        var fileExplorerElem = document.getElementById('fileExplorer');

        fileExplorerElem.innerHTML = '';
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
    };

    changeDir(currentDir);
});
