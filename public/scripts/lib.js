/** @type { (id: string) => HTMLElement } */
export const di = document.getElementById.bind(document);
/** @type { (qs: string) => HTMLElement } */
export const dq = qs => document.querySelector(qs);
/** @type { (qs: string) => HTMLElement[] } */
export const dqs = qs => [...document.querySelectorAll(qs)];

