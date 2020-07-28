function getLinkOrder() {
  $.ajax({
    url: domain + "/api/link",
    type: "GET",
    success: function (resp) {
      $("#link").val(resp.link);
      $("#idLink").val(resp.idLink);
    },
    error: function (error) {
      failedAlert(error.responseText);
    },
  });
}

function updateLinkOrder() {
  var link = $("#link").val();
  var idLink = $("#idLink").val();
  var jsonData = JSON.stringify({
    link,
  });
  $.ajax({
    url: domain + "/api/link/" + idLink,
    type: "PUT",
    headers: { Authorization: "Bearer " + token },
    data: jsonData,
    contentType: "application/json",
    success: function (resp) {
      Swal.fire("Success", "Link telah disimpan.", "success");
    },
    error: function (error) {
      failedAlert(error.responseText);
    },
  });
}
