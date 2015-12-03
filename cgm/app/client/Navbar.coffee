share.Navbar = React.createClass
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
      div className: 'top-nav blue lighten-2',
        div className: 'container',
          div className: 'nav-wrapper',
            a className: 'page-title white-text', @state.title
      div className: 'container',
        a
          href: '#'
          dataActivates: 'nav-mobile'
          className: 'button-collapse top-nav full hide-on-large-only'
          i className: 'mdi-navigation-menu'
      React.createElement(share.LeftContainer, setPage: @setPage)
