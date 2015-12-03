
runGo = function(executable, arguments, stdout, sterr) {
    if (executable.indexOf(';') >= 0) {
        throw new Meteor.Error(500, "Improperly formed executable.");
    }
    Exec.run(executable, arguments, stdout, sterr);
};
