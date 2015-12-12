
RawFastas.on('uploaded', function(fileObj, storeName) {
    var interv = Meteor.setInterval(function() {
        if (fileObj.hasStored('raw_fastas')) {
            Meteor.clearInterval(interv);
            console.log("Running " + fileObj.copies.raw_fastas.key);
            Meteor.call('_callGo', fileObj.copies.raw_fastas.key, "/home/ubuntu/go/src/github.com/j6k4m8/cg/genomik/public/assemblies/" + fileObj.copies.raw_fastas.key);
        } else {
            console.log("Waiting for " + fileObj.name());
        }
    }, 1000);
});
