function getFaq() {
  $.ajax({
    url: domain + "/api/faq",
    type: "GET",
    success: function (resp) {
      $.each(resp.faqs, function (idx, value) {
        $("#faq").append(
          '<div class="form-row" id="faq' +
            value.idFaq +
            '">' +
            '<div class="form-group col-md-5">' +
            "<label>Question</label>" +
            '<input type="text" class="form-control" readonly title="' +
            value.question +
            '" value="' +
            value.question +
            '"/>' +
            "</div>" +
            '<div class="form-group col-md-6">' +
            "<label>Answer</label>" +
            '<input type="text" class="form-control" readonly title="' +
            value.answer +
            '" value="' +
            value.answer +
            '"/>' +
            "</div>" +
            '<div class="form-group col-md-1">' +
            "<label>.</label>" +
            '<div class="input-group-append">' +
            '<button class="btn btn-danger" type="button" onclick=deleteFaq("' +
            value.idFaq +
            '")' +
            ">Delete</button>" +
            "</div>" +
            "</div>" +
            "</div>"
        );
      });
    },
    error: function (error) {
      failedAlert(error.responseText);
    },
  });
}

function addFaq() {
  var question = $("#question").val();
  var answer = $("#answer").val();
  var jsonData = JSON.stringify({
    question,
    answer,
  });

  $.ajax({
    url: domain + "/api/faq",
    type: "POST",
    headers: { Authorization: "Bearer " + token },
    data: jsonData,
    contentType: "application/json",
    success: function (resp) {
      $("#faq").append(
        '<div class="form-row" id="faq' +
          resp.idFaq +
          '">' +
          '<div class="form-group col-md-5">' +
          "<label>Question</label>" +
          '<input type="text" class="form-control" readonly title="' +
          question +
          '" value="' +
          question +
          '"/>' +
          "</div>" +
          '<div class="form-group col-md-6">' +
          "<label>Answer</label>" +
          '<input type="text" class="form-control" readonly title="' +
          answer +
          '" value="' +
          answer +
          '"/>' +
          "</div>" +
          '<div class="form-group col-md-1">' +
          "<label>.</label>" +
          '<div class="input-group-append">' +
          '<button class="btn btn-danger" type="button" onclick=deleteFaq("' +
          resp.idFaq +
          '")>Delete</button>' +
          "</div>" +
          "</div>" +
          "</div>"
      );
      $("#question").val("");
      $("#answer").val("");
      Swal.fire("Success", "Data telah disimpan.", "success");
    },
    error: function (error) {
      failedAlert(error.responseText);
    },
  });
}

function deleteFaq(idFaq) {
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
        url: domain + "/api/faq/" + idFaq,
        type: "DELETE",
        headers: { Authorization: "Bearer " + token },
        success: function (resp) {
          $("#faq" + idFaq).fadeOut(1000);
        },
        error: function (error) {
          failedAlert(error.responseText);
        },
      });
    }
  });
}
