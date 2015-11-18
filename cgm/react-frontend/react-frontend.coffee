if Meteor.isClient
  Meteor.startup ->
    React.render(
      React.createElement(share.App, null), document.getElementById("render-target")
    )
