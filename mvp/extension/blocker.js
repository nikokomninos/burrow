const getCurrentTab = async () => {
  const queryOptions = { active: true, lastFocusedWindow: true };
  const [tab] = await chrome.tabs.query(queryOptions);
  return tab ?? null;
}

const getCurrentURL = (tab) => {
  return tab?.url ?? null;
}

const sendURL = async () => {
  const tab = await getCurrentTab()
  const url = getCurrentURL(tab);
  const res = await fetch("http://localhost:3598/block", {
    method: "POST",
    headers: {
      "Content-Type": "application/json"
    },
    body: JSON.stringify({ url }),
  });
  const data = await res.text();
  console.log(data);
  if (data === "URL is in blocklist" && tab?.id) {
    await chrome.scripting.executeScript({
      target: { tabId: tab.id },
      func: () => {
        document.body.innerHTML = "<h1>Blocked by Burrow</h1>";
        document.addEventListener('contextmenu', e => e.preventDefault());
      },
    });
  }
}

const btn = document.getElementById("get-url");
btn.addEventListener("click", async () => {
  sendURL()
});
