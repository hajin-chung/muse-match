import { di } from "./lib.js"

const lat = parseFloat(di("map").getAttribute("lat"))
const lng = parseFloat(di("map").getAttribute("lng"))

const position = new naver.maps.LatLng(lat, lng)
let mapOptions = {
  center: position,
};

const map = new naver.maps.Map('map', mapOptions);
const marker = new naver.maps.Marker({ position, map })
