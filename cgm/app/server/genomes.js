Meteor.methods({
    _save_genome: function(gen) {
        /*
        Accepts a genome document as a parameter. Checks for its
        validity, and then saves to the database.

        Arguments:
            gen (Genome): The genome to store
        */

        if (Genomes.findOne({
                fileId: gen.fileId
            })) {
            return Genomes.update({
                fileId: gen.fileId
            }, gen);
        } else {
            return Genomes.insert(_.omit(gen, '_id'));
        }
    }
});
