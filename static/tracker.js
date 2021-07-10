(function () {
  function event(name) {
    send('/event', [
      { name: 'e', value: name },
    ]);
  }

  function page() {
    send('/page', [
      { name: 'h', value: `${window.location.protocol}//${window.location.host}` },
      { name: 'p', value: window.location.pathname },
    ]);
  }

  function send(path, params) {
    params.push({ name: 'id', value: website() });
    params.push({ name: 'tz', value: Intl.DateTimeFormat().resolvedOptions().timeZone });
    params.push({
      name: 'xy',
      value: `${Math.round(window.innerWidth / 100) * 100}x${Math.round(window.innerHeight / 100) * 100}`,
    });
    const qs = params.map(v => `${encodeURIComponent(v.name)}=${encodeURIComponent(v.value)}`).join('&');
    fetch(`//${window.location.host}${path}?${qs}`).catch(err => console.log('Failed to send', err));
  }

  let _website = null;

  function website() {
    if (!_website) {
      _website = document.querySelector('script[data-website]').getAttribute('data-website');
    }

    return _website;
  }

  page();
})();
