Files = new Meteor.Collection('files', {
    transform: function(doc) {
        return new File(doc);
    }
});

/*
 *  Files
 *  Files contain metadata about a File (author, date created, etc)
 */

File = function(doc) {
    doc.created = new Date();
    _.extend(this, doc);
};

File.prototype = {
    constructor: File,

    // Verify that a File has requisite params:
    isValid: function() {
        return !!this.path;
    },

    save: function() {
        Meteor.call('_save_file', this, function(err, val) {
            if (!!err) {
                console.log(err);
            }
        });
    }
};
