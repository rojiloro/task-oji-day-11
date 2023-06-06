let dataProject = [];

function addProject(event) {
  event.preventDefault();
  // deklarasi variabel
  let nama = document.getElementById("input-project-name").value;
  let startDate = document.getElementById("input-start-date").value;
  let endDate = document.getElementById("input-end-date").value;
  let description = document.getElementById("description").value;
  let image = document.getElementById("input-project-image").files;

  // Mencari durasi project
  let date1 = new Date(startDate);
  let date2 = new Date(endDate);
  let selisih = date2 - date1;

  let milisecond = 1000;
  let secondInHours = 3600;
  let hoursInDay = 24;
  let dayInMonth = 30;

  let distanceDay = Math.floor(selisih / (milisecond * secondInHours * hoursInDay));
  let distanceMonth = Math.floor(selisih / (milisecond * secondInHours * hoursInDay * dayInMonth));
  if (distanceDay > 0) {
    console.log(distanceDay + " hari");
  } else if (distanceDay > distanceMonth) {
    console.log(distanceMonth + " bulan" + (distanceDay - distanceMonth) + " hari");
  } else if (distanceDay < 0) {
    alert("tanggal yang dimasukan salah");
  }


  // icon
  const playstorejsIcon = '<i class="fa-brands fa-google-play"></i>';
  const androidjsIcon = '<i class="fa-brands fa-android"></i>';
  const javajsIcon = '<i class="fa-brands fa-java"></i>';
  const reactIcon = '<i class="fa-brands fa-react"></i>';

  // cek checkbox

  let cbPlaystorejs = document.getElementById("playstore").checked ? playstorejsIcon : "";
  let cbAndroidjs = document.getElementById("android").checked ? androidjsIcon : "";
  let cbJavajs = document.getElementById("java").checked ? javajsIcon : "";
  let cbReactjs = document.getElementById("react").checked ? reactIcon : "";

  // cek kondisi image
  let imageURL = URL.createObjectURL(image[0]);
  // if (image.target.files.length !== 0) {
  //   this.setState({ image: URL.createObjectURL(e.target.files[0]) });
  // }

  let data = {
    nama,
    startDate,
    endDate,
    distanceDay,
    distanceMonth,
    cbPlaystorejs,
    cbAndroidjs,
    cbJavajs,
    cbReactjs,
    description,
    imageURL,
    postAt: new Date(),
  };

  dataProject.push(data);
  console.log(dataProject);

  renderDetail();
}

function renderDetail() {
    document.getElementById("detail") = ""; 
    
    for (let i = 0; i < dataProject.length; i++){
     document.getElementById('detail').innerHTML += `
     <div class="detail">
         <div class="container-detail" id="detail">
           <h1>${dataProject[i].nama}</h1>
           <div class="hero">
             <div>
               <img src="${dataProject[i].imageURL}"  />
             </div>
             <div>
               <p class="judul">Duration</p>
               <div class="text">
                 <i class="fa-regular fa-calendar"></i>
                 <p>${dataProject[i].startDate} - ${dataProject[i].endDate} </p>
               </div>
               <p class="judul">Technologies</p>
               <div class="tech">
                 <div class="jejer">
                   ${dataProject[i].cbPlaystorejs}
                   <p>Playstore</p>
                   ${dataProject[i].cbAndroidjs}
                   <p>Android</p>
                   </div>
                   <div class="jejer">
                   ${dataProject[i].cbJavajs}
                   <p>Java</p>
                   ${dataProject[i].cbReactjsjs}
                   <p>React</p>
                 </div>
               </div>
             </div>
           </div>
           <div class="paragraph">
             <p>${dataProject[i].description}</p>
           </div>
         </div>
       </div>
     `; 
    }
   }