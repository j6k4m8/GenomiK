Template.assemblies_index.helpers({
    user_owned_assemblies: function() {
        return RawFastas.find({
            'owner': Meteor.userId()
        });
    },
});

Template.assembly_card.helpers({
    ownerName: function() {
        var o = Meteor.users.findOne(this.owner);
        return o.profile.first_name + " " + o.profile.last_name;
    },

    assemblyUrl: function() {
        if (this.copies.raw_fastas.key.endsWith('.gz')) {
            return '/assemblies/' + this.copies.raw_fastas.key;
        } else {
            return '/assemblies/' + this.copies.raw_fastas.key + '.gz';
        }
    }
});
