document.addEventListener('DOMContentLoaded', function () {
    const openMenu = document.querySelector('.open-menu');
    const closeMenu = document.querySelector('.close-menu');
    const contextToggle = document.querySelector('.context-toggle');
    const menu = document.querySelector('div[role="dialog"]');
    const contextMenu = document.querySelector('.context-menu');
    const main = document.querySelector('main');
    const prov = document.querySelector('.provider');
    const provSub = document.querySelector('.provider button');
    const tok = document.querySelector('.token');
    const tokSub = document.querySelector('.token button');
    const conf = document.querySelector('.confirm');
    const confRes = document.querySelector('.confirm button[type="reset"]');
    const confSub = document.querySelector('.confirm button[type="submit"]');
    const labels = document.querySelectorAll('label:not([aria-disabled])');
    const provision = document.querySelector('.provision');
    const proProg = document.querySelector('.provision .progress');
    const proSucc = document.querySelector('.provision .success');

    const startOver = () => {
        prov.setAttribute('aria-checked', 'false');
        prov.setAttribute('aria-busy', 'false');
        tok.setAttribute('aria-checked', 'false');
        conf.setAttribute('aria-checked', 'false');
        // reset the labels
        labels.forEach(function (label) {
            label.classList.remove('!bg-blue-700');
            label.classList.add('bg-gray-800');
        });
    }

    openMenu.addEventListener('click', function () {
        menu.classList.remove('hidden');
    });

    closeMenu.addEventListener('click', function () {
        menu.classList.add('hidden');
    });

    contextToggle.addEventListener('click', function () {
        contextMenu.classList.toggle('hidden');
    });

    labels.forEach(function (label) {
        label.addEventListener('click', function () {
            label.classList.remove('bg-gray-800');
            label.classList.add('!bg-blue-700');
            // remove the class from all other labels
            labels.forEach(function (l) {
                if (l !== label) {
                    l.classList.remove('!bg-blue-700');
                    l.classList.add('bg-gray-800');
                }
            });
        });
    });

    provSub.addEventListener('click', function () {
        prov.setAttribute('aria-busy', 'true');
        tok.setAttribute('aria-checked', 'false');
    });

    tokSub.addEventListener('click', function () {
        prov.setAttribute('aria-checked', 'true');
        tok.setAttribute('aria-checked', 'true');
        conf.setAttribute('aria-checked', 'false');
    });

    confRes.addEventListener('click', function () {
        startOver();
    });

    confSub.addEventListener('click', function (event) {
        event.preventDefault();
        // get provider and token values
        const platform = prov.querySelector('input:checked').value || '';
        const apiKey = tok.querySelector('input').value || '';
        if (platform && apiKey) {
            const data = new Object();
            data.platform = platform;
            data.apiKey = apiKey;
            main.classList.add('hidden');
            provision.classList.remove('hidden');
            provision.classList.add('grid');
            conf.setAttribute('aria-checked', 'true');
            // send the data to the server
            fetch('/provision', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify(data)
            })
            .then(response => response.text())
            .then(data => {
                // if data contains error
                if (data.includes('Error')) {
                    console.log('Error:', data);
                    alert("An error occurred while trying to provision the service. Please try again later.");
                    provision.classList.add('hidden');
                    main.classList.remove('hidden');
                    startOver();
                } else {
                    console.log('Success:', data);
                    proProg.classList.remove('grid');
                    proProg.classList.add('hidden');
                    proSucc.classList.remove('hidden');
                    progSucc.classList.add('grid');
                    proSucc.querySelector('textarea').value = data;
                }
            })
            .catch(error => {
                console.error('Error:', error);
            });
        }

    });

});