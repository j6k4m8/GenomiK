stdErrorHandler = function(args) {
    /*
    A default error handler, if none is specified.

    Arguments:
        args (string[]): Whatever
    */
    throw new Meteor.Error(501, "STDERR: " + args);
};


runGo = function(executable, arguments, stdout, stderr) {
    /*
    Runs the actual go executable in a separate thread.

    Arguments:
        executable (str): The name of the executable
        arguments (str[]): An array of arguments (ordered) to pass
        stdout (function): A handler for standard-output
        stderr (function): A handler for errors, or stdErrorHandler if
            none is specified.
    */
    if (executable.indexOf(';') >= 0) {
        // HACKS!
        throw new Meteor.Error(500, "Improperly formed executable.");
    }

    if (!stderr || typeof stdoutHandler !== 'function') {
        // Use the default if no handler is provided.
        stderr = stdErrorHandler;
    }
    Exec.run(executable, arguments, stdout, stderr);
};
