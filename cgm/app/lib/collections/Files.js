Files = new FS.Collection("files", {
    stores: [new FS.Store.FileSystem("files", {path: "~/uploads"})]
});

Files.allow({
    'insert': function() {
        return true;
    }
});
