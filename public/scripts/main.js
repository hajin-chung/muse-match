import { di, dq, dqs } from "./lib.js"

darkHeader()
function darkHeader() {
  di("header").classList.add("text-white")
  di("header").classList.remove("shadow-md")
  di("header").classList.remove("bg-white")
  dqs("#header img").forEach(img => img.classList.add("invert"))
  dq("#header #picture").classList.remove("invert")
}

function lightHeader() {
  di("header").classList.remove("text-white")
  di("header").classList.add("shadow-md")
  di("header").classList.add("bg-white")
  dqs("#header img").forEach(img => img.classList.remove("invert"))
}

const elementIsVisibleInViewport = (el, partiallyVisible = false) => {
  const { top, left, bottom, right } = el.getBoundingClientRect();
  const { innerHeight, innerWidth } = window;
  return partiallyVisible
    ? ((top >= 0 && top <= innerHeight) ||
      (bottom >= 0 && bottom <= innerHeight)) &&
    ((left > 0 && left < innerWidth) || (right > 0 && right < innerWidth))
    : top >= 0 && left >= 0 && bottom <= innerHeight && right <= innerWidth;
};


new Swiper('#banner', {
  loop: true,
  pagination: {
    el: '.swiper-pagination',
  },
});

document.onscroll = () => {
  if (elementIsVisibleInViewport(di("banner"), true)) {
    darkHeader()
  } else {
    lightHeader()
  }
}
