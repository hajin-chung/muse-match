import { dq, di, dqs } from "./lib.js"

const placeId = document.location.href.split("/").at(-1)

const thumbnailList = di("thumbnail-list")
const thumbnailImageInput = di("thumbnail-image-input")

dqs(".thumbnail").forEach(thumbnailImage => {
  const thumbnail = thumbnailImage.parentNode
  thumbnail.onclick = () => selectImage(thumbnailImage.src)
  thumbnail.querySelector(".delete").onclick = () => 
    thumbnailList.removeChild(thumbnail)
})

di("thumbnail-button").onclick = () => {
  di("thumbnail-image-input").click()
}

thumbnailImageInput.onchange = () => {
  if (thumbnailImageInput.files.length < 1) {
    return
  }

  const file = thumbnailImageInput.files[0]
  const url = URL.createObjectURL(file)
  di("preview").src = url
  selectImage(url)

  let newThumbnail = di("thumbnail-template").content.cloneNode(true)
  di("thumbnail-list").appendChild(newThumbnail)
  newThumbnail = thumbnailList.lastChild
  newThumbnail.querySelector(".thumbnail").src = url
  newThumbnail.querySelector(".delete").onclick = () => {
    thumbnailList.removeChild(newThumbnail)
  }
  newThumbnail.onclick = () => selectImage(url)
}

function selectImage(url) {
  const preview = di("preview")
  if (url !== null) {
    preview.classList.remove("bg-gray-100");
  } else {
    preview.classList.add("bg-gray-100");
  }
  preview.src = url;
}

const linkList = di("link-list")
di("link-button").onclick = () => {
  let newLink = di("link-template").content.cloneNode(true)
  linkList.insertBefore(newLink, di("link-button"))
}

const locationList = di("location-list")
const locationImageInput = di("location-image-input")
let locationTarget = null

dqs(".location").forEach(location => {
  location.querySelector(".upload-button").onclick = () => {
    locationTarget = location
    locationImageInput.click()
  }
  location.querySelector(".delete-button").onclick = () =>
    locationList.removeChild(location)
})

di("location-button").onclick = () => {
  let newLocation = di("location-template").content.cloneNode(true)
  locationList.appendChild(newLocation)
  newLocation = locationList.lastChild
  newLocation.querySelector(".upload-button").onclick = () => {
    locationTarget = newLocation
    locationImageInput.click()
  }
  newLocation.querySelector(".delete-button").onclick = () => {
    locationList.removeChild(newLocation)
  }
}

locationImageInput.onchange = () => {
  if (locationTarget === null) return
  if (locationImageInput.files.length < 1) return

  const file = locationImageInput.files[0]
  const url = URL.createObjectURL(file)
  locationTarget.querySelector(".image").src = url
  locationTarget.querySelector(".image").style.backgroundColor = "white"
}

let isLoading = false;
di("submit").onclick = async () => {
  if (isLoading) return
  isLoading = true
  check.classList.add("hidden");
  spinner.classList.remove("hidden");

  const locations = dqs(".location").map((location) => ({
    title: location.querySelector(".title-input").value,
    description: location.querySelector(".description-input").value,
  }))

  const payload = {
    title: di("title-input").value,
    address: di("address-input").value,
    instagramId: di("instagram-input").value,
    facebookId: di("facebook-input").value,
    twitterId: di("twitter-input").value,
    links: dqs(".link-input").map((linkInput) => linkInput.value),
    imageCount: dqs(".thumbnail").length,
    locations,
  }

  const res = await fetch(`/dashboard/place/${placeId}`, {
    method: "POST",
    body: JSON.stringify(payload)
  })
  const data = await res.json();

  await Promise.all(dqs(".thumbnail").map(async (thumbnail, idx) => {
    const res = await fetch(thumbnail.src)
    const blob = await res.blob()
    await fetch(data.imageUrls[idx], {
      method: "PUT",
      body: blob
    })
  }))

  await Promise.all(dqs(".location .image").map(async (locationImage, idx) => {
    const res = await fetch(locationImage.src)
    const blob = await res.blob()
    await fetch(data.locationImageUrls[idx], {
      method: "PUT",
      body: blob,
    })
  }))

  isLoading = false
  check.classList.remove("hidden");
  spinner.classList.add("hidden");
}
