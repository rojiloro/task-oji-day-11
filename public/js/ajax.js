const promise = new Promise((resolve, reject) => {
  const xhr = new XMLHttpRequest();
  xhr.open("GET", "https://api.npoint.io/498518b2c2b7d0dd4a5b", true);
  //   console.log(xhr);
  xhr.onload = () => {
    if (xhr.status === 200) {
      resolve(JSON.parse(xhr.response));
    } else {
      reject("Error loading data.");
    }
  };

  xhr.onerror = () => {
    reject("Network error.");
  };

  xhr.send();
});

async function getAllTestimonial() {
  const response = await promise;
  console.log(response);
  let testimonialHTML = "";

  response.forEach(function (item) {
    testimonialHTML += `
    <div class="testimonial">
          <img src="${item.image}" class="profile-testimonial" />
          <p class="quote">${item.quote}</p>
          <p class="author">${item.author}</p>
          <p class="author">${item.rating} <i class="fa-solid fa-star"></i></p>
        </div>`;
  });

  document.getElementById("testimonial").innerHTML = testimonialHTML;
}

getAllTestimonial();

async function getFilteredTestimonials(rating) {
  const response = await promise;

  const testimonialFiltered = response.filter((item) => {
    return item.rating === rating;
  });

  //   console.log(testimonialFiltered);

  let testimonialHTML = "";
  if (testimonialFiltered.length === 0) {
    testimonialHTML = "<h1> Data Not found!</h1>";
  } else {
    testimonialFiltered.forEach((item) => {
      testimonialHTML += `
        <div class="testimonial">
        <img src="${item.image}" class="profile-testimonial" />
        <p class="quote">${item.quote}</p>
        <p class="author">${item.author}</p>
        <p class="author">${item.rating}<i class="fa-solid fa-star"></i></p>
      </div>`;
    });
  }

  document.getElementById("testimonial").innerHTML = testimonialHTML;
}
