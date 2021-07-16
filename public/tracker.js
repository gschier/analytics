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
    params.push({ name: 'xy', value: screensize() });
    const oReq = new XMLHttpRequest();
    oReq.mode = 'no-cors';
    oReq.open('GET', `${scriptOrigin()}/api${path}?${(params.map(v => `${v.name}=${encodeURIComponent(v.value)}`).join('&'))}`);
    oReq.send();
  }

  let _script = null;

  function script() {
    if (!_script) {
      _script = document.querySelector('script[data-website]');
    }
    return _script;
  }

  function screensize() {
    return `${Math.round(window.innerWidth / 100) * 100}x${Math.round(window.innerHeight / 100) * 100}`;
  }

  function website() {
    return script().getAttribute('data-website');
  }

  function scriptOrigin() {
    return script().getAttribute('src').match(/https:\/\/[^/]*/)[0];
  }

  page();
})();
