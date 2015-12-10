// Set the page title using the {{ title X }} helper
UI.registerHelper('title', function(title) {
    document.title = title + " | GenomiK";
});
