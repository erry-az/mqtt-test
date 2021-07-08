function handleEnter(e) {
  console.log({ key: e.key });
  if (e.key !== "Enter" || e.shiftKey) return;

  e.preventDefault();
  const { value } = document.getElementById("chat");
  if (value.length < 1) return;

  onSend();
}

function handleEsc(e) {
  if (e.key !== "Escape") return;
  toggleModal();
}
