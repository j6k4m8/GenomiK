share.App = React.createClass
  render: ->
    {div} = React.DOM
    div className: 'container',
      React.createElement(share.LeftContainer, null)