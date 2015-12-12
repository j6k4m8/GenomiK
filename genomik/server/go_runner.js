EXE = "genomik-cli";

stderr = function(args) {
    console.log(args);
};


Meteor.methods({
    '_callGo': function(filename, fileout) {
        var gzflag = "";
        if (filename.endsWith('.gz') || filename.endsWith('.gzip')) {
            gzflag = "--gz";
        }

        Exec.run(EXE, [
            'unitig',
            '-o',
            fileout,
            "/home/ubuntu/gk/uploads/raw_fastas/" + filename,
            gzflag,
        ], console.log, stderr);
    }
});
