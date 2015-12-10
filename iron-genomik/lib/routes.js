// Routes for things like the About page
var infoRoutes = [
    'about',
    'algorithms',
];

// Routes that don't require user to be logged in
var unprotectedRoutes = [
    'home',
    'atSignIn',
].concat(infoRoutes);

// Routes that can be simply 'routed' without any
// extra logic (i.e. not 'Home')
var simpleRoutes = [
].concat(infoRoutes);

// Configure the routes that don't require login
// (this is why we made unprotectedRoutes)
Router.plugin('ensureSignedIn', {
    except: _.pluck(AccountsTemplates.routes, 'name')
             .concat(unprotectedRoutes)
});

// Configure the router itself
Router.configure({
    layoutTemplate: 'main'
});

// Add the 'signIn' route.
AccountsTemplates.configureRoute('signIn');
AccountsTemplates.addFields([
    {
        _id: 'first_name',
        type: 'text',
        displayName: "First Name",
        required: true
    },
    {
        _id: 'last_name',
        type: 'text',
        displayName: "Last Name",
        required: true
    }
]);

// The default (Home) template
Router.route('/', function() {
    this.render('home');
}, {
    name: 'home'
});

// All the simple template names can be routed like this.
for (var i = 0; i < simpleRoutes.length; i++) {
    Router.route(simpleRoutes[i]);
}
