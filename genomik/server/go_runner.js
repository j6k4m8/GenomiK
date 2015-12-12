EXE = "genomik-cli";

stderr = function(args) {
    console.log(args);
};


runCommand = function(cmd, args, stdout) {
    /*
    Run the actual command.

    Arguments:
        cmd (str): The command to run
        args (str[]): Arguments
    */

    // Should do sanitation here first
    Exec.run(cmd, args, stdout, stderr);
};


Meteor.methods({
    '_callGo': function(filename, fileout) {
        var gzflag = "";
        if (filename.endsWith('.gz') || filename.endsWith('.gzip')) {
            gzflag = "--gz";
        }

        runCommand(EXE, [
            'unitig',
            '-o',
            fileout,
            gzflag,
            '"' + "~/gk/uploads/raw_fastas/" + filename + '"'
        ], console.log);
    }
});
