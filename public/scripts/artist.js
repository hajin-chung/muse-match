import { di, dq } from "./lib.js"

/** type {("grid" | "list")} */
let mode = "grid"
di("list").classList.add("hidden");

di("grid-button").onclick = () => {
  if (mode === "grid") return

  mode = "grid";
  di("grid-button").classList.add("bg-black");
  dq("#grid-button > img").classList.add("invert");
  dq("#grid-button > p").classList.add("text-white");
  di("list-button").classList.remove("bg-black");
  dq("#list-button > img").classList.remove("invert");
  dq("#list-button > p").classList.remove("text-white");

  di("list").classList.add("hidden");
  di("grid").classList.remove("hidden");
}

di("list-button").onclick = () => {
  if (mode === "list") return

  mode = "list";
  di("list-button").classList.add("bg-black");
  dq("#list-button > img").classList.add("invert");
  dq("#list-button > p").classList.add("text-white");
  di("grid-button").classList.remove("bg-black");
  dq("#grid-button > img").classList.remove("invert");
  dq("#grid-button > p").classList.remove("text-white");

  di("grid").classList.add("hidden");
  di("list").classList.remove("hidden");
}
