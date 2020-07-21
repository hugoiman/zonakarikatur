function getOffer() {
  var offers = "";
  $.ajax({
    url: domain + "/api/offer",
    type: "GET",
    async: false,
    success: function (resp) {
      $.each(resp.offers, function (idx, value) {
        offers +=
          '<div class="col-12 col-md-4 col-lg-4" id="image' +
          value.idOffer +
          '">' +
          '<article class="article article-style-c">' +
          '<div class="article-header">' +
          '<div class="article-image" data-background="/assets2/images/offer/' +
          value.image +
          '"></div>' +
          '<div class="article-badge">' +
          '<a href="#" class="btn btn-icon btn-danger" data-placement="right" onclick=deleteAlert("/api/offer/' +
          value.idOffer +
          '","' +
          value.idOffer +
          '")><i class="far fa-trash-alt"></i></a>' +
          '<a class="btn btn-icon btn-primary gallery-item" href="/assets2/images/offer/' +
          value.image +
          '"style="height: auto; width: auto;"><i class="far fa-eye"></i></a>' +
          "</div>" +
          "</div>" +
          '<div class="article-details">' +
          '<div class="article-title">' +
          "<h2>" +
          '<a href="/" onclick="return false;">' +
          value.title +
          "</h2>" +
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

  $("#offers").append(offers);
}
