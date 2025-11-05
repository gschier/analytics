(function() {
  function event(name, attributes) {
    const params = new URLSearchParams();
    params.set('e', name);
    if (attributes) params.set('a', JSON.stringify(attributes));
    send('/e', params);
  }

  function page() {
    const { path, protocol, host } = window.location;
    if (path === sessionStorage.lastPathName || !host) return;
    sessionStorage.lastPathName = path;

    const params = new URLSearchParams();
    params.set('h', `${protocol}//${host}`);
    params.set('p', path);

    const urlParams = new URLSearchParams(window.location.search);
    const referrer = urlParams.get('ref') || document.referrer;
    params.set('r', referrer);

    send('/p', params);
  }

  function send(path, params) {
    const id = website();
    params.set('id', id);
    params.set('u', uid());
    params.set('tz', Intl.DateTimeFormat().resolvedOptions().timeZone);
    params.set('xy', screensize());

    for (const [key, value] of params.entries()) {
      if (!value) params.delete(key);
    }

    const url = `${scriptOrigin()}/t${path}?${params.toString()}`;

    if (
      localStorage.disableAnalytics === 'true' ||
      window.location.hostname === 'localhost' ||
      id === ''
    ) {
      console.log('Analytics disabled', params.toString());
      return;
    }

    fetch(url, { mode: 'no-cors' }).catch(console.error);
  }

  let _script = null;

  function script() {
    if (!_script) {
      _script = document.querySelector('script[data-website]');
    }
    return _script;
  }

  function screensize() {
    const w = window.screen.width;
    const h = window.screen.height;
    return `${Math.round(w / 100) * 100}x${Math.round(h / 100) * 100}`;
  }

  function website() {
    return script().getAttribute('data-website') || '';
  }

  function uid() {
    return script().getAttribute('data-uid') || '';
  }

  function scriptOrigin() {
    const src = script().getAttribute('src') || '';
    return src.match(/https:\/\/[^/]*/)[0];
  }

  page();

  window.trackEvent = event;
})();
