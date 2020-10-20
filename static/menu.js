document.addEventListener('DOMContentLoaded', function () {
    var elems = document.querySelectorAll('.sidenav');
    var dropdownElems = document.querySelectorAll('.dropdown-trigger');

    M.Sidenav.init(elems, {});
    M.Dropdown.init(dropdownElems, {});
});