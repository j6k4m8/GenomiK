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
                a = fileObj;
                Meteor.setTimeout(function() {
                    Meteor.call('_callGo', a.getFileRecord().copies.raw_fastas.key, "~/gk/out", function(err, val) {
                        if (!!err) {
                            Materialize.toast(err, 5000);
                        } else {
                            Materialize.toast("Unitigs computed for " + fileObj.jobName);
                        }
                    });
                }, 1000);
            }
        });
    }
});
