import { dqs, di } from "./lib.js";

const bannerInput = di("banner-input");
const pictureInput = di("picture-input");
const histories = di("histories");
const lists = di("lists");

const linkTemplate = di("link-template");
const historyTemplate = di("history-template");
const listTemplate = di("list-template");

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
  const newLink = linkTemplate.content.cloneNode(true);
  di("links").insertBefore(newLink, di("link-button"));
});

dqs(".history").forEach((history) => {
  history.querySelector(".delete-button").onclick = () => {
    history.parentNode.removeChild(history)
  }
})

di("history-button").addEventListener("click", () => {
  let newHistory = historyTemplate.content.cloneNode(true);
  histories.appendChild(newHistory);
  newHistory = histories.lastChild;
  newHistory.querySelector(".delete-button").onclick = () => {
    histories.removeChild(newHistory);
  }
});

dqs(".art-list").forEach(list => {
  list.querySelector(".delete-button").onclick = () => {
    lists.removeChild(list);
  }
  list.querySelector(".update-button").onclick = () => {
    openModal(list)
  }
})

di("list-button").addEventListener("click", () => {
  let newList = listTemplate.content.cloneNode(true);
  lists.insertBefore(newList, di("list-button"));
  newList = lists.children[lists.children.length - 2]
  newList.querySelector(".delete-button").onclick = () => {
    lists.removeChild(newList);
  }
  newList.querySelector(".update-button").onclick = () => openModal(newList)
});

const modal = di("modal");
const artCards = dqs("#modal .art-card");
let modalTarget;
/** @type {string[]} */
let modalList = [];

artCards.forEach(artCard => {
  artCard.onclick = () => {
    const num = artCard.querySelector(".number");
    if (num.style.display === "flex") {
      // remove num and shift numbers
      const index = modalList.findIndex((id) => id === artCard.id);
      for (let i = index + 1; i < modalList.length; i++) {
        modal.querySelector(`#${modalList[i]}`).querySelector(".number").innerText = i;
      }
      modalList.splice(index, 1);
      num.style.display = "none";
    } else {
      num.style.display = "flex";
      num.innerHTML = modalList.length + 1;
      modalList.push(artCard.id);
    }
  }
});

di("list-submit").onclick = closeModal;

function openModal(target) {
  modalTarget = target;
  modalList = [];

  modal.style.display = "flex";

  modalList = [...target.querySelectorAll(".items > .art-card")].map(c => c.id);

  artCards.map(c => {
    const num = c.querySelector(".number")
    const idx = modalList.findIndex((id) => id === c.id)
    if (idx !== -1) {
      num.innerHTML = idx + 1;
      num.style.display = "flex";
    } else {
      num.innerHTML = "";
      num.style.display = "hidden";
    }
  })
}

function closeModal() {
  modal.style.display = "none";
  artCards.map(c => c.querySelector(".number")).forEach(num => {
    num.innerHTML = "";
    num.style.display = "none";
  })

  const items = modalTarget.querySelector(".items");
  items.innerHTML = "";
  modalList.forEach((id) => {
    const item = di(id).cloneNode(true);
    items.appendChild(item);
  });
}

let isLoading = false;

di("submit").addEventListener("click", async () => {
  if (isLoading) return;
  isLoading = true;
  di("check").classList.add("hidden");
  di("spinner").classList.remove("hidden");

  const history = dqs(".history").map(h => {
    const title = h.querySelector(".title-input").value;
    const content = h.querySelector(".content-input").value;
    return { title, content };
  });

  const list = dqs(".art-list").map(l => {
    const title = l.querySelector("#title").value;
    const artIds = [...l.querySelectorAll(".art-card")].map(c => c.id);
    return { title, artIds };
  })

  const profile = {
    name: di("name-input").value,
    description: di("description-input").value,
    instagramId: di("instagram-input").value,
    facebookId: di("facebook-input").value,
    twitterId: di("twitter-input").value,
    links: dqs(".link-input").map(el => el.value),
    note: di("note-input").value,
    history,
    list,
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
