import { dqs, di } from "./lib.js"

dqs(".image-button").forEach(b => b.onclick = () => {
  di("focus").src = `/image?id=${b.id}`
})
