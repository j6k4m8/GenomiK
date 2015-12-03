share.NewGenome = React.createClass
  uploadGenome: ->
    console.log 'Upload Genome!'

  render: ->
    {div, h2, a} = React.DOM
    div {},
      h2 className: 'header', 'Assemble A New Genome'
      a
        className: 'waves-effect waves-light btn'
        onClick: => @uploadGenome
        'Upload'