import { CutRequestBody, CutResponse } from "./types";

const formElement = document.getElementById("cut-form");
const resultElement = document.getElementById("cut-result");

const formInput = document.getElementById("cut-form-url") as HTMLInputElement;
const resultInput = document.getElementById("cut-result-input") as HTMLInputElement;

const resultCopyButton = document.getElementById("cut-result-copy") as HTMLButtonElement;

function toggleElementsVisiblity() {
    formElement.classList.toggle("hidden");
    resultElement.classList.toggle("hidden");
}

function nextUrl(initValue = "") {
    toggleElementsVisiblity();

    resultInput.value = "";
    formInput.value = initValue;
}

async function submit(event: SubmitEvent) {
    event.preventDefault();

    const url = formInput.value;

    const reponse = await fetch("/cut", { method: 'POST', body: JSON.stringify({ url } as CutRequestBody) });
    const data: CutResponse = await reponse.json();

    resultInput.value = data.url;
    toggleElementsVisiblity();

    resultInput.addEventListener("input", () => {
        nextUrl(resultInput.value);
        formInput.focus();
    }, { once: true });
}

function copyUrl() {
    navigator.clipboard.writeText(resultInput.value);
    resultCopyButton.innerText = "Copied"
    setTimeout(() => resultCopyButton.innerText = "Copy", 3 * 1000);
    
}

formElement.addEventListener("submit", submit);
resultCopyButton.addEventListener("click", copyUrl);