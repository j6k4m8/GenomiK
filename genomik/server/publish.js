Meteor.publish('raw_fastas', function() {
    return RawFastas.find({
        $or: [
            { owner: this.userId },
            { public: true }
        ]
    });
});
