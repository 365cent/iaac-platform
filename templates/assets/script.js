document.addEventListener('DOMContentLoaded', function() {
    const openMenu = document.querySelector('.open-menu');
    const closeMenu = document.querySelector('.close-menu');
    const contextToggle = document.querySelector('.context-toggle');
    const menu = document.querySelector('div[role="dialog"]');
    const contextMenu = document.querySelector('.context-menu');

    const labels = document.querySelectorAll('label');

    openMenu.addEventListener('click', function() {
        menu.classList.remove('hidden');
    });

    closeMenu.addEventListener('click', function() {
        menu.classList.add('hidden');
    });

    contextToggle.addEventListener('click', function() {
        contextMenu.classList.toggle('hidden');
    });

    labels.forEach(function(label) {
        label.addEventListener('click', function() {
            label.classList.remove('bg-gray-800');
            label.classList.add('!bg-blue-700');
            // remove the class from all other labels
            labels.forEach(function(l) {
                if (l !== label) {
                    l.classList.remove('!bg-blue-700');
                    l.classList.add('bg-gray-800');
                }
            });
        });
    });
});