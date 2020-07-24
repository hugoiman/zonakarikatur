window.onload = function () {
  getAdmin();
};
var token = Cookies.get("token");

function getAdmin() {
  $.ajax({
    url: domain + "/api/admin",
    type: "GET",
    headers: { Authorization: "Bearer " + token },
    success: function (resp) {
      $(".name").text(resp.name);
      $(".name").val(resp.name);
      $(".username").val(resp.username);
      $(".email").val(resp.email);
    },
    error: function (error) {
      window.location.href = "/login";
    },
  });
}

function deleteAlert(api, idImage) {
  Swal.fire({
    title: "Are you sure?",
    text: "You won't be able to revert this!",
    icon: "warning",
    showCancelButton: true,
    confirmButtonColor: "#3085d6",
    cancelButtonColor: "#d33",
    confirmButtonText: "Yes, delete it!",
  }).then((result) => {
    if (result.value) {
      $.ajax({
        url: domain + api,
        type: "DELETE",
        headers: { Authorization: "Bearer " + token },
        success: function (resp) {
          $("#image" + idImage).fadeOut(1000);
        },
        error: function (error) {
          failedAlert(error.responseText);
        },
      });
    }
  });
}

function failedAlert(error) {
  Swal.fire({
    icon: "error",
    title: "Failed!",
    text: error,
  });
}

function logout() {
  Cookies.remove("token", { path: "/" });
  window.location.href = "/login";
}
