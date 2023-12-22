const dq = document.querySelector.bind(document);

const thumbnailList = dq("#thumbnail-list");
const addButton = dq("#add-button");
const imageInput = dq("#image-input");
const preview = dq("#preview")
const thumbnailTemplate = dq("#thumbnail-template")

const imageList = [];

addButton.addEventListener("click", () => {
  imageInput.click();
});

imageInput.addEventListener("change", () => {
  if (imageInput.files.length < 1) {
    return
  }

  const file = imageInput.files[0]
  const url = URL.createObjectURL(file);
  imageInput.value = ""
  addImage(file, url);
  selectImage(url);
})

let draggedItem = null;

function addImage(image, url) {
  imageList.push({ image, url })
  const thumbnail = thumbnailTemplate.cloneNode(true);
  const thumbnailImage = thumbnail.querySelector("#thumbnail-image");
  const thumbanilRemoveButton = thumbnail.querySelector("#thumbnail-remove");
  thumbnail.id = "thumbnail";
  thumbnail.style.display = "block";
  thumbnailImage.src = url;
  thumbnailImage.addEventListener("click", () => selectImage(url));
  thumbanilRemoveButton.addEventListener("click", () => removeImage(thumbnail));
  thumbnailList.appendChild(thumbnail)
}

function removeImage(thumbnail) {
  const idx = imageList.findIndex(({ itemUrl }) => itemUrl === thumbnail.src);
  if (idx !== -1) imageList.splice(idx, 0);
  thumbnailList.removeChild(thumbnail)
  preview.removeAttribute("src")
  selectImage(null);
}

function selectImage(url) {
  if (url !== null) {
    preview.classList.remove("bg-gray-100");
  } else {
    preview.classList.add("bg-gray-100");
  }
  preview.src = url;
}

let isLoading = false;
const submit = dq("#submit");
const spinner = dq("#spinner");
const check = dq("#check");

submit.addEventListener("click", async () => {
  if (isLoading) return;
  isLoading = true;
  check.classList.add("hidden");
  spinner.classList.remove("hidden");


  const payload = {
    name: dq("#name").value,
    description: dq("#description").value,
    price: parseInt(dq("#price").value),
    info: dq("#info").value,
    imageLength: imageList.length,
    tags: tags
  }

  const res = await fetch("/dashboard/art/new", {
    method: "POST",
    body: JSON.stringify(payload),
  });
  // TODO: handle error 

  const data = await res.json();
  const uploadUrls = data.uploadUrls

  await Promise.all(uploadUrls.map(async (url, idx) => {
    return await fetch(url, {
      method: "PUT",
      body: imageList[idx].image,
    })
  }));


  isLoading = false;
  check.classList.remove("hidden");
  spinner.classList.add("hidden");
})

const tagList = dq("#tag-list")
const tagTemplate = dq("#tag-template")
const tagInput = dq("#tag-input")
const tags = []

tagList.addEventListener("click", () => {
  tagInput.focus();
})

tagInput.addEventListener("keydown", (e) => {
  if (e.key === "Enter") {
    const isNew = addTag(tagInput.value);
    if (isNew) tagInput.value = "";
  } else if (e.key === "Backspace" && tagInput.value.length === 0 && tags.length > 0) {
    const lastTag = tagList.querySelector(".tag:last-of-type")
    removeTag(lastTag, tags[tags.length - 1]);
  }
})

function addTag(tagName) {
  if (tags.findIndex((tn) => tn === tagName) !== -1) return false;

  tags.push(tagName);
  const tag = tagTemplate.cloneNode(true);
  tag.querySelector("#tag-name").innerHTML = tagName;
  tag.id = null
  tag.classList.add("tag");
  tag.classList.remove("hidden");
  tag.classList.add("flex");

  tag.querySelector("#remove-tag").addEventListener("click", () => {
    removeTag(tag, tagName);
  })
  tagList.insertBefore(tag, tagInput);
  return true;
}

function removeTag(elem, name) {
  const idx = tags.findIndex((tn) => tn === name);
  tags.splice(idx, 1);

  tagList.removeChild(elem);
}
