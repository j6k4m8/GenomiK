Sequences = new Meteor.Collection('sequences', {
    transform: function(doc) {
        return new Sequence(doc);
    }
});

/*
 *  Sequences
 *  Sequences contain metadata about a Sequence (author, etc) and
 *  a fileId that refers to a location on disk where a Sequence's file can
 *  be found.
 */

Sequence = function(doc) {
    doc.created = new Date();
    _.extend(this, doc);
};

Sequence.prototype = {
    constructor: Sequence,

    // Verify that a Sequence has requisite params:
    isValid: function() {
        return !!this.fileId;
    },

    save: function() {
        Meteor.call('_save_sequence', this, function(err, val) {
            if (!!err) {
                console.log(err);
            }
        });
    }
};
