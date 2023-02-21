jQuery(window).load(function () {
    $('.pre-load').stop().delay(200).animate({opacity:0}, 500, function(){$('.pre-load').css({'display':'none'});})
});

$(document).ready(function() {

    $('#owl-demo .item').css({height: $(window).height()});
    $('#owl-demo .item img, .fixed-bg img, .main-content').css({'min-height': $(window).height()});

    $(window).resize(function (){
        $('#owl-demo .item').css({height: $(window).height()});
        $('#owl-demo .item img, .fixed-bg img, .main-content').css({'min-height': $(window).height()});
    });


    var owl_pro = $("#owl-demo-products");
    owl_pro.owlCarousel({
        itemsCustom : [
            [0, 1],
            [480, 2],
            [768, 3],
            [992, 4],
            [1200, 5]
        ],
        navigation : false
    });
    $(".next-pro").click(function(){
        owl_pro.trigger('owl.next');
    });
    $(".prev-pro").click(function(){
        owl_pro.trigger('owl.prev');
    });


    var owl = $("#owl-demo");

    owl.owlCarousel({
        navigation : false,
        slideSpeed : 500,
        paginationSpeed : 1000,
        singleItem : true,
        loop : true,
        autoPlay : true,
        autoPlayTimeout : 700
    });

    $(".next").click(function(){
        owl.trigger('owl.next');
    })
    $(".prev").click(function(){
        owl.trigger('owl.prev');
    })

    jQuery(document.documentElement).keyup(function (event) {
        var owl = jQuery(".owl-carousel");

        // handle cursor keys
        if (event.keyCode == 37) {
            // go left
            owl.trigger('owl.prev');
        } else if (event.keyCode == 39) {
            // go right
            owl.trigger('owl.next');
        }
    });

});