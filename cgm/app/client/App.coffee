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
    {div, a, i} = React.DOM
    div {},
      div className: 'top-nav red lighten-1',
        div className: 'container',
          div className: 'nav-wrapper',
            a className: 'page-title white-text', @state.title
      div className: 'container',
        a
          href: '#'
          dataActivates: 'nav-mobile'
          className: 'button-collapse top-nav full hide-on-large-only'
          i className: 'mdi-navigation-menu'
      #Need to go through and add the navbar here to make it responsive
      React.createElement(share.LeftContainer, setPage: @setPage)
      # div className: 'title-bar', @state.title
      div className: 'container',
        div className: 'row',
          div className: 'col s12 m9 l10',
            if @state.page is 'NewGenome'
              React.createElement(share.NewGenome, null)
            else
              div {}, "This Page Not Implemented Yet"