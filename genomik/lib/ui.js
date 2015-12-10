// Set the page title using the {{ title X }} helper
UI.registerHelper('title', function(title) {
    document.title = title + " | GenomiK";
});


var image_urls = {
    _img_about_card:    function() { return 'http://jordan.matelsky.com/art/img/Brain01.png' },
    _img_tech_card:     function() { return '/campus.png' },
    
    _img_yeast:     function() { return '/yeast.png' },
};

for (var k in image_urls) {
    UI.registerHelper(k, image_urls[k]);
}
