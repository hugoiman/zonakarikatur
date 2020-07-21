let flagTestimony = 0;
function loadTestimony() {
  var testimonies = "";
  $.ajax({
    url: domain + "/api/testimony",
    type: "GET",
    data: {
      offset: flagTestimony,
      limit: 12,
    },
    async: false,
    success: function (resp) {
      $.getScript("/assets/js/scripts.js");
      $.each(resp.testimonies, function (idx, value) {
        testimonies +=
          '<div class="col-6 col-md-3 col-lg-3 all ' +
          value.category +
          '" id="image' +
          value.idTestimony +
          '">' +
          '<article class="article article-style-c">' +
          '<div class="article-header">' +
          '<div class="article-image" data-background="/assets2/images/testimony/' +
          value.image +
          '"></div>' +
          '<div class="article-badge">' +
          '<a href="#" class="btn btn-icon btn-danger" onclick=deleteAlert("/api/testimony/' +
          value.idTestimony +
          '","' +
          value.idTestimony +
          '") data-placement="right"><i class="far fa-trash-alt"></i></a>' +
          '<a href="/assets2/images/testimony/' +
          value.image +
          '" class="btn btn-icon btn-primary gallery-item" style="height: auto; width: auto;"><i class="far fa-eye"></i></a>' +
          "</div>" +
          "</div>" +
          "</article>" +
          "</div>";
      });
    },
    error: function (error) {
      console.log("error");
    },
  });

  $("#testimonies").append(testimonies);
  flagTestimony += 12;
}
