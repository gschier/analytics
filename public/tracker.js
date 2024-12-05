(function() {
  function event(name, attributes) {
    const params = new URLSearchParams();
    params.set('e', name);
    if (attributes) params.set('a', JSON.stringify(attributes));
    send('/e', params);
  }

  function page() {
    const { pathname, protocol, host } = window.location;
    if (pathname === sessionStorage.lastPathName || !host) return;
    sessionStorage.lastPathName = pathname;

    const params = new URLSearchParams();
    params.set('h', `${protocol}//${host}`);
    params.set('p', pathname);
    params.set('r', document.referrer);
    send('/p', params);
  }

  function send(path, params) {
    params.set('id', website());
    params.set('u', uid());
    params.set('tz', Intl.DateTimeFormat().resolvedOptions().timeZone);
    params.set('xy', screensize());
    const url = `${scriptOrigin()}/t${path}`;

    if (
      localStorage.disableAnalytics === 'true' ||
      window.location.hostname === 'localhost'
    ) {
      console.log('Analytics disabled', url, params.toString());
      return;
    }

    navigator.sendBeacon(url, params);
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
    return script().getAttribute('data-website');
  }

  function uid() {
    return script().getAttribute('data-uid');
  }

  function scriptOrigin() {
    return script()
      .getAttribute('src')
      .match(/https:\/\/[^/]*/)[0];
  }

  page();

  window.trackEvent = event;
})();
