var infoRoutes = [
    'about',
    'algorithms',
];

var unprotectedRoutes = [
    'home',
].concat(infoRoutes);

var simpleRoutes = [

].concat(infoRoutes);


Router.plugin('ensureSignedIn', {
    except: _.pluck(AccountsTemplates.routes, 'name')
             .concat(unprotectedRoutes)
});

Router.configure({
    layoutTemplate: 'main'
});


AccountsTemplates.configureRoute('signIn');


Router.route('/', function() {
    this.render('home');
}, {
    name: 'home'
});


for (var i = 0; i < simpleRoutes.length; i++) {
    Router.route(simpleRoutes[i]);
}
