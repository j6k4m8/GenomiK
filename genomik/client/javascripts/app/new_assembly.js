Template.new_assembly.events({
    'change #fasta-upload': function(event, template) {
        var fastaFile = new FS.File(event.target.files[0]);
        fastaFile.owner = Meteor.userId();
        fastaFile.jobName = $('#job-name').val();
        fastaFile.public = $('#job-public').val() == 'on';

        RawFastas.insert(fastaFile, function(err, fileObj) {
            if (err) {
                Materialize.toast('Invalid FASTA.', 5000);
            } else {
                Router.go('home');
                Materialize.toast('New upload started: ' + fileObj.jobName, 5000);
            }
        });
    }
});
