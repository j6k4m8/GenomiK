share.NewGenome = React.createClass
  uploadGenome: ->
    console.log 'Upload Genome!'

  render: ->
    {div, h2} = React.DOM
    div {},
      h2 className: 'header', 'Assemble A New Genome'
      div
        className: 'upload-button'
        onClick: => @uploadGenome
        'Upload'