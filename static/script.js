document.addEventListener("DOMContentLoaded", function() {
    const shortenButton = document.getElementById("shortenButton");
    const shortenInput = document.getElementById("shortenInput");
    const customInput = document.getElementById("customInput");
    const shortAnotherButton = document.getElementById("shortAnotherButton");

    shortenButton.addEventListener("click", async function(event) {
        event.preventDefault();

        if (shortenButton.innerText === "Shorten") {
            await shorten();
        }

        if (shortenButton.innerText === "Copy") {
            copyToClipboard();
        }
    });

    shortAnotherButton.addEventListener("click", function(event) {
        event.preventDefault();
        shortenInput.value = "";
        shortenButton.innerText = "Shorten";
        shortAnotherButton.style.display = "none";
    });

    function copyToClipboard() {
        const copyText = document.getElementById("shortenInput");
        copyText.select();
        copyText.setSelectionRange(0, 99999);
        navigator.clipboard.writeText(copyText.value).then(_ => console.log("Copied to clipboard"));
    }

    async function shorten() {
        const url = shortenInput.value.trim();
        const custom = customInput.value.trim();
        if (url === "") {
            return;
        }

        // Regex to check if it is a valid url
        const urlRegex = new RegExp(/^(http|https):\/\/[^ "]+$/);
        if (!urlRegex.test(url)) {
            shortenInput.value = "Invalid URL";
            return;
        }

        const payload = { url };
        if (custom !== "") {
            payload.custom = custom;
        }

        try {
            const response = await fetch("/shorten", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json"
                },
                body: JSON.stringify({ url })
            });

            const data = await response.json();
            console.log(data);
            if (data.error) {
                shortenInput.value = data.error;
                return;
            }

            // get current website domain
            const port = window.location.port ? ":" + window.location.port : "";
            shortenInput.value = "https://" + window.location.hostname + port + window.location.port + "/" + data["short"];
            shortenButton.innerText = "Copy";

            shortAnotherButton.style.display = "inline-block";
        } catch (error) {
            console.error(error);
            shortenInput.value = "An error occurred";
        }
    }
});