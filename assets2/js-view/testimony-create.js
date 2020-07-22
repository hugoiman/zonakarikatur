function uploadFile() {
  var image = $("#image")[0].files;
  var fd = new FormData();
  var today = new Date();

  var fileName = "";

  for (var i = 0; i < image.length; i++) {
    fileName =
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

    fileName = fileName + "-" + image[i].name;
    fileName = fileName.split(" ").join("-");
    fd.append("files", image[i], fileName);
  }

  $.ajax({
    url: domain + "/api/testimony-file",
    type: "POST",
    headers: { Authorization: "Bearer " + token },
    data: fd,
    processData: false,
    contentType: false,
    success: function (resp) {
      for (var j = 0; j < resp.length; j++) {
        createTestimony(resp[j]);
      }
    },
    error: function (error) {
      failedAlert(error.responseText);
    },
  });
}

function createTestimony(fileName) {
  var image = fileName;
  var jsonData = JSON.stringify({
    image,
  });

  $.ajax({
    url: domain + "/api/testimony",
    type: "POST",
    headers: { Authorization: "Bearer " + token },
    data: jsonData,
    contentType: "application/json",
    success: function (resp) {
      Swal.fire("Success", "Data telah disimpan.", "success");
      $("#image").val("");
    },
    error: function (error) {
      failedAlert(error.responseText);
    },
  });
}
