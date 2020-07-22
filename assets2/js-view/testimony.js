let flagTestimony = 0;
function loadTestimony() {
  $.ajax({
    url: domain + "/api/testimony",
    type: "GET",
    data: {
      offset: flagTestimony,
      limit: 12,
    },
    async: false,
    success: function (resp) {
      $.each(resp.testimonies, function (idx, value) {
        $("#testimonies").append(
          '<div class="col-6 col-md-3 col-lg-3">' +
            '<article class="article article-style-c">' +
            '<div class="article-header">' +
            '<div class="article-image" data-background="/assets2/images/testimony/' +
            value.image +
            '"></div>' +
            '<div class="article-badge">' +
            '<a href="/assets2/images/testimony/' +
            value.image +
            '" class="btn btn-icon btn-primary gallery-item" style="height: auto; width: auto;"><i class="far fa-eye"></i></a>' +
            '<a href="#" class="btn btn-icon btn-danger" onclick=deleteAlert("/api/testimony/' +
            value.idTestimony +
            '","' +
            value.idTestimony +
            '") data-placement="right"><i class="far fa-trash-alt"></i></a>' +
            "</div>" +
            "</div>" +
            "</article>" +
            "</div>"
        );
      });
    },
    error: function (error) {
      console.log("error");
    },
  });
  flagTestimony += 12;
}
