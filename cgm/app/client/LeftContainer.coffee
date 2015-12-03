share.LeftContainer = React.createClass
  getInitialState: ->
    page: 'None'

  changePage: (val) ->
    if @state.page is val
      @setState page: 'None'
    else
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
          li className: 'bold',
            a
              className: 'nav-element collapsible-header waves-effect waves-teal'
              'Genomes'
            div
              className: 'collapsible-body'
              ul {},
                li {},
                  a className: 'sub-nav waves-effect', 'New Genome'
                li {},
                  a className: 'sub-nav waves-effect', 'My Genomes'
          li className: 'bold',
            a
              className: 'nav-element collapsible-header waves-effect waves-teal'
              'Reads'
            div className: 'collapsible-body',
              ul {},
                li {},
                  a className: 'sub-nav waves-effect', 'New Query'
                li {},
                  a className: 'sub-nav waves-effect', 'My Reads'
      li className: 'nav-element',
        a className: 'waves-effect waves-teal', 'About'
        # div className: 'side-nav',
        #   div className: 'search-element',
        #     img
        #       className: 'logo-inner'
        #       src: '/resources/logo.png'
        #       alt: 'Genomik'
        #   div
        #     className: 'search-element clickable'
        #     onClick: => @changePage 'Genomes'
        #     "Genomes"
        #   if @state.page is 'Genomes'
        #     div {},
        #       div
        #         className: 'sub-search'
        #         onClick: => @props.setPage 'NewGenome'
        #         'New Genome'
        #       div
        #         className: 'sub-search'
        #         onClick: => @props.setPage 'SearchGenome'
        #         'My Genomes'
        #   div
        #     className: 'search-element clickable'
        #     onClick: => @changePage 'Reads'
        #     "Reads"
        #   if @state.page is 'Reads'
        #     div {},
        #       div
        #         className: 'sub-search'
        #         onClick: => @props.setPage 'NewQuery'
        #         'New Query'
        #       div
        #         className: 'sub-search'
        #         onClick: => @props.setPage 'SearchQuery'
        #         'My Queries'
        #   div
        #     className: 'search-element clickable'
        #     onClick: => @aboutPage()
        #     "About"
