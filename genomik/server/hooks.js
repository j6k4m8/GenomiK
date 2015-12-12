
RawFastas.on('uploaded', function(fileObj, storeName) {
    var interv = Meteor.setInterval(function() {
        console.log(fileObj.hasStored('raw_fastas'))
        if (fileObj.hasStored('raw_fastas')) {
            Meteor.clearInterval(interv);
            console.log("Running " + fileObj.copies.raw_fastas.key);
            Meteor.call('_callGo', fileObj.copies.raw_fastas.key, "~/gk/out");
        } else {
            console.log("Waiting for " + fileObj.name());
        }
    }, 1000);
});
