Meteor.methods({
    _save_file: function(f) {
        /*
        Accepts a file document as a parameter. Checks for its
        validity, and then saves to the database.

        Arguments:
            f (File): The file to store
        */

        if (Files.findOne({
                _id: f._id
            })) {
            return Files.update({
                _id: f._id
            }, f);
        } else {
            return Files.insert(_.omit(f, '_id'));
        }
    }
});
