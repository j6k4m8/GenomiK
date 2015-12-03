Meteor.methods({
    _save_sequence: function(seq) {
        /*
        Accepts a sequence document as a parameter. Checks for its
        validity, and then saves to the database.

        Arguments:
            seq (Sequence): The Sequence to store
        */

        if (Sequences.findOne({
                fileId: seq.fileId
            })) {
            return Sequences.update({
                fileId: seq.fileId
            }, seq);
        } else {
            return Sequences.insert(_.omit(seq, '_id'));
        }
    }
});
