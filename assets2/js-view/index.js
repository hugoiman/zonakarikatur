var domain = "https://zonakarikatur.herokuapp.com";
// var domain = "http://localhost:8080";

function getBtnWhatsapp() {
  var options = {
    whatsapp: 6281286185009, // WhatsApp number
    call_to_action: "Hi, ada pertanyaan?", // Call to action
    position: "right", // Position may be 'right' or 'left'
  };
  var proto = document.location.protocol,
    host = "getbutton.io",
    url = proto + "//static." + host;
  var s = document.createElement("script");
  s.type = "text/javascript";
  s.async = true;
  s.src = url + "/widget-send-button/js/init.js";
  s.onload = function () {
    WhWidgetSendButton.init(host, proto, options);
  };
  var x = document.getElementsByTagName("script")[0];
  x.parentNode.insertBefore(s, x);
}

function getOffer() {
  var offers = "";
  $.ajax({
    url: domain + "/api/offer",
    type: "GET",
    async: false,
    success: function (resp) {
      $.each(resp.offers, function (idx, value) {
        offers +=
          '<div class="item"><div class="media d-block mb-4 text-center site-media site-animate border-0"><a href="/assets2/images/offer/' +
          value.image +
          '" class="site-thumbnail image-popup"><img src="/assets2/images/offer/' +
          value.image +
          '" class="img-fluid"></a><div class="media-body p-md-2 p-1"><h5 class="mt-0 h4">' +
          value.title +
          '</h5></div></div></div>"';
      });
    },
    error: function (error) {
      console.log("error");
    },
  });

  $("#offers").append(offers);
  $(".rupiah").mask("000.000.000", { reverse: true });
}

function getPromo() {
  var tab_promo = "";
  var promo_list = "";
  $.ajax({
    url: domain + "/api-promos",
    type: "GET",
    async: false,
    success: function (resp) {
      $.each(resp.promo_list, function (idx, value) {
        tab_promo +=
          '<li class="nav-item site-animate"><a class="nav-link' +
          (idx == 0 ? " active" : "") +
          '" id="pills-' +
          value.id_promo +
          '-tab" data-toggle="pill" href="#pills-' +
          value.id_promo +
          '" role="tab" aria-' +
          'controls="pills-' +
          value.id_promo +
          '" aria-selected="true">' +
          value.title +
          "</a></li>";

        promo_list +=
          '<div class="tab-pane fade' +
          (idx == 0 ? " show active" : "") +
          '" id="pills-' +
          value.id_promo +
          '" role="tabpanel" aria-labelledby="pills-' +
          value.id_promo +
          '-tab"><div clas' +
          's="col-md-12"><div class="row">' +
          (!value.image.includes(".")
            ? '<div class="col-md-5 site-animate"><iframe width="400" height="315" src="https' +
              "://www.youtube.com/embed/" +
              value.image +
              '" frameborder="0" allow="accelerom' +
              'eter; autoplay; encrypted-media; gyroscope; picture-in-picture" allowfullscree' +
              "n></iframe>"
            : '<div class="col-md-3 site-animate"><a href="/assets2/images/promo/' +
              value.image +
              '" class="site-thumbnail image-popup"><img src="/assets2/images/promo/' +
              value.image +
              '" class="img-fluid"></a>') +
          "</div>" +
          (!value.image.includes(".")
            ? '<div class="col-md-7">'
            : '<div class="col-md-9">') +
          '<h5 class="mt-0 h4">' +
          value.title +
          "</h5>" +
          value.description +
          '<p class="mb-4">You get: <br>' +
          value.benefit +
          '</p><p class="mb-4"><small>Note: ' +
          "<br>" +
          value.note +
          "</small></p></div></div></div></div>";
      });
    },
    error: function (error) {
      console.log("error");
    },
  });
  $("#pills-tab").append(tab_promo);
  $("#promo_list").append(promo_list);
}

var flagTestimony = 0;
function loadTestimony() {
  var testimonies = "";
  var content = "";
  open_div =
    '<div class="carousel-item ' +
    (flagTestimony === 0 ? " active" : "") +
    '"><div class="row">';

  close_div = "</div></div>";

  $.ajax({
    url: domain + "/api/testimony",
    type: "GET",
    data: {
      offset: flagTestimony,
      limit: 6,
    },
    async: false,
    success: function (resp) {
      $.getScript("/assets2/js/main.js", function () {
        //
      });
      $.each(resp.testimonies, function (idx, value) {
        content +=
          '<div class="col-md-2 col-4 site-animate"><a href="/assets2/images/testimoni/' +
          value.image +
          '" class="site-thumbnail image-popup"><img src="/assets2/images/testimoni/' +
          value.image +
          '" class="img-fluid"></a></div>';
      });
    },
    error: function (error) {
      console.log("error");
    },
  });
  testimonies = open_div + content + close_div;
  $("#testimonies").append(testimonies);
  flagTestimony += 6;
}

function getGallery() {
  var galleries = "";

  $.ajax({
    url: domain + "/api/gallery",
    type: "GET",
    data: {
      offset: 0,
      limit: 12,
    },
    async: false,
    success: function (resp) {
      galleries = "";
      $.each(resp.galleries, function (idx, value) {
        galleries +=
          '<div class="col-md-2 col-4 site-animate g' +
          idx +
          " " +
          value.category +
          '"><a href="/assets2/images/gallery/' +
          value.image +
          '" class="site-thumbnail image-popup"><img src="/assets2/images/gallery/' +
          value.image +
          '" class="img-fluid" alt="' +
          value.category +
          '"></a></div>';
      });
    },
    error: function (error) {
      console.log("error");
    },
  });
  $(".site-custom-gutters").append(galleries);
}

var flagGallery = 0;
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
      if (resp.galleries == null) {
        $("#image-loader").attr("hidden", true);
      }
      $.getScript("/assets2/js/main.js", function () {
        //
      });

      galleries = "";
      $.each(resp.galleries, function (idx, value) {
        galleries +=
          '<div class="col-md-2 col-4 site-animate g' +
          idx +
          " " +
          value.category +
          ' all"><a href="/assets2/images/gallery/' +
          value.image +
          '" class="site-thumbnail image-popup"><img src="/assets2/images/gallery/' +
          value.image +
          '" class="img-fluid" alt="' +
          value.category +
          '"></a></div>';
      });
    },
    error: function (error) {
      console.log("error");
    },
  });
  $(".site-custom-gutters").append(galleries);
  flagGallery += 12;
}

// if ('serviceWorker' in navigator) {   window.addEventListener('load',
// function() {
// navigator.serviceWorker.register('/sw').then(function(registration) {
// Registration was successful       console.log('ServiceWorker registration
// successful with scope: ', registration.scope);     }, function(err) {
// registration failed :(       console.log('ServiceWorker registration failed:
// ', err);     });   }); } else {   console.log("ServiceWorker belum didukung
// browser ini."); }
