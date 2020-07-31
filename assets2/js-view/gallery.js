let flagGallery = 0;
function loadGallery() {
  $.ajax({
    url: domain + "/api/gallery",
    type: "GET",
    data: {
      offset: flagGallery,
      limit: 12,
    },
    async: false,
    success: function (resp) {
      $.each(resp.galleries, function (idx, value) {
        $("#galleries").append(
          '<div class="col-6 col-md-3 col-lg-3 all ' +
            value.category +
            '" id="image' +
            value.idGallery +
            '">' +
            '<article class="article article-style-c" title="' +
            value.category +
            '">' +
            '<div class="article-header">' +
            '<div class="article-image" data-background="' +
            value.image +
            '"></div>' +
            '<div class="article-badge">' +
            '<a href="' +
            value.image +
            '" class="btn btn-icon btn-primary gallery-item" style="height: auto; width: auto;"><i class="far fa-eye"></i></a>' +
            '<a href="#" class="btn btn-icon btn-danger" onclick=deleteAlert("/api/gallery/' +
            value.idGallery +
            '","' +
            value.idGallery +
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
  flagGallery += 12;
}
