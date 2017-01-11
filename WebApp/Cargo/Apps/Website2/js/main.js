// top navigation
$('#bs-example-navbar-collapse-1 li').click(function(e) {
    $('#bs-example-navbar-collapse-1 li.active').removeClass('active');
    var $this = $(this);
    if (!$this.hasClass('active')) {
        $this.addClass('active');
    }
    e.preventDefault();
    hidePages(this.id)
});

$('#a-cargo').click(function(e) {
    $('#bs-example-navbar-collapse-1 li.active').removeClass('active');
    e.preventDefault();
    hidePages(this.id)
});

function hidePages(exceptId){
	$('.pageRow').addClass('hidden')
	$('#page-' + exceptId.split('-')[1]).removeClass('hidden')
}

// tutorials navigation
$(".tutorials-link").click(function() {
    $('html, body').animate({
        scrollTop: $('#tutorialsSection-' + this.id.split('-')[1]).offset().top
    }, 500);
});

// main
main()

function main() {
     
}
