document.addEventListener('DOMContentLoaded', function () {
    var modalElems = document.querySelectorAll('.modal-myfiles');
    var inputValueCounter = document.querySelectorAll('.counterInput');
    var currentDir = "/";
    var columnFilter = ['IS_DIR', 'FName', 'FSize'];
    var CHUNK_UPLOAD_SIZE = 800;
    var bytesToSize = function (bytes) {
        var sizes = ['Bytes', 'KB', 'MB', 'GB', 'TB'];
        if (bytes == 0) return '0 Byte';
        var i = parseInt(Math.floor(Math.log(bytes) / Math.log(1024)));
        return Math.round(bytes / Math.pow(1024, i), 2) + ' ' + sizes[i];
    };
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
                if (fileTable.classList.contains('centered')) {
                    fileTable.classList.add('centered');
                }

                for (var n = 0; n < currentList.length; n++) {
                    var item = currentList[n];
                    var newLineElement = document.createElement('tr');

                    newLineElement.setAttribute('data-path', item.FName);
                    newLineElement.classList.add('hand');
                    newLineElement.addEventListener('mouseover', function () {
                        this.classList.add('light-blue');
                        this.classList.add('darken-1');
                        this.classList.add('white-text');
                    });

                    newLineElement.addEventListener('mouseout', function () {
                        this.classList.remove('light-blue');
                        this.classList.remove('darken-1');
                        this.classList.remove('white-text');
                    });

                    newLineElement.addEventListener('click', function (evt) {
                        evt.preventDefault();
                        changeDir(this.getAttribute('data-path'));
                    });

                    for (var c = 0; c < columnFilter.length; c++) {
                        var cItem = columnFilter[c];
                        var newColumnElement = document.createElement('td');
                        var result = item[cItem];

                        if (c === 0) {
                            cItem = columnFilter[c + 1];
                            result = item[cItem];

                            var fStatus = (result.indexOf('.') === -1 ? 'folder' : 'description');
                            result = '<i class="material-icons">' + fStatus + '</i>';
                        } else if (c === 2) {
                            result = bytesToSize(result);
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
    var FileUpload = function (file) {
        var reader = new FileReader();
        var xhr = new XMLHttpRequest();
        var fileName = file.name;
        var lastChunk = 0;
        var currentChunk = 0;
        var result = {
            currentd: currentDir,
            chunk: ''
        };

        xhr.upload.addEventListener("progress", function (e) {
            if (e.lengthComputable) {
                var percentage = Math.round((e.loaded * 100) / e.total);
                global.add_notification('upload_f_' + fileName, 'Upload file ' + fileName + ' with success!', percentage);
            }
        }, false);

        xhr.upload.addEventListener("load", function (e) {
            global.add_notification('upload_f_' + fileName, 'Upload file ' + fileName + ' with success!', 100);
        }, false);

        xhr.open('POST', '/api/fshared/upload/' + fileName + '/' + lastChunk);

        reader.onload = function (evt) {
            var data = evt.target.result;
            var totalData = data.length;

            if (totalData <= CHUNK_UPLOAD_SIZE) {
                lastChunk = 1;
                result.chunk = data;
                xhr.send(JSON.stringify(result));
            }
        };

        reader.readAsBinaryString(file);
    };

    var dir_name = document.getElementById('dir_name');
    var create_dir_btn = document.getElementById('createDirBtn');
    var new_dir_form = document.getElementById('newDirForm');
    dir_name.addEventListener('input', function (event) {
        var is_valid = new_dir_form.checkValidity();

        if (!is_valid && !create_dir_btn.classList.contains('disabled')) {
            create_dir_btn.classList.add('disabled');
        } else if (is_valid && create_dir_btn.classList.contains('disabled')) {
            create_dir_btn.classList.remove('disabled');
        }
    });

    create_dir_btn.addEventListener('click', function () {
        if (new_dir_form.checkValidity()) {
            var dir_name = document.getElementById('dir_name').value;

            global.makeRequest('/api/fshared/mk/' + dir_name, 'POST', {
                currentd: currentDir
            }, function (xhr, data_json) {
                if (xhr.status === 200) {
                    M.toast({ html: '<i class="material-icons">done_outline</i>&nbsp;' + data_json.message, classes: 'rounded blue' });
                    document.getElementById('closeCreateDirBtn').click();
                    changeDir();
                }
            }, function (xhr, data_json) {
                if (xhr.status === 406) {
                    M.toast({ html: '<i class="material-icons">error</i>&nbsp;' + data_json.message, classes: 'rounded' });
                } else {
                    M.toast({ html: '<i class="material-icons">error</i>&nbsp; Cannot create this directory at this time, try later!', classes: 'rounded' });
                }
            });
        } else {
            M.toast({ html: '<i class="material-icons">error</i>&nbsp; Directory name is invalid, try another one!', classes: 'rounded red' });
        }
    });

    changeDir();
    M.Modal.init(modalElems, {
        onOpenStart: function (evt) {
            var id = evt.getAttribute('id');

            if (id === 'newDirModal') {
                if (!create_dir_btn.classList.contains('disabled')) {
                    create_dir_btn.classList.add('disabled');
                }

                document.getElementById('dir_name').focus();
            }
        },
        onCloseEnd: function (evt) {
            var id = evt.getAttribute('id');

            switch (id) {
                case 'newDirModal':
                    document.getElementById('newDirForm').reset();
                    break;

                case 'uploadFileModal':
                    document.getElementById('uploadFileForm').reset();
                    break;
            }
        }
    });

    if (inputValueCounter.length !== 0) {
        new M.CharacterCounter(inputValueCounter[0], {});
    }


    // File Upload
    var uploadFile = document.getElementById('uploadNewFiles');
    var dragDropUpload = document.getElementById('dragDropUpload');
    uploadFile.addEventListener('change', function () {
        var files = this.files;

        document.getElementById('closeUploadBtn').click();
        M.toast({ html: '<i class="material-icons">warning</i>&nbsp;Files are uploading right now! See the notification cloud!', classes: 'rounded' });

        for (var c = 0; c < files.length; c++) {
            var file = files[c];

            new FileUpload(file);
        }
    });

    dragDropUpload.addEventListener('dragenter', function (e) {
        e.stopPropagation();
        e.preventDefault();

        dragDropUpload.classList.remove('darken-1');
        dragDropUpload.classList.add('darken-3');
    });

    dragDropUpload.addEventListener('dragleave', function (e) {
        e.stopPropagation();
        e.preventDefault();

        dragDropUpload.classList.add('darken-1');
        dragDropUpload.classList.remove('darken-3');
    });

    dragDropUpload.addEventListener('dragover', function (e) {
        e.stopPropagation();
        e.preventDefault();
    });

    dragDropUpload.addEventListener('drop', function (e) {
        e.stopPropagation();
        e.preventDefault();

        dragDropUpload.classList.add('darken-1');
        dragDropUpload.classList.remove('darken-3');

        var dt = e.dataTransfer;
        var files = dt.files;

        document.getElementById('closeUploadBtn').click();
        M.toast({ html: '<i class="material-icons">warning</i>&nbsp;Files are uploading right now! See the notification cloud!', classes: 'rounded' });

        for (var c = 0; c < files.length; c++) {
            var file = files[c];

            new FileUpload(file);
        }
    });
});
