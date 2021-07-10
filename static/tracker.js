async function track(eventName) {
  const script = document.querySelector('script[data-website]');
  const website = script.getAttribute('data-website');
  await fetch(`//${window.location.host}/event?name=${eventName}&website=${website}`);
}

track('dummy').catch(err => console.error('Failed to track event', err));
