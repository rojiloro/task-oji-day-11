function submitData() {
  let name = document.getElementById("input-name").value;
  let email = document.getElementById("input-email").value;
  let phone = document.getElementById("input-phone").value;
  let subject = document.getElementById("subject").value;
  let message = document.getElementById("input-message").value;

  if (name == "") {
    return alert("Nama Harus diisi");
  } else if (email == "") {
    return alert("Email harus diisi");
  } else if (phone == "") {
    return alert("Telepon harus diisi");
  } else if (subject == "") {
    return alert("Subject Harus diisi");
  }

  let emailReceiver = "kelengoji79@gmail.com";

  let a = document.createElement("a");
  a.href = `https://mail.google.com/mail/?view=cm&fs=1&to=:${emailReceiver}?subject=${subject}&body=Halo, Nama saya ${name}, ${message}. Silahkan kontak saya ${phone}`;
  a.click();

  console.log(name);
  console.log(email);
  console.log(phone);
  console.log(subject);
  console.log(message);
}
