Meteor.startup ->
  React.render(
    React.createElement(share.App, null), document.getElementById("render-target")
  )
  # document.ready ->
  #   document.getElementByClass('.collapsible').collapsible({
  #     accordion: false
  #   });
