import { getShortenedUrl } from './apiClient';

const formElement = document.getElementById('cut-form') as HTMLElement;
const resultElement = document.getElementById('cut-result') as HTMLElement;

const formInput = document.getElementById('cut-form-url') as HTMLInputElement;
const resultInput = document.getElementById('cut-result-input') as HTMLInputElement;

const resultCopyButton = document.getElementById('cut-result-copy') as HTMLButtonElement;

function toggleElementsVisiblity() {
  formElement.classList.toggle('hidden');
  resultElement.classList.toggle('hidden');
}

function nextUrl(initValue = '') {
  toggleElementsVisiblity();

  resultInput.value = '';
  formInput.value = initValue;

  formInput.focus();
}

async function submit(event: SubmitEvent) {
  event.preventDefault();

  const url = await getShortenedUrl(formInput.value);

  resultInput.value = url;
  toggleElementsVisiblity();

  resultInput.addEventListener('input', () => nextUrl(resultInput.value), { once: true });
}

function copyUrl() {
  navigator.clipboard.writeText(resultInput.value);
  resultCopyButton.innerText = 'Copied';
  setTimeout(() => (resultCopyButton.innerText = 'Copy'), 3 * 1000);
}

formElement.addEventListener('submit', submit);
resultCopyButton.addEventListener('click', copyUrl);
