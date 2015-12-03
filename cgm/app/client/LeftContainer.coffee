share.LeftContainer = React.createClass
  getInitialState: ->
    page: 'None'

  changePage: (val) ->
    if @state.page is val
      @setState page: 'None'
    else
      @setState page: val

  setPageWithHighlight: (val) ->
    @props.setPage val
    @setState page: val

  aboutPage: ->
    @setState page: 'About'
    @props.setPage 'About'

  render: ->
    {div, img, nav, ul, li, a} = React.DOM
    ul
      className: 'side-nav fixed'
      id: 'nav-mobile'
      li className: 'logo',
        div className: 'search-element',
          img
            className: 'logo-inner'
            src: '/resources/logo.png'
            alt: 'Genomik'
      li className: 'no-padding',
        ul className: 'collapsible collapsible-accordion',
          li className: 'bold nav-header',
            a
              className: 'nav-element collapsible-header waves-effect waves-teal'
              'Genomes'
            div
              className: 'collapsible-body'
              ul {},
                li className: "#{'red lighten-1' if @state.page is 'NewGenome'}",
                  a
                    className: "sub-nav clickable #{'white-text' if @state.page is 'NewGenome'}"
                    onClick: => @setPageWithHighlight 'NewGenome'
                    'New Genome'
                li className: "#{'red lighten-1' if @state.page is 'SearchGenome'}",
                  a
                    className: "sub-nav clickable #{'white-text' if @state.page is 'SearchGenome'}"
                    onClick: => @setPageWithHighlight 'SearchGenome'
                    'My Genomes'
          li className: 'bold',
            a
              className: 'nav-element collapsible-header waves-effect waves-teal'
              'Reads'
            div className: 'collapsible-body',
              ul {},
                li className: "#{'red lighten-1' if @state.page is 'NewQuery'}",
                  a
                    className: "sub-nav clickable #{'white-text' if @state.page is 'NewQuery'}"
                    onClick: => @setPageWithHighlight 'NewQuery'
                    'New Query'
                li className: "#{'red lighten-1' if @state.page is 'SearchQuery'}",
                  a
                    className: "sub-nav clickable #{'white-text' if @state.page is 'SearchQuery'}"
                    onClick: => @setPageWithHighlight 'SearchQuery'
                    'My Reads'
      li className: 'nav-element',
        a
          className: 'waves-effect waves-teal'
          onClick: => @setPageWithHighlight 'About'
          'About'
