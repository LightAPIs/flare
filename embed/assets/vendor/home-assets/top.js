(() => {
  function scrollTop() {
    const doc = document.documentElement;
    doc.scrollTo({
      top: 0,
      behavior: 'smooth',
    });
  }

  const topBtn = document.getElementById('btn-scroll-top');
  if (topBtn) {
    topBtn.addEventListener('click', ev => {
      ev.stopPropagation();
      ev.preventDefault();
      scrollTop();
    });
    topBtn.addEventListener(
      'touchstart',
      ev => {
        ev.stopPropagation();
        ev.preventDefault();
        scrollTop();
      },
      { passive: false }
    );
  }
})();
