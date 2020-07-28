function getAbout() {
  $.ajax({
    url: domain + "/api/about",
    type: "GET",
    headers: { Authorization: "Bearer " + token },
    success: function (resp) {},
    error: function (error) {
      failedAlert(error.responseText);
    },
  });
}
