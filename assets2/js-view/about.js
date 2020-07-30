function getAbout() {
  $.ajax({
    url: domain + "/api/about",
    type: "GET",
    headers: { Authorization: "Bearer " + token },
    success: function (resp) {
      $("#description").html(resp.description);
      $("#idAbout").val(resp.idAbout);
      $("#description").summernote({
        toolbar: [
          ["style", ["bold", "italic", "underline"]],
          ["font", ["strikethrough", "superscript", "subscript"]],
        ],
        height: 150,
      });
    },
    error: function (error) {
      failedAlert(error.responseText);
    },
  });
}

function updateAbout() {
  var description = $("#description").summernote("code");
  var idAbout = $("#idAbout").val();
  var jsonData = JSON.stringify({
    description,
  });
  $.ajax({
    url: domain + "/api/about/" + idAbout,
    type: "PUT",
    headers: { Authorization: "Bearer " + token },
    data: jsonData,
    contentType: "application/json",
    success: function (resp) {
      Swal.fire("Success", "Deskripsi telah disimpan.", "success");
    },
    error: function (error) {
      failedAlert(error.responseText);
    },
  });
}
