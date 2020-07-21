function uploadFile() {
  var fd = new FormData();
  var image = $("#image")[0].files[0];
  var today = new Date();

  var fileName =
    today.getFullYear() +
    "" +
    (today.getMonth() + 1) +
    "" +
    today.getDate() +
    "-" +
    today.getHours() +
    "" +
    today.getMinutes() +
    "" +
    today.getSeconds();

  var extension = image.name.substr(image.name.lastIndexOf(".") + 1);
  fileName = fileName + "." + extension;

  fd.append("files", image, fileName);

  $.ajax({
    url: domain + "/api/offer-file",
    type: "POST",
    headers: { Authorization: "Bearer " + token },
    data: fd,
    processData: false,
    contentType: false,
    success: function (resp) {
      createOffer(fileName);
    },
    error: function (error) {
      failedAlert(error.responseText);
    },
  });
}

function createOffer(fileName) {
  var title = $("#title").val();
  var image = fileName;
  var jsonData = JSON.stringify({
    title,
    image,
  });

  $.ajax({
    url: domain + "/api/offer",
    type: "POST",
    headers: { Authorization: "Bearer " + token },
    data: jsonData,
    contentType: "application/json",
    success: function (resp) {
      Swal.fire("Success", "Data telah disimpan.", "success");
      $("#title").val("");
      $("#image").val("");
    },
    error: function (error) {
      failedAlert(error.responseText);
    },
  });
}
