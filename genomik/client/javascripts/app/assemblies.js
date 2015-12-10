Template.assemblies_index.helpers({
    user_owned_assemblies: function() {
        return RawFastas.find({
            'owner': Meteor.userId()
        });
    }
});
