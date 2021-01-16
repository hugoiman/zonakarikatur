let flagClient = 0;
function loadClient() {
  $.ajax({
    url: domain + "/api/client",
    type: "GET",
    data: {
      offset: flagClient,
      limit: 12,
    },
    async: false,
    success: function (resp) {
      $.each(resp.clients, function (idx, value) {
        $("#clients").append(
          '<div class="col-6 col-md-3 col-lg-3" id="image' +
            value.idClient +
            '">' +
            '<article class="article article-style-c">' +
            '<div class="article-header">' +
            '<div class="article-image" data-background="' +
            value.image +
            '"></div>' +
            '<div class="article-badge">' +
            '<a href="' +
            value.image +
            '" class="btn btn-icon btn-primary gallery-item" style="height: auto; width: auto;"><i class="far fa-eye"></i></a>' +
            '<a href="#" class="btn btn-icon btn-danger" onclick=deleteAlert("/api/client/' +
            value.idClient +
            '","' +
            value.idClient +
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
  flagClient += 12;
}
