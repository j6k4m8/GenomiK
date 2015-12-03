Genomes = new Meteor.Collection('genomes', {
    transform: function(doc) {
        return new Genome(doc);
    }
});

/*
 *  Genomes
 *  Genomes contain metadata about a genome (author, date created, etc) and
 *  a fileId that refers to a location on disk where the genome's file can
 *  be found.
 */

Genome = function(doc) {
    doc.created = new Date();
    _.extend(this, doc);
};

Genome.prototype = {
    constructor: Genome,

    // Verify that a Genome has requisite params:
    isValid: function() {
        return !!this.fileId;
    },

    save: function() {
        Meteor.call('_save_genome', this, function(err, val) {
            if (!!err) {
                console.log(err);
            }
        });
    }
};
