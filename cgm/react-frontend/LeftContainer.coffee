share.LeftContainer = React.createClass
  render: ->
    {div, img} = React.DOM
    div className: 'search-container',
      div className: 'search-element',
        img
          className: 'logo-inner'
          src: '/resources/logo.png'
          alt: 'Genomik'
      div className: 'search-element', "Genomes"
      div className: 'search-element', "Reads"
      div className: 'search-element', "About"
