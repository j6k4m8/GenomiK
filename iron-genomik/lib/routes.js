Router.configure({
    layoutTemplate: 'main'
});


AccountsTemplates.configureRoute('signIn');


Router.route('/', function() {
    this.render('home');
});
