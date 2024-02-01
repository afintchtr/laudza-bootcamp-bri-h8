let imageUrl;

const loadImage = async (url) => {
  try {
    const response = await url;
    fetch(url);

    const image = document.getElementById("image");
    image.setAttribute("src", response);
    image.style.display = "block";
  } catch (err) {
    const imageContainer = document.getElementById("imageContainer");
    imageContainer.innerHTML = "Gambar tidak ditemukan!";
  }
};

const imageLinkButton = document.getElementById("formForImage");
imageLinkButton.addEventListener("submit", (event) => {
  event.preventDefault();
  imageUrl = document.getElementById("imageLink").value;
  loadImage(imageUrl);
});
