EXE = "genomik-cli";

stderr = function(args) {
    console.log(args);
};


Meteor.methods({
    '_callGo': function(filename, fileout) {
        var gzflag = "";
        if (filename.endsWith('.gz') || filename.endsWith('.gzip')) {
            gzflag = "--gz";
            Exec.run(EXE, [
                'unitig',
                "/home/ubuntu/gk/uploads/raw_fastas/" + filename,
                gzflag,
                '-o',
                fileout,
            ], console.log, stderr);
        } else {
            Exec.run(EXE, [
                'unitig',
                "/home/ubuntu/gk/uploads/raw_fastas/" + filename,
                '-o',
                fileout + ".gz",
            ], console.log, stderr);
        }

    }
});
