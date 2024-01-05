import { dqs, di } from "./lib.js";

const bannerInput = di("banner-input");
const pictureInput = di("picture-input");

const linkTemplate = di("link-template");
const historyTemplate = di("history-template");

di("banner-button").addEventListener("click", () => {
  bannerInput.click();
});

bannerInput.addEventListener("change", () => {
  if (bannerInput.files.length < 1) {
    return
  }

  const file = bannerInput.files[0]
  const url = URL.createObjectURL(file);
  di("banner").src = url
})

di("picture-button").addEventListener("click", () => {
  pictureInput.click();
})

pictureInput.addEventListener("change", () => {
  if (pictureInput.files.length < 1) {
    return
  }

  const file = pictureInput.files[0];
  const url = URL.createObjectURL(file);
  di("picture").src = url;
});

di("link-button").addEventListener("click", () => {
  newLink = linkTemplate.content.cloneNode(true);
  di("links").insertBefore(newLink, di("link-button"));
});

di("history-button").addEventListener("click", () => {
  newHistory = historyTemplate.content.cloneNode(true);
  di("histories").appendChild(newHistory);
});

let isLoading = false;

di("submit").addEventListener("click", async () => {
  if (isLoading) return;
  isLoading = true;
  di("check").classList.add("hidden");
  di("spinner").classList.remove("hidden");

  const history = dqs("#history").map(h => {
    const title = h.querySelector("#title-input").value;
    const content = h.querySelector("#content-input").value;
    return { title, content };
  })

  const profile = {
    name: di("name-input").value,
    description: di("description-input").value,
    instagramId: di("instagram-input").value,
    facebookId: di("facebook-input").value,
    twitterId: di("twitter-input").value,
    links: dqs("#link-input").map(el => el.value),
    note: di("note-input").value,
    history: history,
  };

  const res = await fetch("/dashboard/profile", {
    method: "POST",
    body: JSON.stringify(profile),
  });
  const data = await res.json();

  if (bannerInput.files.length > 0) {
    await fetch(data.bannerUrl, {
      method: "PUT",
      body: bannerInput.files[0]
    });
  }

  if (pictureInput.files.length > 0) {
    await fetch(data.pictureUrl, {
      method: "PUT",
      body: pictureInput.files[0],
    });
  }

  isLoading = false;
  di("check").classList.remove("hidden");
  di("spinner").classList.add("hidden");
})
