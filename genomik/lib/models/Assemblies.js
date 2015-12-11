Assemblies = new FS.Collection("assemblies", {
    stores: [new FS.Store.FileSystem("assemblies", {path: "~/gk/uploads/assemblies"})],
    allow: {
        extensions: ['fa', 'fasta']
    },
    onInvalid: function(message) {
        if (Meteor.isClient) {
            Materialize.toast(message, 5000);
        } else {
            console.log(message);
        }
    }
});


Assemblies.allow({
    'insert': function () {
        return Meteor.isServer;
    }
});
