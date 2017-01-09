// top-navbar
$('#top-navbar li').click(function(e) {
    $('#top-navbar li.active').removeClass('active');
    var $this = $(this);
    if (!$this.hasClass('active')) {
        $this.addClass('active');
    }
    e.preventDefault();
});

$('#a-cargo').click(function(e) {
    $('#top-navbar li.active').removeClass('active');
    e.preventDefault();
});


main()

function main() {
    //document.getElementsByTagName("body")[0].innerHTML = "A"

}

function testFct(){
	alert('a')
}