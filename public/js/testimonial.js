// // // create class
// // class Testimonial {
// //   #quote = "";
// //   #image = "";

// //   constructor(quote, image) {
// //     //method inisialisasi
// //     this.#quote = quote;
// //     this.#image = image;
// //   }

// //   get quote() {
// //     return this.#quote;
// //   }

// //   get image() {
// //     return this.#image;
// //   }

// //   get author() {
// //     throw new Error("getAuthor() method must be implemented");
// //   }
// //   //   pholymorpic
// //   get testimonialHTML() {
// //     return `<div class="testimonial">
// //         <img src="${this.image}" class="profile-testimonial" />
// //         <p class="quote">${this.quote}</p>
// //         <p class="author">${this.author}</p>
// //       </div>`;
// //   }
// // }

// // //   sublass
// // class AuthorTestimonials extends Testimonial {
// //   #author = "";

// //   constructor(author, quote, image) {
// //     super(quote, image);
// //     this.#author = author;
// //   }

// //   get author() {
// //     return this.#author;
// //   }
// // }
// // //   Subclass
// // class CompanyTestimonials extends Testimonial {
// //   #company = "";

// //   constructor(company, quote, image) {
// //     super(quote, image);
// //     this.#company = company;
// //   }

// //   get author() {
// //     return this.#company + " Company";
// //   }
// // }

// // const testimonial1 = new AuthorTestimonials(
// //   "Mohammad Fatkhuroji",
// //   "cintailah produk-produk indonesia",
// //   "https://images.unsplash.com/photo-1453396450673-3fe83d2db2c4?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D&auto=format&fit=crop&w=687&q=80"
// // );
// // const testimonial2 = new AuthorTestimonials(
// //   "bambang suryadi",
// //   "sangat mengkeren",
// //   "https://images.unsplash.com/photo-1564564321837-a57b7070ac4f?ixlib=rb-4.0.3&ixid=MnwxMjA3fDB8MHxzZWFyY2h8OHx8bWFufGVufDB8fDB8fA%3D%3D&auto=format&fit=crop&w=500&q=60"
// // );
// // const testimonial3 = new CompanyTestimonials(
// //   "Daihatsu",
// //   "Desain Yang Mantap!! 🔥🔥🔥",
// //   "https://images.unsplash.com/photo-1568602471122-7832951cc4c5?ixlib=rb-4.0.3&ixid=MnwxMjA3fDB8MHxzZWFyY2h8M3x8bWFufGVufDB8fDB8fA%3D%3D&auto=format&fit=crop&w=500&q=60"
// // );

// // let testimonialData = [testimonial1, testimonial2, testimonial3];
// // let testimonialHTML = "";

// // for (let i = 0; i < testimonialData.length; i++) {
// //   testimonialHTML += testimonialData[i].testimonialHTML;
// // }

// // document.getElementById("testimonial").innerHTML = testimonialHTML;

// // hof and callback
// const testimonialData = [
//   {
//     author: "Mohammad Fatkhuroji",
//     quote: "cintailah produk-produk indonesia",
//     image: "https://images.unsplash.com/flagged/photo-1570612861542-284f4c12e75f?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D&auto=format&fit=crop&w=870&q=80",
//     rating: 5,
//   },
//   {
//     author: "Andrey Arshavin",
//     quote: "Come on Young Gunners",
//     image: "https://images.unsplash.com/photo-1564564321837-a57b7070ac4f?ixlib=rb-4.0.3&ixid=MnwxMjA3fDB8MHxzZWFyY2h8OHx8bWFufGVufDB8fDB8fA%3D%3D&auto=format&fit=crop&w=500&q=60",
//     rating: 3,
//   },
//   {
//     author: "Bukayo Saka",
//     quote: "you are great",
//     image: "https://images.unsplash.com/photo-1453396450673-3fe83d2db2c4?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D&auto=format&fit=crop&w=687&q=80",
//     rating: 3,
//   },
// ];

// function allTestimonials() {
//   let testimonialHTML = "";

//   testimonialData.forEach(function (item) {
//     testimonialHTML += `
//     <div class="testimonial">
//     <img src="${item.image}" class="profile-testimonial" />
//       <p class="quote">${item.quote}</p>
//       <p class="author">${item.author}</p>
//       <p class="author">${item.rating} <i class = "fa-solid fa-star"></i></p>
//     </div>
//   `;
//   });

//   document.getElementById("testimonial").innerHTML = testimonialHTML;
// }

// allTestimonials();

// function filterTestimonials(rating) {
//   let testimonialHTML = "";

//   const testimonialFiltered = testimonialData.filter(function (item) {
//     return item.rating === rating;
//   });

//   if (testimonialFiltered.length === 0) {
//     testimonialHTML += `<h1> Data not Found!</h1>`;
//   } else {
//     testimonialFiltered.forEach(function (item) {
//       testimonialHTML += `
//       <div class="testimonial">
//         <img src="${item.image}" class="profile-testimonial" />
//         <p class="quote">${item.quote}</p>
//         <p class="author">${item.author}</p>
//         <p class="author">${item.rating} <i class = "fa-solid fa-star"></i></p>
//       </div>
//     `;
//     });
//   }

//   document.getElementById("testimonial").innerHTML = testimonialHTML;
// }
