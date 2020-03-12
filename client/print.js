export const print = message => {
  const output = document.getElementById("output")
  const d = document.createElement("div");
  d.innerHTML = message;
  output.appendChild(d);
};