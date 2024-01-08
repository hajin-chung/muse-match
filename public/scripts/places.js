import { di } from "./lib.js"

const refreshButton = di("refresh-button")
const placeTemplate = di("place-template")
const map = new naver.maps.Map('map');
let didChange = false;

navigator.geolocation.getCurrentPosition((pos) => {
  const lat = pos.coords.latitude
  const lng = pos.coords.longitude
  map.setOptions({
    center: { lat, lng }
  })
}, () => console.log("failed"))

naver.maps.Event.addListener(map, "center_changed", () => {
  if (!didChange) {
    didChange = true
    loadPlaces()
    return
  }

  showRefreshButton()
})

let markers = []

async function loadPlaces() {
  refreshButton.style.display = "none"

  const bounds = map.getBounds()
  const maxLat = bounds._max._lat
  const maxLng = bounds._max._lng
  const minLat = bounds._min._lat
  const minLng = bounds._min._lng

  const res = await fetch(`/place?max_lat=${maxLat}&max_lng=${maxLng}&min_lat=${minLat}&min_lng=${minLng}`);
  const data = await res.json()

  di("place-list").innerHTML = ""
  data.places.forEach((place) => {
    let placeElement = placeTemplate.content.cloneNode(true)
    di("place-list").appendChild(placeElement)
    placeElement = di("place-list").lastChild
    placeElement.href = `/place/${place.id}`
    placeElement.querySelector(".image").src = `/image?id=${place.thumbnail}`
    placeElement.querySelector(".title").innerHTML = place.title
    placeElement.querySelector(".art-count").innerHTML = `${place.artCount} 작품`
    placeElement.querySelector(".address").innerHTML = place.address
  })

  markers.forEach(marker => marker.setMap(null))
  markers = []
  data.places.forEach(place => {
    const position = new naver.maps.LatLng(place.lat, place.lng)
    const marker = new naver.maps.Marker({ position, map })
    markers.push(marker)
  })
}

function showRefreshButton() {
  refreshButton.style.display = "flex"
}

refreshButton.onclick = () => {
  loadPlaces()
}
