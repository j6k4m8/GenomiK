share.LeftContainer = React.createClass
  getInitialState: ->
    page: 'None'

  changePage: (val) ->
    if @state.page is val
      @setState page: 'None'
    else
      @setState page: val

  render: ->
    {div, img} = React.DOM
    div className: 'search-container',
      div className: 'search-element',
        img
          className: 'logo-inner'
          src: '/resources/logo.png'
          alt: 'Genomik'
      div
        className: 'search-element clickable'
        onClick: => @changePage 'Genomes'
        "Genomes"
      if @state.page is 'Genomes'
        div {},
          div className: 'sub-search', 'New Genome'
          div className: 'sub-search', 'Search'
      div
        className: 'search-element clickable'
        onClick: => @changePage 'Reads'
        "Reads"
      if @state.page is 'Reads'
        div {},
          div className: 'sub-search', 'New Query'
          div className: 'sub-search', 'Search'
      div
        className: 'search-element clickable'
        onClick: => @changePage 'About'
        "About"
