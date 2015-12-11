RawFastas = new FS.Collection("raw_fastas", {
    stores: [new FS.Store.FileSystem("raw_fastas", {path: "~/gk/uploads/raw_fastas"})],
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


RawFastas.allow({
    'insert': function() {
        return !!Meteor.userId();
    }
});
