share.App = React.createClass
  getInitialState: ->
    page: 'None'
    title: 'About'

  setPage: (val) ->
    @setState page: val
    if val is 'NewGenome'
      @setState title: 'New Genome'
    else
      @setState title: 'About'

  render: ->
    {div} = React.DOM
    div className: 'container',
      React.createElement(share.LeftContainer, setPage: @setPage)
      div className: 'title-bar', @state.title
      div className: 'main-content',
        if @state.page is 'NewGenome'
          React.createElement(share.NewGenome, null)
        else
          div {}, "This Page Not Implemented Yet"