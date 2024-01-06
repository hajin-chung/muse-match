/** @type { (id: string) => HTMLElement } */
export const di = document.getElementById.bind(document);
/** @type { (qs: string) => HTMLElement } */
export const dq = qs => document.querySelector(qs);
/** @type { (qs: string) => HTMLElement[] } */
export const dqs = qs => [...document.querySelectorAll(qs)];

// document.addEventListener('click', function(e) {
//   let targetElement = e.target;
//   while (targetElement && targetElement.tagName !== 'A') {
//     targetElement = targetElement.parentElement;
//   }
//
//   // If an 'a' tag was found
//   if (targetElement) {
//     e.preventDefault(); // Prevent the default link behavior
//     const href = targetElement.getAttribute('href');
//     loadContent(href);
//   }
// });
//
// function random() {
//   return Math.ceil(Math.random() * 100)
// }
//
// function reloadScripts() {
//   dqs("body script").forEach((s) => {
//     const parent = s.parentNode;
//     const src = `${s.src}?${random()}`
//     parent.removeChild(s)
//     const script = document.createElement("script");
//     script.src = src
//     script.type = "module"
//     parent.appendChild(script)
//   })
//
// }
//
// function loadContent(url) {
//   const currentUrl = document.location;
//   fetch(url)
//     .then(response => response.text())
//     .then(html => {
//       const parser = new DOMParser();
//       const newHtml = parser.parseFromString(html, "text/html")
//       const body = newHtml.querySelector("body")
//       document.body.innerHTML = body.innerHTML;
//       if (url.href !== currentUrl.href)
//         history.pushState({ body: document.body.innerHTML }, '', url);
//       reloadScripts();
//     });
// }
//
// window.addEventListener('popstate', function(event) {
//   if (event.state && event.state.body) {
//     // If there's content in the state, use it
//     document.body.innerHTML = event.state.body;
//     reloadScripts();
//   } else {
//     // Otherwise, you might want to load content based on the URL
//     loadContent(document.location);
//   }
// });
//
// window.onload = function() {
//   loadContent(document.location);
// };
