(() => {
  const hide = 'search-result-hide';
  const showBtn = 'show-clear-button';
  let timer = null;
  let debounced = null;
  function debounce(func, delay, immediate) {
    window.clearTimeout(timer);

    if (immediate) {
      //* 立即执行
      if (timer === null) {
        func();
      }
    }

    timer = window.setTimeout(() => {
      func();
      timer = null;
    }, delay);

    return {
      cancel: () => {
        window.clearTimeout(timer);
        timer = null;
      },
    };
  }

  function restoreDisplay(appList, bookmarkList) {
    if (debounced) {
      debounced.cancel();
      debounced = null;
    }

    appList.forEach(app => {
      app.ele.classList.remove(hide);
    });
    bookmarkList.forEach(bm => {
      bm.ele.classList.remove(hide);
      bm.items.forEach(item => {
        item.ele.classList.remove(hide);
      });
    });
  }

  /** 搜索框 */
  const searchInput = document.getElementById('search');
  /** 清空按钮 */
  const clearBtn = document.getElementById('clear-search-content');

  if (searchInput) {
    // 获取当前已有的网址信息
    const curAppList = [];
    document.querySelectorAll('.apps-container .app-container').forEach(app => {
      curAppList.push({
        ele: app,
        url: app.querySelector('a').href,
        title: app.querySelector('.app-title').innerText.trim().toLowerCase(),
        desc: app.querySelector('.app-desc').innerText.trim().toLowerCase(),
      });
    });
    const curBookmarkList = [];
    document.querySelectorAll('.bookmark-group-container').forEach(group => {
      const items = [];
      group.querySelectorAll('.bookmark-list li').forEach(item => {
        items.push({
          ele: item,
          url: item.querySelector('a').href,
          title: item.querySelector('a > span').innerText.trim().toLowerCase(),
        });
      });

      curBookmarkList.push({
        ele: group,
        items,
      });
    });

    searchInput.addEventListener('keydown', ev => {
      if (ev.key === 'Escape') {
        restoreDisplay(curAppList, curBookmarkList);
        searchInput.value = '';
        clearBtn && clearBtn.classList.remove(showBtn);
        return false;
      }

      debounced = debounce(() => {
        const val = ev.target.value.trim().toLowerCase();
        curAppList.forEach(app => {
          if (app.url.includes(val) || app.title.includes(val) || app.desc.includes(val)) {
            app.ele.classList.remove(hide);
          } else {
            app.ele.classList.add(hide);
          }
        });
        curBookmarkList.forEach(bm => {
          let hadShowItem = false;
          bm.items.forEach(item => {
            if (item.url.includes(val) || item.title.includes(val)) {
              item.ele.classList.remove(hide);
              hadShowItem = true;
            } else {
              item.ele.classList.add(hide);
            }
          });

          if (hadShowItem) {
            bm.ele.classList.remove(hide);
          } else {
            bm.ele.classList.add(hide);
          }
        });

        if (clearBtn) {
          if (val) {
            clearBtn.classList.add(showBtn);
          } else {
            clearBtn.classList.remove(showBtn);
          }
        }
      }, 250);
    });

    if (clearBtn) {
      clearBtn.addEventListener('click', ev => {
        ev.stopPropagation();
        restoreDisplay(curAppList, curBookmarkList);
        searchInput.value = '';
        clearBtn.classList.remove(showBtn);
      });
    }
  }
})();
