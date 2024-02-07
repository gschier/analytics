(function () {
  function event(name, attributes) {
    send('/e', [
      { name: 'e', value: name },
      {
        name: 'a',
        value: JSON.stringify(attributes || {}),
      },
    ]);
  }

  function page() {
    const { pathname, protocol, host } = window.location;
    if (pathname === sessionStorage.lastPathName || !host) return;
    sessionStorage.lastPathName = pathname;
    send('/p', [
      {
        name: 'h',
        value: `${protocol}//${host}`,
      },
      { name: 'p', value: pathname },
      { name: 'r', value: document.referrer },
    ]);
  }

  function send(path, params) {
    if (
      localStorage.disableAnalytics === 'true' ||
      window.location.hostname === 'localhost'
    ) {
      console.log('Analytics disabled', path, params);
      return;
    }

    params.push({ name: 'id', value: website() });
    params.push({
      name: 'tz',
      value: Intl.DateTimeFormat().resolvedOptions().timeZone,
    });
    params.push({ name: 'xy', value: screensize() });
    const qs = params
      .filter(Boolean)
      .map((v) => `${v.name}=${encodeURIComponent(v.value)}`)
      .join('&');
    const url = `${scriptOrigin()}/t${path}?${qs}`;
    fetch(url, { mode: 'no-cors' }).catch((err) => console.log('Error:', err));
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

  function scriptOrigin() {
    return script()
      .getAttribute('src')
      .match(/https:\/\/[^/]*/)[0];
  }

  page();

  window.trackEvent = event;
})();
