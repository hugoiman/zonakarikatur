let flagFrame = 0;
function loadFrame() {
  $.ajax({
    url: domain + "/api/frame",
    type: "GET",
    data: {
      offset: flagFrame,
      limit: 12,
    },
    async: false,
    success: function (resp) {
      $.each(resp.frames, function (idx, value) {
        $("#frames").append(
          '<div class="col-6 col-md-3 col-lg-3 all ' +
            value.model +
            '" id="image' +
            value.idFrame +
            '">' +
            '<article class="article article-style-c" title="' +
            value.model +
            '">' +
            '<div class="article-header">' +
            '<div class="article-image" data-background="' +
            value.image +
            '"></div>' +
            '<div class="article-badge">' +
            '<a href="' +
            value.image +
            '" class="btn btn-icon btn-primary gallery-item" style="height: auto; width: auto;"><i class="far fa-eye"></i></a>' +
            '<a href="#" class="btn btn-icon btn-danger" onclick=deleteAlert("/api/frame/' +
            value.idFrame +
            '","' +
            value.idFrame +
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
  flagFrame += 12;
}
